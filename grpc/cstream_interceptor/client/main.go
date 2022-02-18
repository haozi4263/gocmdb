package main

import (
	"bufio"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc/protos"
	"os"
)

// 实现grpc.ClientStream 接口
type clientStream struct {
	grpc.ClientStream
}

func NewClientStream(s grpc.ClientStream) grpc.ClientStream  {
	return &clientStream{s}
}

// 覆盖SendMsg方法
func (c *clientStream)SendMsg(m interface{}) error  {
	fmt.Println("SendMsg: ", m)
	return c.ClientStream.SendMsg(m)
}
// 覆盖RecvMsg方法
func (c *clientStream)RecvMsg(m interface{}) error {
	fmt.Println("RecvMsg: ", m)
	return c.ClientStream.RecvMsg(m)
}

// desc:stream描述 cc:链接信息 调用方法 streamer：客户端流多个  grpc.ClientStream：服务端流
func CStreamInterceptor(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error){
	stream, err := streamer(ctx,desc, cc, method, opts...)
	if err != nil {
		return nil, err
	}
	return NewClientStream(stream), err
}


func main()  {
	//链接grpc服务
	cli, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure(), grpc.WithStreamInterceptor(CStreamInterceptor)) // 使用不安全的链接
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
