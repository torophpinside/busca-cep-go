package main

import (
	rpc2 "busca-cep-go/buscaCep/cep/transport/http/rpc"
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:42586")
	if err != nil {
		log.Fatal(err)
	}

	var reply rpc2.CepRpcServiceResponse
	err = client.Call("CepRpcService.FindCep", "88036100", &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)
}
