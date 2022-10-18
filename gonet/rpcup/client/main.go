package main

import "shopping/gonet/rpcup/client_proxy"

func main() {

	/*conn, err := net.Dial("tcp", ":8090")
	if err != nil {
		panic(conn)
	}

	cli := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))


	var reply = ""

	err = cli.Call("HelloService.Hello", "bobby", &reply)
	if err != nil {
		panic(err)
	}

	fmt.Println(reply)
*/

	//想要做成cli.Hello的方式。那么就需要在rpc的基础上进行封装
	cli := client_proxy.NewHelloClient()
	var reply = ""
	cli.Hello("bobby", reply)
}
