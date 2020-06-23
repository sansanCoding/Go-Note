package TestServer

import (
	"Go-Note/RPC/gRPC/Test"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"strconv"
)

//消息前缀
var msgPrefix = "TestServer-"

//监听端口
var GRPCPort = "52052"

//主入口-服务端监听
func ServerListen(){
	//本地消息前缀
	localMsgPrefix := msgPrefix+"ServerInit-"

	//监听tcp协议+地址端口
	listen,listenErr := net.Listen("tcp",":"+GRPCPort)
	if listenErr!=nil {
		panic(localMsgPrefix+"netListenErr:"+listenErr.Error())
	}

	//创建grpc服务端
	grpcServer := grpc.NewServer()
	//注册UserService服务
	Test.RegisterUserServiceServer(grpcServer,&UserService{})

	//接收监听的消息
	serveErr := grpcServer.Serve(listen)
	if serveErr!=nil {
		panic(localMsgPrefix+"grpcServerServeErr:"+serveErr.Error())
	}

	fmt.Print(localMsgPrefix+"listen_GRPCPort:"+GRPCPort)
}

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