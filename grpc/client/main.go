package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	hellopb "shopping/grpc/protobuf/api/gen/v1"
	"time"
)

func main() {

	//定义拦截器
	InterFun := func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {

		//可以在这里进行统一的ctx操作，比如操作metadata添加特殊的请求头部新信息或jwt认证信息等

		//更换ctx,或向ctx中添加认证信息等
		md := metadata.New(map[string]string{"a":"b"})
		metadata.NewOutgoingContext(ctx,md)

		//invoker和服务端的拦截器handler一样，在这里可以计算耗时
		start := time.Now()
		err := invoker(ctx, method, req, reply, cc, opts...)
		if err != nil {
			panic(err)
		}
		since := time.Since(start)
		fmt.Println(since)
		return err
	}

	//注册拦截器
	interceptor := grpc.WithUnaryInterceptor(InterFun)
	credentials := grpc.WithTransportCredentials(insecure.NewCredentials())

	opts := make([]grpc.DialOption,0)
	opts = append(opts,interceptor,credentials)
	conn, err := grpc.Dial(":8090", opts...)
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	cli := hellopb.NewHelloClient(conn)

	//提供一个超时机制来确保服务安全
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	rep, err := cli.HelloWorld(ctx, &hellopb.HelloRequest{})

	if err != nil {
		panic(err)
	}

	fmt.Println(rep)

}