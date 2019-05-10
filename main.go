package main

import (
	cepController "busca-cep-go/buscaCep/cep/transport/http/rest/controller"
	router "busca-cep-go/config"
	"busca-cep-go/config/rpc"
)

func main() {

	rpc.RegisterRpcServer()

	cepController.Routes()
	router.Run()
}
