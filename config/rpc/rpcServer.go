package rpc

import (
	"log"
	"net"
	"net/http"
	"net/rpc"

	cepRpcBase "busca-cep-go/buscaCep/cep/transport/http/rpc"
)

func RegisterRpcServer() {
	cepRpc := new(cepRpcBase.CepRpc)
	err := rpc.Register(cepRpc)
	if err != nil {
		log.Fatal("Format of service Task isn't correct. ", err)
	}

	rpc.HandleHTTP()

	listener, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("Listen error: ", e)
	}

	log.Printf("Serving RPC server on port %d", 1234)

	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatal("Error serving: ", err)
	}
}
