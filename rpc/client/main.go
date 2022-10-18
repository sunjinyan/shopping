package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Client() {
	tr := &http.Transport{
		DisableKeepAlives: true,
	}

	client := http.Client{
		Transport:     tr,
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       0,
	}

	res, err := client.Get("http://localhost:80/add?a=1&b=2")
	if err != nil {
		panic(err)
	}

	//var p []byte

	read, err := ioutil.ReadAll(res.Body)

	//read, err := res.Body.Read(p)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(read))
}

func main() {
	Client()
	fmt.Println("i am client")
}
