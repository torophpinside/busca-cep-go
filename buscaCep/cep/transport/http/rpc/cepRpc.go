package rpc

import (
	application "busca-cep-go/buscaCep/cep/application/findCep"
	model "busca-cep-go/buscaCep/cep/domain/model"
	infraModel "busca-cep-go/buscaCep/cep/infra/domain/model"
)

type CepRpc int

func (cepRpc *CepRpc) FindCep(cep string) *model.Cep {
	app := application.GetInstance(infraModel.GetGormInstance())
	cepData, _ := app.FindCep(cep)
	return cepData
}
