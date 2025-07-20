package infrastructure

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func NewStorage(ctx context.Context, env *Environment) *pgxpool.Pool {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", env.DB_USER, env.DB_PASSWORD, env.DB_HOST, env.DB_PORT, env.DB_NAME)
	dbpool, err := pgxpool.New(ctx, connStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	return dbpool
}
