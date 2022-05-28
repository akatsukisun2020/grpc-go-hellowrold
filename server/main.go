package main

import (
	"context"
	"fmt"
	"net"

	"helloworld/hellopb"

	"google.golang.org/grpc"
)

func main() {
	listen, err := net.Listen("tcp", ":8001")
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		return
	}

	svr := grpc.NewServer()
	hellopb.RegisterDataServer(svr, &server{})
	defer func() {
		svr.Stop()
		listen.Close()
	}()

	fmt.Println("Serving 8001...")
	_ = svr.Serve(listen)
	fmt.Println("Serving Quit...")
}

type server struct {
	hellopb.UnimplementedDataServer
}

func (s *server) GetUser(ctx context.Context, req *hellopb.UserRq) (*hellopb.UserRp, error) {
	rsp := &hellopb.UserRp{
		Name: fmt.Sprintf("张三_%d", req.GetId()),
	}
	fmt.Printf("This is sever, req:%v, rsp:%v\n", req, rsp)
	return rsp, nil
}
