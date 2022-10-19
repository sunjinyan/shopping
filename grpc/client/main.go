package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	hellopb "shopping/grpc/protobuf/api/gen/v1"
)

func main() {
	conn, err := grpc.Dial(":8090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	cli := hellopb.NewHelloClient(conn)

	rep, err := cli.HelloWorld(context.Background(), &hellopb.HelloRequest{})

	if err != nil {
		panic(err)
	}

	fmt.Println(rep)

}