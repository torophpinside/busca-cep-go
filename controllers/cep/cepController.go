package cep

import (
	"encoding/json"
	"net/http"

	"busca-cep-go/config"

	"github.com/gin-gonic/gin"
)

type Cep struct {
	Cep         string
	Logradouro  string
	Complemento string
	Bairro      string
	Localidade  string
	UF          string
	Unidade     string
	Ibge        string
	Gia         string
}

func FindCep(context *gin.Context) {
	cep := &Cep{}

	cepKey := "cep_key_" + context.Param("cep")
	val, err := config.GetRedis().Get(cepKey).Result()
	if err == nil && val != "" {
		err = json.Unmarshal([]byte(val), cep)
		context.JSON(200, cep)
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

	context.JSON(200, cep)
}
