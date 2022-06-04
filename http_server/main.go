package main

import (
	"context"
	"fmt"
	"helloworld/hellopb"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type server struct {
	hellopb.UnimplementedDataServer // 隐式声明
}

func (s *server) GetUser(ctx context.Context, req *hellopb.UserRq) (*hellopb.UserRp, error) {

	rsp := &hellopb.UserRp{
		Name: fmt.Sprintf("name_%d", req.GetId()),
	}
	return rsp, nil
}

func main() {
	s := grpc.NewServer()
	hellopb.RegisterDataServer(s, &server{})

	// 1. 启动grpc服务
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen:%v", err)
	}

	log.Println("serving grpc on 0.0.0.0")
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve:%v", err)
		}
	}()

	// 2. 启动http服务
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalln("failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux()
	err = hellopb.RegisterDataHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    ":8080",
		Handler: gwmux,
	}

	log.Println("Serving grpc-gateway on http:0.0.0.0 :8080")
	log.Fatalln(gwServer.ListenAndServe())
}
