package hello

import (
	"context"
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
