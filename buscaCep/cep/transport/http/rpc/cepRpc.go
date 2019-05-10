package rpc

import (
	application "busca-cep-go/buscaCep/cep/application/findCep"
	infraModel "busca-cep-go/buscaCep/cep/infra/domain/model"
	"encoding/json"
	"log"
	"net/rpc"
)

type CepRpcService int

type CepRpcServiceResponse struct {
	Cep string
}

func RegisterCepService() {
	err := rpc.Register(new(CepRpcService))
	if err != nil {
		log.Fatal(err)
	}
}

func (cepRpcService *CepRpcService) FindCep(cep string, data *CepRpcServiceResponse) error {
	app := application.GetInstance(infraModel.GetGormInstance())
	cepData, err := app.FindCep(cep)
	if err != nil {
		return err
	}

	cepDataOut, err := json.Marshal(cepData)
	if err != nil {
		return err
	}

	data.Cep = string(cepDataOut)
	return nil
}