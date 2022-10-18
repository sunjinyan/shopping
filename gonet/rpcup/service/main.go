package main

import (
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"shopping/gonet/rpcup/service_proxy"
	"shopping/gonet/rpcup/svc"
)



func main() {

	lis, err := net.Listen("tcp", ":8090")
	if err != nil {
		panic(err)
	}

	//注册
	//err = rpc.RegisterName(svc.HelloServiceName, &svc.HelloService{})
	err = service_proxy.RegisterHelloService(&svc.HelloService{})
	if err != nil {
		panic(err)
	}


	for  {
		conn, err := lis.Accept()
		if err != nil {
			panic(err)
		}
		rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
