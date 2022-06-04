package main

import (
	"context"
	"fmt"
	"net"

	"helloworld/hellopb"

	"google.golang.org/grpc"
)

func main() {
	// 创建一个监听句柄
	listen, err := net.Listen("tcp", ":8001")
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		return
	}

	// 这里是创建的一个服务server，并不是一个service！！
	// 一个server，可以提供多个service。
	// 每一个service，代表对外提供的几种服务！！
	svr := grpc.NewServer()

	// 这个实际上就是通过生成的stub存根，
	// (1) 将对应的存根的service的结构体和名字，先注册到server中
	// 将对应的接口实现，先“赋值”
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
