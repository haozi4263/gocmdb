package main

import (
	"htgolang-201906/course/20191214/code/gocmdb/rpc/demo1/service"
	"log"
	"net"
	"net/rpc"
)


/*
rpc服务最多的优点就是 我们可以像使用本地函数一样使用 远程服务上的函数, 因此有几个关键点:

远程连接: 类似于我们的pkg
函数名称: 要表用的函数名称
函数参数: 这个需要符合RPC服务的调用签名, 及第一个参数是请求，第二个参数是响应
函数返回: rpc函数的返回是 连接异常信息, 真正的业务Response不能作为返回值

*/

var _ service.HelloService = (*HelloService(nil))

type HelloService struct{}

func (h *HelloService)Hello(request string, reply *string) error  {
	*reply = "hello" + request
	return nil
}

func main()  {
	// 把我们的对象注册成一个rpc的 receiver
	// 其中rpc.Register函数调用会将对象类型中所有满足RPC规则的对象方法注册为RPC函数，
	// 所有注册的方法会放在“HelloService”服务空间之下
	rpc.RegisterName("HelloService", new(HelloService))

	// 建立TCP链接
	listener,err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTcp error", err)
	}

	// 通过rpc.ServeConn函数在该TCP链接上为对方提供RPC服务。
	// 没Accept一个请求，就创建一个goroutie进行处理
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}

		// 前面都是tcp的知识, 到这个RPC就接管了
		// 因此 你可以认为 rpc 帮我们封装消息到函数调用的这个逻辑,
		// 提升了工作效率, 逻辑比较简洁，可以看看他代码
		go rpc.ServeConn(conn)
	}

}