package controller

import (
	application "busca-cep-go/buscaCep/cep/application/findCep"
	model "busca-cep-go/buscaCep/cep/infra/domain/model"
	router "busca-cep-go/config"

	"github.com/gin-gonic/gin"
)

func Routes() {
	r := router.GetRouter()
	r.GET("/api/cep/find/:cep", findCep)
}

func findCep(context *gin.Context) {
	app := application.GetInstance(model.GetInstance())
	cep := app.FindCep(context.Param("cep"))
	context.JSON(200, cep)
}
