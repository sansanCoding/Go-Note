package main

import (
	"Go-Note/RPC/gRPC/TestClient"
	"Go-Note/RPC/gRPC/TestServer"
	"time"
)

/*
	//@todo 参考文章地址:https://doc.oschina.net/grpc?t=58008

	gRPC使用的流程说明
	@todo 1.定义服务
	@todo 	第一步是定义一个服务：一个 RPC 服务通过参数和返回类型来指定可以远程调用的方法。
	@todo 	gRPC默认使用 protocol buffers(即使用proto文件写入相关服务定义，如请求参数定义，响应参数定义等);
	@todo 	创建一个test.proto文件并写入相关RPC请求与响应的参数定义。

	@todo 2.将定义的服务(即写好的proto文件)生成go语言的代码文件
	@todo 	protoc --go_out=plugins=grpc:. test.proto

	@todo 3.go引用gRPC包
	@todo go get -u -v google.golang.org/grpc

	@todo 4.准备服务实现
	@todo	参考TestServer/testServer.go代码
		//服务实现
		type UserService struct {

		}
		//获取用户信息
		func (us *UserService) GetUser(ctx context.Context,request *Test.UserRequest) (response *Test.UserResponse,err error){
			return &Test.UserResponse{
				UserID:request.UserID,
				UserName:strconv.Itoa(int(request.UserID))+":test",
			},nil
		}

	@todo 5.服务端实现：准备服务端代码和监听TCP协议和地址端口。
	@todo	需要提供一个 gRPC 服务的另一个主要功能是让这个服务实在在网络上可用。
	@todo 	参考TestServer/testServer.go代码

	@todo 6.创建客户端请求
	@todo	写一个客户端，客户端的 gRPC 非常简单。在这一步，我们将用生成的代码写一个简单的客户程序来访问服务端。
	@todo	参考TestClient/testClient.go代码
 */

//@todo 直接CD进入到Go-Note/RPC/gRPC目录下,使用go run testMain.go介入测试!

func main(){
	//先启动服务端监听
	go TestServer.ServerListen()

	//5秒后由客户端发起请求
	time.Sleep(5*time.Second)

	//发起客户端请求
	TestClient.ClientRequest()
}