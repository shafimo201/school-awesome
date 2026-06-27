package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgres(dsn string, maxOpenConns, maxIdleConns int) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}

	config.MaxConns = int32(maxOpenConns)
	config.MinConns = int32(maxIdleConns)
	config.MaxConnLifetime = time.Minute * 30

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, err
	}

	return pool, nil
}

func Close(ctx context.Context, pool *pgxpool.Pool) {
	if pool != nil {
		pool.Close()
	}
	select {
	case <-ctx.Done():
		return
	case <-time.After(5 * time.Second):
		return
	}
}
