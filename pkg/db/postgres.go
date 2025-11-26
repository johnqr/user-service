package db

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPostgresPool(ctx context.Context, dsn string) (*pgxpool.Pool, error) {
	cfg, err := pgxpool.ParseConfig(dsn)
	if err != nil { return nil, err }
	p, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil { return nil, err }
	return p, nil
}
