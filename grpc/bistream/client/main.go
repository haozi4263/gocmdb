package main

import (
	"bufio"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc/protos"
	"io"
	"os"
	"sync"
)

func main() {
	//链接grpc服务
	cli, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure()) // 使用不安全的链接
	if err != nil {
		panic(err)
	}

	defer cli.Close()
	c := protos.NewEchoServiceClient(cli)
	reader := bufio.NewReader(os.Stdin)
	var wg sync.WaitGroup
	for {
		stream, err := c.BiStreamEcho(context.Background())
		if err != nil {
			panic(err)
		}
		wg.Add(1)
		go func() { // 发送数据
			defer wg.Done()
			for {
				// 从控制台读取数据
				line, _, err := reader.ReadLine()
				if err != nil {
					panic(err)
				}
				if string(line) == "close"{
					if err = stream.CloseSend(); err != nil {
						panic(err)
					}
					break
				}
				req := protos.EchoRequest{Req: string(line)}
				stream.Send(&req)
			}
		}()

		wg.Add(1)
		go func() { //接受
			defer wg.Done()
			for {
				res, err := stream.Recv()
				if err == io.EOF{

				}
				if err != nil { //接受数据出错，继续接受
					continue
				}
				fmt.Println("received:", res.GetRsp() )
			}
		}()

		wg.Wait()
	}
}
