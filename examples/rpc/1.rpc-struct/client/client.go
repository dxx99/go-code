package main

import (
	"fmt"
	"net/rpc"
)


type Req struct {
	A int
	B int
}

type Resp struct {
	Total int
}


func main() {
	client, err := rpc.Dial("tcp", ":8001")
	if err != nil {
		panic(err)
	}
	resp := new(Resp)

	err = client.Call("Cal.Add", Req{
		A: 1,
		B: 2,
	}, resp)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.Total)


}
