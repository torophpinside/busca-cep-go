package cep

import (
	"encoding/json"
	"net/http"

	"busca-cep-go/config"
	"busca-cep-go/model"

	"github.com/gin-gonic/gin"
)

func FindCep(context *gin.Context) {
	cep := &model.Cep{}

	cepKey := "cep_key_" + context.Param("cep")
	val, err := config.GetRedis().Get(cepKey).Result()
	if err == nil && val != "" {
		err = json.Unmarshal([]byte(val), cep)
		buildResponse(context, cep)
		return
	}

	resp, err := http.Get("https://viacep.com.br/ws/" + context.Param("cep") + "/json/")

	if err != nil {
		context.JSON(200, gin.H{
			"message": "cep_not_found",
		})
		return
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(cep)
	if err != nil {
		context.JSON(200, gin.H{
			"message": "cep not found",
		})
		return
	}
	out, err := json.Marshal(cep)
	if err != nil {
		panic(err)
	}

	err = config.GetRedis().Set(cepKey, string(out), 0).Err()
	if err != nil {
		panic(err)
	}

	buildResponse(context, cep)
}

func buildResponse(context *gin.Context, cep *model.Cep) {
	context.JSON(200, gin.H{
		"data": cep,
	})
}
