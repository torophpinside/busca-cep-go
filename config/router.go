package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var router *gin.Engine

func init() {
	envFile := os.Getenv("ENV_FILE")
	if envFile == "" {
		pathFile, err := filepath.Abs(os.Getenv("GOPATH") + "/src/busca-cep-go/.env")
		if err != nil {
			panic("env file not found")
		}
		envFile = pathFile
	}

	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	gin.SetMode(os.Getenv("GIN_ROUTER_MODE"))

	r := gin.Default()

	router = r
}

func GetRouter() *gin.Engine {
	return router
}

func Run() {
	log.Fatal(router.Run(os.Getenv("GIN_ROUTER_PORT")))
}
