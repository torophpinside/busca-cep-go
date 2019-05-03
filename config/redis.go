package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/go-redis/redis"
	"github.com/joho/godotenv"
)

var redisClient *redis.Client

func init() {
	envFile := os.Getenv("ENV_FILE")
	if envFile == "" {
		envFile = ".env"
	}
	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, _ := strconv.Atoi(os.Getenv("REDIS_DB"))
	conn := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       db,
	})

	_, err = conn.Ping().Result()
	if err != nil {
		fmt.Println("error connecting redis")
		os.Exit(1)
	}
	redisClient = conn
}

func GetRedis() *redis.Client {
	return redisClient
}
