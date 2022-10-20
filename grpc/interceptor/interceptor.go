package interceptor

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"time"
)

type Interceptor struct {

}

func (i *Interceptor) MyInter(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error){

	//在这里就可以进行验证、jwt令牌、反爬等操作

	//调用原来的调用方式之前
	start := time.Now()
	hand, err := handler(ctx, req)
	if err != nil {
		panic(err)
	}
	//可以在此中间进行一个操作，比如计算业务处理时常等
	since := time.Since(start)
	fmt.Println(since)
	//调用原来的调用方式之后

	return hand,err
}