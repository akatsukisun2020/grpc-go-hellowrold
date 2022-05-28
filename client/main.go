package main

import (
	"context"
	"fmt"

	"helloworld/hellopb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// https://baijiahao.baidu.com/s?id=1730547563404362689&wfr=spider&for=pc

func main() {
	host := "127.0.0.1:8001"

	conn, err := grpc.Dial(host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	client := hellopb.NewDataClient(conn)
	req := &hellopb.UserRq{
		Id: 752111111,
	}
	rsp, err := client.GetUser(context.Background(), req)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("This is client, req:%v, rsp:%v\n", req, rsp)

}
