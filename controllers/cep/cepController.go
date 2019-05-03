package cep

import (
	"encoding/json"
	"fmt"
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

func FindAll(context *gin.Context) {
	var i int64 = 11111111
	for i <= 11111999 {
		val := fmt.Sprintf("%08d", i)
		save(string(val))
		i++
	}
}

func save(val string) {
	cep := &model.Cep{}
	resp, _ := call(val)
	defer resp.Body.Close()

	err := json.NewDecoder(resp.Body).Decode(cep)
	if err != nil {
		panic(err)
	}

	out, err := json.Marshal(cep)
	if err != nil {
		panic(err)
	}

	err = config.GetRedis().Set(val, string(out), 0).Err()
	if err != nil {
		panic(err)
	}
}

func call(cep string) (*http.Response, error) {
	resp, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
	return resp, err
}

func buildResponse(context *gin.Context, cep *model.Cep) {
	context.JSON(200, gin.H{
		"data": cep,
	})
}
