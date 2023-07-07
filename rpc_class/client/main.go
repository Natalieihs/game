package main

import (
	"net/rpc"
	"time"
)

type Request struct {
	NumOne int
	NumTwo int
}

type Response struct {
	Num int
}

func main() {
	req := Request{NumOne: 1, NumTwo: 20}
	client, err := rpc.DialHTTP("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	var res Response
	ca := client.Go("Server.Add", req, &res, nil)

	for {
		select {
		case <-ca.Done:
			println(res.Num)
			return
		default:
			time.Sleep(time.Second)
			println("wait...")
		}
	}
	println(res.Num)
}
