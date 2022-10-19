package main

import (
	"google.golang.org/grpc"
	"net"
	streamspb "shopping/streamgrpc/api/gen/v1"
	stream "shopping/streamgrpc/service"
)

func main() {


	l, err := net.Listen("tcp", ":8090")
	if err != nil {
		panic(err)
	}

	//grpc.WithTransportCredentials(insecure.NewCredentials())
	s := grpc.NewServer()

	//注册服务
	streamspb.RegisterStreamServiceServer(s,&stream.Service{})

	err = s.Serve(l)
	if err != nil {
		panic(err)
	}
}
