package main

import (
	"bufio"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc/protos"
	"os"
)



func main()  {
	//链接grpc服务
	cli, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure()) // 使用不安全的链接
	if err != nil {
		panic(err)
	}

	defer cli.Close()
	c := protos.NewEchoServiceClient(cli)

	reader := bufio.NewReader(os.Stdin)

	stream, err := c.CStreamEcho(context.Background())
	if err != nil {
		panic(err)
	}
	for {
		line,_, err := reader.ReadLine()
		if err != nil {
			panic(err)
		}
		if string(line) == "close" { // 最后一次发送
			res, err := stream.CloseAndRecv()
			if err != nil {
				panic(err)
			}
			//打印所有服务端的返回
			fmt.Printf("clinet received:%s\n", res.GetRsp())
		}
		// 客户端调用远程grpc方法
		req := protos.EchoRequest{Req: string(line)}
		stream.Send(&req)
	}
}
