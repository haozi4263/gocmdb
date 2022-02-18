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
	for {
		line,_, err := reader.ReadLine()
		if err != nil {
			panic(err)
		}
		// 客户端调用远程grpc方法
		res ,err := c.UnaryEcho(context.Background(), &protos.EchoRequest{
			Req: string(line),
		})
		if err != nil {
			panic(err)
		}
		fmt.Println(res.GetRsp())
	}
}
