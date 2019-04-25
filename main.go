package main

import (
	"busca-cep-go/controllers/cep"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/find/:cep", cep.FindCep)
	r.Run() // listen and serve on 0.0.0.0:8080
}
