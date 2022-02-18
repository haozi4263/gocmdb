package main

import (
	"fmt"
	"google.golang.org/grpc"
	"grpc/protos"
	"io"
	"net"
	"sync"
)

type EchoService struct {
	protos.UnimplementedEchoServiceServer
}

// bidirectional rpc, 双向数据流
func (es *EchoService)BiStreamEcho(stream protos.EchoService_BiStreamEchoServer) error {
	var wg sync.WaitGroup
	var ch = make(chan string)

	wg.Add(1)
	go func() { //接受请求
		defer wg.Done()
		for {
			req, err := stream.Recv()
			if err == io.EOF { // 结束
				close(ch)
				break
			}
			if err != nil {
				panic(err)
			}
			fmt.Println("received: ", req.GetReq())
			// todo 要发送给另外一个协程
			ch <- req.GetReq()
		}
	}()

	wg.Add(1)
	go func() { // 发送请求
		defer wg.Done()
		for v := range ch {
			if v == "" { // close ch后接受到空的管道需要退出循环
				break
			}
			err := stream.Send(&protos.EchoResponse{Rsp: v})
			for {
				if err != nil { //出现错误，继续发送
					fmt.Println("send err:", err)
					continue
				}
				break
			}

		}

	}()
	wg.Wait()
	return   nil
}

func main()  {
	rpcs := grpc.NewServer() //创建grpc server
	protos.RegisterEchoServiceServer(rpcs, &EchoService{}) //绑定
	lis, err := net.Listen("tcp",":8080")
	if err != nil {
		panic(err)
	}
	defer lis.Close()
	rpcs.Serve(lis)

}
