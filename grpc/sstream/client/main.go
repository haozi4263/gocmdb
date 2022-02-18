package main

import (
	"bufio"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc/protos"
	"io"
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
	for {
		line,_, err := reader.ReadLine()
		if err != nil {
			panic(err)
		}
		// 客户端调用远程grpc方法
		req := protos.EchoRequest{Req: string(line)}
		stream, err := c.SStreamEcho(context.Background(), &req)
		if err != nil {
			panic(err)
		}

		for {
			rsp, err := stream.Recv()
			if err == io.EOF { //服务端返回最后一个值
				break
			}
			if err != nil {
				continue
			}
			fmt.Println("received: " + rsp.GetRsp())
		}
	}
}
