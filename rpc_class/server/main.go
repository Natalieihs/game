package main

import (
	"net"
	"net/http"
	"net/rpc"
	"time"
)

type Server struct {
}

type Request struct {
	NumOne int
	NumTwo int
}

type Response struct {
	Num int
}

// 写个Add 方法
func (s *Server) Add(req Request, res *Response) error {
	//
	time.Sleep(time.Second * 5)
	res.Num = req.NumOne + req.NumTwo
	return nil
}

func main() {
	rpc.Register(new(Server))
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	http.Serve(l, nil)
}
