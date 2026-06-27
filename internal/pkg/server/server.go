package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/smoha201/school-awesome/internal/core/port"
	"github.com/smoha201/school-awesome/internal/core/usecase"
	"github.com/smoha201/school-awesome/internal/pkg/auth"
	"github.com/smoha201/school-awesome/internal/pkg/config"
	"github.com/smoha201/school-awesome/internal/adapter/api"
)

func New(cfg *config.Config, userRepo port.UserRepository, userService *usecase.UserService, jwtManager *auth.JWTManager, logger zerolog.Logger) *http.Server {
	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(api.LoggingMiddleware(logger))
	engine.Use(api.RepositoryMiddleware(userRepo))
	engine.GET("/health", api.HealthHandler)

	v1 := engine.Group("/api/v1")
	{
		v1.POST("/auth/login", api.LoginHandler(userService, jwtManager))
		v1.GET("/health", api.HealthHandler)
		v1.GET("/me", api.AuthMiddleware(jwtManager), api.MeHandler)
		v1.GET("/users", api.AuthMiddleware(jwtManager), api.ListUsersHandler)
		admin := v1.Group("/admin", api.AuthMiddleware(jwtManager), api.AdminMiddleware())
		{
			admin.POST("/students", api.CreateStudentHandler(userService))
			admin.POST("/teachers", api.CreateTeacherHandler(userService))
		}
	}

	httpServer := &http.Server{
		Addr: fmt.Sprintf(":%d", cfg.Server.Port),
		Handler: engine,
		ReadTimeout: cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
		IdleTimeout: cfg.Server.IdleTimeout,
	}

	return httpServer
}
