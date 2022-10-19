package stream

import (
	"context"
	"fmt"
	"io"
	streamspb "shopping/streamgrpc/api/gen/v1"
	"strconv"
	"sync"
	"time"
)

type Service struct {
	streamspb.UnimplementedStreamServiceServer
}

func (s *Service) GetStream(req *streamspb.StreamRequest,res streamspb.StreamService_GetStreamServer) (err error) {
	i := 0
	for  {

		i++
		err = res.Send(&streamspb.StreamResponse{Msg: "我是消息: " + strconv.Itoa(i) + " 发给: " + req.Msg})
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

	return err
}

func (s *Service) PutStream(cliStr streamspb.StreamService_PutStreamServer) (err error) {

	for  {
		recv, err := cliStr.Recv()
		if err != nil {
			if err == io.EOF && err == context.Canceled {
				break
			}
			panic(err)
		}
		fmt.Println(cliStr.Context().Value("client"),recv.Msg)
	}

	return err
}

func (s *Service) AllStream(ss streamspb.StreamService_AllStreamServer) (err error) {
	//使用协程是为了避免阻塞，从而达到了接受发送不影响的情况
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		for {
			recv, err := ss.Recv()
			if err != nil {
				defer wg.Done()
				if err == io.EOF && err == context.Canceled {
					break
				}
				panic(err)
			}
			fmt.Println("接受All Stream Client",recv.Msg)
		}
	}()

	go func() {
		i := 0
		for  {
			i++
			err = ss.Send(&streamspb.StreamResponse{Msg: "我是All 消息: " + strconv.Itoa(i)})
			if err != nil {
				defer wg.Done()
				if err == io.EOF && err == context.Canceled {
					break
				}
				panic(err)
			}
			if i >= 10  {
				defer  wg.Done()
				break
			}
		}
	}()

	wg.Wait()
	return err
}