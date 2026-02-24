package client

import (
	"context"
	"flagd/config"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

func ConnectPostgres(config config.Config) *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), config.DatabaseURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
	return conn
}
