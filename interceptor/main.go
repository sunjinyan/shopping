package main

import (
	"google.golang.org/grpc"
	"net"
	"shopping/grpc/interceptor"
	hellopb "shopping/grpc/protobuf/api/gen/v1"
	"shopping/grpc/service/hello"
)

func main() {

	l, err := net.Listen("tcp", ":8090")
	if err != nil {
		panic(err)
	}

	//一元的调用所使用的拦截器，stream的服务有stream的拦截器
	//通过这个函数将自己写的interceptor注册进去，看他的参数，写一个同样的函数就可以了
	inte := interceptor.Interceptor{}
	option := grpc.UnaryInterceptor(inte.MyInter)

	//stream 服务的拦截器
	//grpc.StreamInterceptor()

	s := grpc.NewServer(option)
	hellopb.RegisterHelloServer(s,&hello.Service{})

	err = s.Serve(l)

	if err != nil {
		panic(err)
	}
}