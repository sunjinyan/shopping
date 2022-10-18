package client_proxy

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"shopping/gonet/rpcup/svc"
)

type HelloClient struct {
	*rpc.Client
}

func NewHelloClient() *HelloClient {
	conn, err := net.Dial("tcp", ":8090")
	if err != nil {
		panic(conn)
	}

	cli := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

	return &HelloClient{
		Client: cli,
	}
}

func (hc *HelloClient)Hello(args interface{}, reply interface{}) {

	//还要解决这个名字的问题
	err := hc.Call(svc.HelloServiceName + ".Hello", args, &reply)

	if err != nil {
		panic(err)
	}

	fmt.Println(reply)
}