package main

import (
	"google.golang.org/grpc"
	"net"
	hellopb "shopping/grpc/protobuf/api/gen/v1"
	"shopping/grpc/service/hello"
)

func main() {

	l, err := net.Listen("tcp", ":8090")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	hellopb.RegisterHelloServer(s,&hello.Service{})

	err = s.Serve(l)

	if err != nil {
		panic(err)
	}
}
