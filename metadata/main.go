package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	hellopb "shopping/grpc/protobuf/api/gen/v1"
)

func main() {

	//第一种方式，使用原始的方式自行创建
	md := make(metadata.MD,0)
	s := []string{
		"abc",
	}
	md["key"] = s

	s = append(s,"def")
	md["key1"] = s

	fmt.Println(md)

	//第二种方式，使用metadata自带的new
	m := metadata.New(map[string]string{"key": "val1", "key2": "val2", "key3": "val3"})
	fmt.Println(m)

	//第三种方式，使用metadata自带的Pairs,参数数量和必须是偶数
	pmd := metadata.Pairs("key", "val1", "key2", "val2", "key3", "val3")
	fmt.Println(pmd)


	//发送metadata

	ctx := metadata.NewOutgoingContext(context.Background(), m)

	conn, err := grpc.Dial("localhost:8090",grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		panic(err)
	}

	cli := hellopb.NewHelloClient(conn)

	pong, err := cli.Ping(ctx, &hellopb.Empty{})
	if err != nil {
		panic(err)
	}
	fmt.Println(pong)
}
