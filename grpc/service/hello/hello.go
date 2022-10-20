package hello

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	hellopb "shopping/grpc/protobuf/api/gen/v1"
)

type Service struct {

	hellopb.UnimplementedHelloServer
}


func (s *Service)HelloWorld(c context.Context,req *hellopb.HelloRequest) (rep *hellopb.HelloResponse,err error){

	//return nil, status.Errorf(codes.OK, "method HelloWorld ")
	return &hellopb.HelloResponse{

	}, nil
}

func (s *Service) Ping(ctx context.Context,he  *hellopb.Empty) (p *hellopb.Pong,err  error) {
	md, b := metadata.FromIncomingContext(ctx)
	if !b {
		return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
	}

	fmt.Println(md)

	return &hellopb.Pong{
		Id: "I am Ping",
	},nil
}
func (s *Service) Pings(ctx context.Context,ee *emptypb.Empty) (pong *hellopb.Pong,err error) {


	md, b := metadata.FromIncomingContext(ctx)
	if !b {
		return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
	}

	fmt.Println(md)

	return &hellopb.Pong{
		Id: "I am Pings",
	},nil
}