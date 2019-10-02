package config

import (
	"busca-cep-go/utils"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

var router *gin.Engine

func init() {
	err := utils.LoadEnv()
	if err != nil {
		log.Fatalf("error_loading_env:%s", err.Error())
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
