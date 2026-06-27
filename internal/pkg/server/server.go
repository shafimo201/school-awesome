package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	"github.com/smoha201/school-awesome/internal/pkg/config"
	"github.com/smoha201/school-awesome/internal/adapter/api"
)

func New(cfg *config.Config, db *pgxpool.Pool, logger zerolog.Logger) *http.Server {
	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(api.LoggingMiddleware(logger))
	engine.GET("/health", api.HealthHandler)

	v1 := engine.Group("/api/v1")
	{
		v1.GET("/health", api.HealthHandler)
		v1.GET("/users", api.ListUsersHandler)
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
