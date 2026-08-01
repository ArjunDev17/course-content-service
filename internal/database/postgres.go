package database

import (
	"context"
	"fmt"
	"time"

	"github.com/ArjunDev17/course-content-service/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgreSQL struct {
	Pool *pgxpool.Pool
}

func New(cfg config.DatabaseConfig) (*PostgreSQL, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	poolConfig, err := pgxpool.ParseConfig(cfg.DSN())
	if err != nil {
		return nil, fmt.Errorf("failed to parse postgres config: %w", err)
	}

	// Connection Pool Configuration
	poolConfig.MaxConns = 20
	poolConfig.MinConns = 5
	poolConfig.MaxConnLifetime = time.Hour
	poolConfig.MaxConnIdleTime = 30 * time.Minute
	poolConfig.HealthCheckPeriod = time.Minute

	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create postgres pool: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("failed to ping postgres: %w", err)
	}

	return &PostgreSQL{
		Pool: pool,
	}, nil
}

func (p *PostgreSQL) Close() {
	if p.Pool != nil {
		p.Pool.Close()
	}
}