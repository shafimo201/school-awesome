package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/smoha201/school-awesome/internal/core/port"
	"github.com/smoha201/school-awesome/internal/core/usecase"
	"github.com/smoha201/school-awesome/internal/pkg/auth"
)

type loginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=8"`
}

type loginResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresAt   string `json:"expires_at"`
}

type createUserRequest struct {
	Username string `json:"username" binding:"required"`
	FullName string `json:"full_name" binding:"required,min=3"`
	Password string `json:"password" binding:"required,min=8"`
}

func LoginHandler(userService *usecase.UserService, jwtManager *auth.JWTManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req loginRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse{Message: "invalid request payload"})
			return
		}

		user, err := userService.Authenticate(c.Request.Context(), "default-school", req.Username, req.Password)
		if err != nil {
			c.JSON(http.StatusUnauthorized, ErrorResponse{Message: "invalid credentials"})
			return
		}

		token, expiresAt, err := jwtManager.Generate(user.ID, user.SchoolID, user.Email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "failed to generate token"})
			return
		}

		c.JSON(http.StatusOK, loginResponse{AccessToken: token, ExpiresAt: expiresAt.UTC().Format("2006-01-02T15:04:05Z07:00")})
	}
}

func MeHandler(c *gin.Context) {
	userRepoVal, ok := c.Get("userRepo")
	if !ok {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "user repository not configured"})
		return
	}

	userRepo, ok := userRepoVal.(port.UserRepository)
	if !ok {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "user repository invalid"})
		return
	}

	claimsVal, ok := c.Get("authClaims")
	if !ok {
		c.JSON(http.StatusUnauthorized, ErrorResponse{Message: "authentication claims missing"})
		return
	}

	claims, ok := claimsVal.(*auth.Claims)
	if !ok {
		c.JSON(http.StatusUnauthorized, ErrorResponse{Message: "invalid auth claims"})
		return
	}

	user, err := userRepo.GetByID(c.Request.Context(), claims.SchoolID, claims.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "failed to load user"})
		return
	}
	if user == nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Message: "user not found"})
		return
	}

	c.JSON(http.StatusOK, buildUserResponse(user))
}

func CreateStudentHandler(userService *usecase.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req createUserRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse{Message: "invalid request payload"})
			return
		}

		user, err := userService.Register(c.Request.Context(), "default-school", "admin", usecase.RegisterUserInput{
			Username: req.Username,
			FullName: req.FullName,
			Password: req.Password,
			RoleID:   "student",
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "failed to create student"})
			return
		}

		c.JSON(http.StatusCreated, buildUserResponse(user))
	}
}

func CreateTeacherHandler(userService *usecase.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req createUserRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse{Message: "invalid request payload"})
			return
		}

		user, err := userService.Register(c.Request.Context(), "default-school", "admin", usecase.RegisterUserInput{
			Username: req.Username,
			FullName: req.FullName,
			Password: req.Password,
			RoleID:   "teacher",
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "failed to create teacher"})
			return
		}

		c.JSON(http.StatusCreated, buildUserResponse(user))
	}
}
