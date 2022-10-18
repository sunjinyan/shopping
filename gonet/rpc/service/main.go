package main

import (
	"net"
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

	//1、实例化一个server，监听那个端口
	lis, err := net.Listen("tcp", ":8090")

	if err != nil {
		panic(err)
	}

	//注册服务
	//rpc.Register: type HelloService has no exported methods of
	//suitable type (hint: pass a pointer to value of that type)
	// &HelloService{} 以上错误表示需要使用指针
	err = rpc.RegisterName("HelloService", &HelloService{})

	if err != nil {
		panic(err)
	}
	for  {
		conn, err := lis.Accept()
		if err != nil {
			panic(err)
		}


		//使用的序列化协议是Gob，就不再使用rpc的包内置函数
		//rpc.Accept(lis)
		//rpc.NewServer()

		//将原有的rpc包的方法gob协议改成json协议
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}



	//rpc调用的几个问题
	//1 call id
	//2 序列化喝反序列化 编码和解码
	//客户端调用比较麻烦
	//dial.Call("HelloService.Hello", "bobby", &reply) 这种调用模式不友好
	//希望做到 client.Hello()

}
