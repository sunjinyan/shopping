package main

import (
	"io"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloService struct {

}

func (s *HelloService)Hello(req string, reply *string) error  {

	*reply = "hello, "+req

	return nil
}

func main() {

	err := rpc.RegisterName("HelloService", &HelloService{})

	if err != nil {
		panic(err)
	}

	http.HandleFunc("/jsonrpc", func(writer http.ResponseWriter, request *http.Request) {
		var conn io.ReadWriteCloser = struct {
			io.Writer
			io.ReadCloser
		}{
			ReadCloser: request.Body,
			Writer: writer,
		}
		/**
		固定格式，使用jsonrpc.NewServerCodec后，所需要的访问个是就是如下固定形式
		type serverRequest struct {
			Method string           `json:"method"`
			Params *json.RawMessage `json:"params"`
			Id     *json.RawMessage `json:"id"`
		}
		 */
		err = rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
		if err != nil {
			panic(err)
		}
	})

	err = http.ListenAndServe(":8090", nil)

	if err != nil {
		panic(err)
	}
}
