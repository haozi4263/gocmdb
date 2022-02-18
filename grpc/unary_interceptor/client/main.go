package main

import (
	"bufio"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc/protos"
	"os"
	"time"
)

// 拦截器功能 method:远程调用方法  cc:grpc链接信息 invoker:远程调用方法
func UnaryInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error{
	// 实现连接器功能 统计时间
	start := time.Now()
	err := invoker(ctx, method, req, reply, cc, opts...)

	end := time.Now()
	fmt.Printf("UnaryInterceptor rpc-name: %s, start: %s end: %s, req:%v rsp:%v\n",
		method, start.Format(time.RFC822), end.Format(time.RFC822),  req, reply)
	return err
}

func UnaryInterceptorEx(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error{
	// 实现连接器功能 统计时间
	start := time.Now()
	err := invoker(ctx, method, req, reply, cc, opts...)

	end := time.Now()
	fmt.Printf("UnaryInterceptorEx rpc-name: %s, start: %s end: %s, req:%v rsp:%v\n",
		method, start.Format(time.RFC822), end.Format(time.RFC822),  req, reply)
	return err
}

func main()  {
	//链接grpc服务
	//cli, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure(), grpc.WithUnaryInterceptor(UnaryInterceptor)) // 使用不安全的链接, 一元拦截器
	// 多个拦截器，执行顺序。 先入后执行 后入先执行
	cli, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure(), grpc.WithChainUnaryInterceptor(UnaryInterceptor, UnaryInterceptorEx)) //  一元拦截器

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
