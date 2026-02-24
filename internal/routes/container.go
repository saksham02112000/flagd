package routes

import (
	"flagd/config"
	handler "flagd/internal/handlers"
	"flagd/internal/service"
	"flagd/internal/store/repository"

	"github.com/go-redis/redis/v8"
	"github.com/jackc/pgx/v5"
)

type Container struct {
	DB *pgx.Conn
	Redis *redis.Client
	Config *config.Config


	FlagRepo repository.FlagRepository

	FlagService *service.FlagService

	FlagHandler *handler.FlagHandler


}


func Build(cfg *config.Config, db *pgx.Conn, rdb *redis.Client) *Container{
	c:= &Container{
		DB: db,
		Redis: rdb,
		Config: cfg,
	}

	c.FlagRepo = repository.NewFlagRepository()
	c.FlagService = service.NewFlagService(c.FlagRepo)
	c.FlagHandler = handler.NewFlagHandler(c.FlagService)

	return c
}