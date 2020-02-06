package TCPCode1

import (
	"net"
)

type TCPClient struct {
	conn net.Conn
}

var tcpClientMsgPrefix = "TCPCode1_TCPClient-"

func NewTCPClient() *TCPClient {
	obj := new(TCPClient)
	return obj
}

//调用示例:
//1.链式调用
//Link().Send().Response()
//2.顺序调用
//obj := NewTCPClient()
//obj.Link()
//obj.Send()
//obj.Response()

//TCP客户端-创建链接
func (this *TCPClient) Link() *TCPClient {
	localMsgPrefix := tcpClientMsgPrefix+"Link-"

	//存储创建链接产生的错误
	var linkErr error
	//创建链接
	this.conn,linkErr = net.Dial(TCPNetWork,TCPAddress)
	//若创建链接产生错误,则以异常抛出
	if linkErr!=nil {
		panic(localMsgPrefix+"netDialErr:"+linkErr.Error())
	}

	return this
}

//TCP客户端-发送消息
func (this *TCPClient) Send(writeByte []byte) *TCPClient{
	localMsgPrefix := tcpClientMsgPrefix+"Send-"

	//发送消息到服务端
	_,err := this.conn.Write(writeByte)
	//若发送产生错误,则以异常抛出
	if err!=nil {
		panic(localMsgPrefix+"connWrite:"+err.Error())
	}

	return this
}

//TCP客户端-获取响应
func (this *TCPClient) Response() string {
	localMsgPrefix := tcpClientMsgPrefix+"Response-"

	//存储响应字节
	var bytes [512]byte
	//读取服务端的响应消息
	n,err := this.conn.Read(bytes[:])
	//若读取产生错误,则以异常抛出
	if err!=nil {
		panic(localMsgPrefix+"connReadErr:"+err.Error())
	}

	//返回服务端的响应消息
	return string(bytes[:n])
}