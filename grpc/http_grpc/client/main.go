package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc/protos"
)



func main()  {
	//链接grpc服务
	cli, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure()) // 使用不安全的链接
	if err != nil {
		panic(err)
	}

	defer cli.Close()
	c := protos.NewEchoServiceClient(cli)
	res ,err := c.UnaryEcho(context.Background(), &protos.EchoRequest{
		Req: "grpc-rpc",
	})
	fmt.Println("resp: ", res.GetRsp())

}
