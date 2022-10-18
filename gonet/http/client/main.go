package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	tr := &http.Transport{
		DisableKeepAlives: true,
	}

	cli := http.Client{
		Transport:     tr,
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       0,
	}

	//{"method":"HelloService.Hello","params":["bobby"],"id":0}
	/**
	固定格式，使用jsonrpc.NewServerCodec后，所需要的访问个是就是如下固定形式
	type serverRequest struct {
		Method string           `json:"method"`
		Params *json.RawMessage `json:"params"`
		Id     *json.RawMessage `json:"id"`
	}
	*/
	data := struct {//固定格式
		Method string  `json:"method"`
		Params []string  `json:"params"`
		Id int         `json:"id"`
	}{
		Method: "HelloService.Hello",
		Params: []string {"bobby"},
		Id: 1,//请求的trance 序列号
	}
	reqData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	reader := bytes.NewReader(reqData)
	post, err := cli.Post("http://localhost:8090/jsonrpc", "application/json", reader)
	if err != nil {
		panic(err)
	}
	body, _ := ioutil.ReadAll(post.Body)
	fmt.Println(string(body))
	//{"id":1,"result":"hello, bobby","error":null}

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

	/*conn, err := net.Dial("tcp", ":8090")
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

	fmt.Println(reply)*/
}
