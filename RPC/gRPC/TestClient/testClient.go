package TestClient

import (
	"Go-Note/RPC/gRPC/Test"
	"Go-Note/RPC/gRPC/TestServer"
	"context"
	"fmt"
	"google.golang.org/grpc"
)

//消息前缀
var msgPrefix = "TestClient-"

//主入口-客户端发起请求
func ClientRequest(){
	//本地消息前缀
	localMsgPrefix := msgPrefix+"ClientInit-"


	//1.grpc.Dial("127.0.0.1:"+TestServer.GRPCPort,多个配置参数...):
	//	客户端创建连接的时候默认必须使用加密传输，否则会直接报错如下: grpc: no transport security set (use grpc.WithInsecure() explicitly or set credentials)
	//	可使用grpc.WithInsecure()避免加密传输
	//2.开启 TLS/SSL 加密
	//	TLS 是一种常见的端对端安全传输协议
	//	grpc 可以使用 credentials.TransportCredentials 结构来方便的开启安全传输
	// 参考文章地址：
	// https://blog.csdn.net/DAGU131/article/details/105815620


	//创建连接
	conn,connErr := grpc.Dial("127.0.0.1:"+TestServer.GRPCPort,grpc.WithInsecure())
	if connErr!=nil {
		panic(localMsgPrefix+"grpcDialErr:"+connErr.Error())
	}
	defer conn.Close()

	//获取UserService客户端
	client := Test.NewUserServiceClient(conn)

	//@todo 如果想发起多个并行请求,参考地址:https://studygolang.com/articles/4370

	//准备请求参数
	userRequest := &Test.UserRequest{
		UserID: 111,
	}

	//这一步就是客户端调用远程服务端的方法
	userResponse,userResponseErr := client.GetUser(context.Background(),userRequest)
	if userResponseErr!=nil {
		panic(localMsgPrefix+"clientGetUserErr:"+userResponseErr.Error())
	}

	//输出结果:
	fmt.Println("userResponse.GetUserID:",userResponse.GetUserID())
	fmt.Println("userResponse.GetUserName:",userResponse.GetUserName())

	//输出结果:
	//userResponse.GetUserID: 111
	//userResponse.GetUserName: 111:test
}