package controller

import (
	application "busca-cep-go/buscaCep/cep/application/findcep"
	model "busca-cep-go/buscaCep/cep/infra/domain/model"
	router "busca-cep-go/config"

	"github.com/gin-gonic/gin"
)

func Routes() {
	r := router.GetRouter()
	r.GET("/api/cep/find/:cep", findCep)
}

func findCep(context *gin.Context) {
	app := application.GetInstance(model.GetGormInstance())
	cep, err := app.FindCep(context.Param("cep"))
	if err != nil {
		context.AbortWithStatusJSON(404, gin.H{
			"cep":     context.Param("cep"),
			"message": err.Error(),
		})
		return
	}
	context.JSON(200, cep)
}
