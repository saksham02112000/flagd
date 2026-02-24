package client

import (
	"flagd/config"

	"github.com/go-redis/redis/v8"
)


func ConnectRedis(config config.Config) *redis.Client {
	return redis.NewClient(&redis.Options{
        Addr:     config.RedisURL,
        Password: "", 
        DB:       0,  
    })
}