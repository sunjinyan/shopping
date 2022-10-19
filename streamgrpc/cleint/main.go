package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	streamspb "shopping/streamgrpc/api/gen/v1"
	"strconv"
	"sync"
	"time"
)

func main() {
	con, err := grpc.Dial(":8090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer  con.Close()

	cli := streamspb.NewStreamServiceClient(con)

	//服务端源源不断地返回新消息，服务端流模式
	scli, err := cli.GetStream(context.Background(), &streamspb.StreamRequest{Msg: "get stream"})
	if err != nil {
		panic(err)
	}
	for {
		recv, err := scli.Recv()
		if err != nil  {
			if err.Error() == "EOF"  {
				break
			}
			panic(err)
		}

		fmt.Println(recv.Msg)
	}

	//客户端源源不断的发送数据给客户端，客户端流模式
	//ctx, cancel := context.WithCancel(context.Background())
	//defer cancel()
	ctx  := context.WithValue(context.Background(), "client", "12")

	pcli,err := cli.PutStream(ctx)
	if err != nil {
		panic(err)
	}
	i := 0
	for  {
		i++
		err = pcli.Send(&streamspb.StreamRequest{Msg: "我是客户端发送过来的第 " + strconv.Itoa(i) + " 条消息"})
		if err != nil {
			if err.Error() == "EOF"  {
				break
			}
			panic(err)
		}
		time.Sleep(time.Second)
		if i > 10 {
			break
		}
	}


	//客户端与服务端双向流模式
	acli, err := cli.AllStream(ctx)
	if err != nil {
		panic(err)
	}
	wg := sync.WaitGroup{}
	wg.Add(2)
	j := 0
	go func() {
		for  {
			j++
			err = acli.Send(&streamspb.StreamRequest{
				Msg: "我是客户端发送过来的 All 第 " + strconv.Itoa(j) + " 条消息",
			})
			if err != nil {
				defer wg.Done()
				if err.Error() == "EOF"  {
					break
				}
				panic(err)
			}
			time.Sleep(time.Second)
			if j >= 10 {
				defer wg.Done()
				break
			}
		}
	}()

	go func() {
		for  {
			recv, err := acli.Recv()
			if err != nil {
				defer wg.Done()
				if err.Error() == "EOF"  {
					break
				}
				panic(err)
			}
			fmt.Println("接受All Service ",recv.Msg)
		}
	}()
	wg.Wait()
}
