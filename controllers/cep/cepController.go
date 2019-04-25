package cep

import (
	"encoding/json"
	"net/http"
	"os"

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
	resp, err := http.Get("https://viacep.com.br/ws/" + context.Param("cep") + "/json/")
	if err != nil {
		context.JSON(200, gin.H{
			"message": "cep not found",
		})
		os.Exit(1)
	}

	defer resp.Body.Close()

	var cep Cep
	err = json.NewDecoder(resp.Body).Decode(&cep)
	if err != nil {
		context.JSON(200, gin.H{
			"message": "cep not found",
		})
		os.Exit(1)
	}

	context.JSON(200, cep)
}
