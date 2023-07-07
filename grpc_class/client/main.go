package main

import (
	"context"
	"fmt"
	hello_grpc "grpc/pb"

	"google.golang.org/grpc"
)

func main() {
	conn, e := grpc.Dial("localhost:8888", grpc.WithInsecure())
	if e != nil {
		fmt.Println(e)
	}
	defer conn.Close()
	client := hello_grpc.NewHelloGRPCClient(conn)
	req, c := client.Hello(context.Background(), &hello_grpc.Request{Message: "Hello"})
	fmt.Println(c)
	fmt.Println(req)

}
