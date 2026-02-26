package client

import (
	"context"
	"flagd/config"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectPostgres(config config.Config) *pgxpool.Pool {
	pool, err := pgxpool.New(context.Background(), config.DatabaseURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	return pool
}
