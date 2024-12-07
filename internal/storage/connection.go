package storage

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	"log/slog"
)

func NewConnectionPool(ctx context.Context, connString string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(ctx, connString)
	if err != nil {
		return pool, err
	}
	defer pool.Close()

	err = pool.Ping(ctx)
	if err != nil {
		return pool, err
	}

	slog.Info("Connected postgres to database")

	return pool, nil
}

func NewRedisClient(ctx context.Context, connString string) (*redis.Client, error) {
	opt, err := redis.ParseURL(connString)
	if err != nil {
		return nil, err
	}

	rdb := redis.NewClient(opt)

	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	slog.Info("Connected redis to database")

	return rdb, nil
}
