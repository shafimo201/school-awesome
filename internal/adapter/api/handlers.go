package api

import (
	"errors"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/smoha201/school-awesome/internal/core/domain"
	"github.com/smoha201/school-awesome/internal/core/port"
	"github.com/smoha201/school-awesome/internal/pkg/auth"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func HealthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "healthy"})
}

func ListUsersHandler(c *gin.Context) {
	const defaultLimit = 20
	const defaultOffset = 0

	users, err := fetchUsers(c, defaultLimit, defaultOffset)
	if err != nil {
		log.Printf("ListUsersHandler error: %v", err)
		c.Error(err)
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "failed to fetch users"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": users})
}

func AuthMiddleware(jwtManager *auth.JWTManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, ErrorResponse{Message: "missing authorization header"})
			return
		}

		parts := strings.Fields(authHeader)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, ErrorResponse{Message: "invalid authorization header"})
			return
		}

		claims, err := jwtManager.Validate(parts[1])
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, ErrorResponse{Message: "invalid or expired token"})
			return
		}

		c.Set("authClaims", claims)
		c.Next()
	}
}

func fetchUsers(c *gin.Context, limit, offset int) ([]interface{}, error) {
	userRepoVal, ok := c.Get("userRepo")
	if !ok {
		return nil, errors.New("user repository not configured")
	}

	userRepo, ok := userRepoVal.(port.UserRepository)
	if !ok {
		return nil, errors.New("user repository invalid")
	}

	claimsVal, ok := c.Get("authClaims")
	if !ok {
		return nil, errors.New("authentication claims missing")
	}

	claims, ok := claimsVal.(*auth.Claims)
	if !ok {
		return nil, errors.New("invalid auth claims")
	}

	users, err := userRepo.ListBySchool(c.Request.Context(), claims.SchoolID, limit, offset)
	if err != nil {
		return nil, err
	}

	result := make([]interface{}, 0, len(users))
	for _, u := range users {
		result = append(result, buildUserResponse(u))
	}
	return result, nil
}

func buildUserResponse(user *domain.User) gin.H {
	return gin.H{
		"id":          user.ID,
		"username":    user.Email,
		"full_name":   user.FullName,
		"role_id":     user.RoleID,
		"status":      user.Status,
		"last_login":  user.LastLoginAt,
	}
}

func RepositoryMiddleware(userRepo port.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("userRepo", userRepo)
		c.Next()
	}
}

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userRepoVal, ok := c.Get("userRepo")
		if !ok {
			c.AbortWithStatusJSON(http.StatusInternalServerError, ErrorResponse{Message: "user repository not configured"})
			return
		}
		userRepo, ok := userRepoVal.(port.UserRepository)
		if !ok {
			c.AbortWithStatusJSON(http.StatusInternalServerError, ErrorResponse{Message: "user repository invalid"})
			return
		}

		claimsVal, ok := c.Get("authClaims")
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, ErrorResponse{Message: "authentication claims missing"})
			return
		}
		claims, ok := claimsVal.(*auth.Claims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, ErrorResponse{Message: "invalid auth claims"})
			return
		}

		user, err := userRepo.GetByID(c.Request.Context(), claims.SchoolID, claims.UserID)
		if err != nil || user == nil {
			c.AbortWithStatusJSON(http.StatusForbidden, ErrorResponse{Message: "forbidden"})
			return
		}

		if user.RoleID != "admin" {
			c.AbortWithStatusJSON(http.StatusForbidden, ErrorResponse{Message: "admin role required"})
			return
		}

		c.Next()
	}
}

func LoggingMiddleware(logger zerolog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		logger.Info().
			Str("method", c.Request.Method).
			Str("path", c.Request.URL.Path).
			Int("status", c.Writer.Status()).
			Dur("duration", time.Since(start)).
			Msg("request")
	}
}
