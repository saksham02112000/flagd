package routes

import (
	"flagd/config"
	handler "flagd/internal/handlers"
	"flagd/internal/service"
	"flagd/internal/store/postgres"
	"flagd/internal/store/repository"
	"log/slog"

	"github.com/go-redis/redis/v8"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Container struct {
	DB     *pgxpool.Pool
	Redis  *redis.Client
	Config *config.Config
	Logger *slog.Logger

	FlagRepo repository.FlagRepository

	FlagService *service.FlagService

	FlagHandler *handler.FlagHandler
}

func Build(cfg *config.Config, db *pgxpool.Pool, rdb *redis.Client, logger *slog.Logger) *Container {
	c := &Container{
		DB:     db,
		Redis:  rdb,
		Config: cfg,
		Logger: logger,
	}

	c.FlagRepo = postgres.NewPostgresFlagRepository(c.DB)
	c.FlagService = service.NewFlagService(c.FlagRepo)
	c.FlagHandler = handler.NewFlagHandler(c.FlagService)

	return c
}
