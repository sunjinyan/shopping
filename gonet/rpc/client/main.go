package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	/*dial, err := rpc.Dial("tcp", "localhost:8090")
	if err != nil {
		panic(err)
	}


	reply := ""

	err = dial.Call("HelloService.Hello", "bobby", &reply)

	if err != nil {
		panic(err)
	}

	fmt.Println(reply)*/

    //没用到
	//conn, err := net.Dial("tcp", ":8090")
	//client := rpc.NewClient(conn)

	conn, err := net.Dial("tcp", ":8090")
	if err != nil {
		panic(err)
	}

	cli := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

	reply := ""
	err = cli.Call("HelloService.Hello", "bobby", &reply)
	//{"method":"HelloService.Hello","params":["bobby"],"id":0}
	//使用Call调用一个最基础的服务，这个服务内注册了所有rpc服务，
	//剩下的寻找对应的服务都是在Call中的args中去寻找
	if err != nil {
		panic(err)
	}

	fmt.Println(reply)
}
