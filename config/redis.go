package config

import (
	"busca-cep-go/utils"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/go-redis/redis"
)

var redisClient *redis.Client

func init() {
	err := utils.LoadEnv()
	if err != nil {
		log.Fatalf("error_loading_env:%s", err.Error())
	}

	db, _ := strconv.Atoi(os.Getenv("REDIS_DB"))
	conn := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       db,
	})

	result, err := conn.Ping().Result()
	if err != nil {
		fmt.Println("error connecting redis")
		os.Exit(1)
	}

	fmt.Println("redis connected:" + result)
	redisClient = conn
}

func GetRedis() *redis.Client {
	return redisClient
}
