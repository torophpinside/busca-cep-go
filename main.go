package main

import (
	"busca-cep-go/controllers/cep"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	gin.SetMode(os.Getenv("GIN_ROUTER_MODE"))
	r := gin.Default()
	r.GET("/api/find/:cep", cep.FindCep)
	r.GET("/api/findall", cep.FindAll)
	r.Run(os.Getenv("GIN_ROUTER_PORT"))
}
