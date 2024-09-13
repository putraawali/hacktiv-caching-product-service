package cache

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
)

func InitRedis() *redis.Client {
	redisDb := os.Getenv("REDIS_DB")

	var db int

	if redisDb == "" {
		db = 0
	} else {
		db, _ = strconv.Atoi(redisDb)
	}

	client := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_ADDR"), // localhost:6379
		DB:   db,
	})

	if _, err := client.Ping(context.Background()).Result(); err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Success connect to redis")
	return client
}
