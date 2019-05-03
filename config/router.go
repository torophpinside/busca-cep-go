package config

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var router *gin.Engine

func init() {
	err := godotenv.Load()
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
	router.Run(os.Getenv("GIN_ROUTER_PORT"))
}