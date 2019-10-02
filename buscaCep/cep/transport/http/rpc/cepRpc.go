package rpc

import (
	application "busca-cep-go/buscaCep/cep/application/findcep"
	infraModel "busca-cep-go/buscaCep/cep/infra/domain/model"
	"encoding/json"
	"log"
	"net/rpc"
	"time"
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
	startTime := time.Now()
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
	log.Printf("Execution time: Nanoseconds - %d, Seconds %f", time.Since(startTime).Nanoseconds(), time.Since(startTime).Seconds())
	return nil
}
