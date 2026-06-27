package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/smoha201/school-awesome/internal/pkg/config"
	"github.com/smoha201/school-awesome/internal/pkg/logger"
	"github.com/smoha201/school-awesome/internal/pkg/server"
	"github.com/smoha201/school-awesome/internal/adapter/db"
	"github.com/smoha201/school-awesome/internal/adapter/api"
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

	httpServer := server.New(cfg, dbConn, logg)
	if err := httpServer.Run(); err != nil {
		logg.Fatal().Err(err).Msg("server failed")
	}
}
