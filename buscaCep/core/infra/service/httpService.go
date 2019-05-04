package service

import (
	"busca-cep-go/buscaCep/cep/domain/model"
	"net/http"
)

func Call(url string, cep *model.Cep) *http.Response {
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != 200 {
		return nil
	}
	defer resp.Body.Close()

	return resp
}
