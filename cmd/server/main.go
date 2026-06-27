package main

import (
	"context"
	"log"
	"net/http"
	"time"


	"github.com/smoha201/school-awesome/internal/adapter/db"
	"github.com/smoha201/school-awesome/internal/core/usecase"
	"github.com/smoha201/school-awesome/internal/pkg/auth"
	"github.com/smoha201/school-awesome/internal/pkg/config"
	"github.com/smoha201/school-awesome/internal/pkg/logger"
	"github.com/smoha201/school-awesome/internal/pkg/server"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	cfg, err := config.Load("configs/config.yaml")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	logg := logger.New(cfg.Logger.Level)

	dbConn, err := db.NewPostgres(cfg.Database.DSN, cfg.Database.MaxOpenConns, cfg.Database.MaxIdleConns)
	if err != nil {
		logg.Fatal().Err(err).Msg("failed to connect database")
	}
	defer db.Close(ctx, dbConn)

	userRepo := db.NewUserRepository(dbConn)
	hasher := auth.NewBcryptHasher(12)
	userService := usecase.NewUserService(userRepo, hasher, logg)
	jwtManager := auth.NewJWTManager(cfg.Auth.JWTSecret, cfg.Auth.AccessTokenTTL)

	// Ensure a default admin user exists for local development.
	// Password is intentionally fixed for now per developer request.
	go func() {
		ctxSeed, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		email := "test@school.org"
		// Try to fetch; if not found, Register will create the user.
		if existing, err := userRepo.GetByEmail(ctxSeed, "default-school", email); err != nil || existing == nil {
			// create the admin if not exists
			_, err := userService.Register(ctxSeed, "default-school", "system", usecase.RegisterUserInput{
				Email:    email,
				FullName: "Admin",
				Password: "Shafi@123",
				RoleID:   "admin",
			})
			if err != nil {
				logg.Error().Err(err).Msg("failed to seed admin user")
			} else {
				logg.Info().Msg("seeded admin user: test@school.org (password: Shafi@123)")
			}
		} else {
			// ensure existing admin has the requested password for local dev
			if err := userRepo.UpdatePasswordByEmail(ctxSeed, "default-school", email, func() string {
				h, _ := hasher.Hash("Shafi@123")
				return h
			}(), "system"); err != nil {
				logg.Error().Err(err).Msg("failed to update admin password")
			} else {
				logg.Info().Msg("ensured admin password is set to Shafi@123")
			}
		}
	}()

	httpServer := server.New(cfg, userRepo, userService, jwtManager, logg)
	if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logg.Fatal().Err(err).Msg("server failed")
	}
}
