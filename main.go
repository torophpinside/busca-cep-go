package main

import (
	cepController "busca-cep-go/buscaCep/cep/transport/http/rest/controller"
	router "busca-cep-go/config"
)

func main() {
	cepController.Routes()
	router.Run()
}
