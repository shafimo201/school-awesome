package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/school-erp/project-school/internal/pkg/config"
	"github.com/school-erp/project-school/internal/pkg/logger"
	"github.com/school-erp/project-school/internal/pkg/server"
	"github.com/school-erp/project-school/internal/adapter/db"
	"github.com/school-erp/project-school/internal/adapter/api"
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
