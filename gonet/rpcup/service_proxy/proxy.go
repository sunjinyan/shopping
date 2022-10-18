package service_proxy

import (
	"net/rpc"
	"shopping/gonet/rpcup/svc"
)

type NewServer interface {
	Hello	(input string,output *string) error
}

func RegisterHelloService(srv  NewServer)  error {
	err := rpc.RegisterName(svc.HelloServiceName,srv)
	if err != nil {
		panic(err)
	}

	return err
}