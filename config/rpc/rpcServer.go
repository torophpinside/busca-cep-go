package rpc

import (
	"log"
	"net"
	"net/rpc"

	cepRpcBase "busca-cep-go/buscaCep/cep/transport/http/rpc"
)

func RegisterRpcServer() {
	addy, err := net.ResolveTCPAddr("tcp", "0.0.0.0:42586")
	if err != nil {
		log.Fatal(err)
	}

	inbound, err := net.ListenTCP("tcp", addy)
	if err != nil {
		log.Fatal(err)
	}

	cepRpcBase.RegisterCepService()
	rpc.Accept(inbound)
}
