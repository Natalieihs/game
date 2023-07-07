package main

import (
	"context"
	"fmt"
	hello_grpc "grpc/pb"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	hello_grpc.UnimplementedHelloGRPCServer
}

func (s *server) Hello(ctx context.Context, req *hello_grpc.Request) (*hello_grpc.Response, error) {

	fmt.Println(req.GetMessage())
	return &hello_grpc.Response{Message: "Hi " + req.GetMessage()}, nil
}

func main() {

	netListen, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println(err)
	}
	s := grpc.NewServer()
	hello_grpc.RegisterHelloGRPCServer(s, &server{})
	s.Serve(netListen)
}
