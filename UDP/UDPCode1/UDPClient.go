package UDPCode1

import (
	"Go-Note/config"
	"net"
)

var udpClientMsgPrefix = "UDPCode1_UDPClient-"

type UDPClient struct {
	conn *net.UDPConn
}

func NewUDPClient() *UDPClient {
	obj := new(UDPClient)
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

//UDP客户端-创建链接
func (this *UDPClient) Link() *UDPClient {
	localMsgPrefix := udpClientMsgPrefix+"Link-"

	//存储创建链接产生的错误
	var connErr error
	//创建链接
	this.conn,connErr = net.DialUDP(config.UDP.NetWork,nil,&net.UDPAddr{
		IP:net.ParseIP(config.UDP.IP),
		Port:config.UDP.Port,
	})
	//若创建链接失败,则以异常抛出
	if connErr!=nil {
		panic(localMsgPrefix+"netDialUDPErr:"+connErr.Error())
	}

	return this
}

//UDP客户端-发送消息
func (this *UDPClient) Send(writeByte []byte) *UDPClient {
	localMsgPrefix := udpClientMsgPrefix+"Send-"

	//发送消息到服务端
	_,err := this.conn.Write(writeByte)
	if err!=nil {
		panic(localMsgPrefix+"connWriteErr:"+err.Error())
	}

	return this
}

//UDP客户端-获取响应
func (this *UDPClient) Response() map[string]interface{} {
	localMsgPrefix := udpClientMsgPrefix+"Response-"

	//如果以下无法接收数据换成如下注释的方式获取试试
	// data := make([]byte, 4096)
	// n, remoteAddr, err := socket.ReadFromUDP(data) // 接收数据
	//存储响应字节
	var data [4096]byte
	//读取UDP服务端的响应消息
	n,addr,err := this.conn.ReadFromUDP(data[:])
	if err!=nil {
		panic(localMsgPrefix+"connReadErr:"+err.Error())
	}

	//调试输出:
	//fmt.Printf("recv:%v addr:%v count:%v\n", string(data[:n]), addr, n)

	return map[string]interface{}{
		"recv":string(data[:n]),
		"addr":addr.String(),
		"count":n,
	}
}

//UDP客户端-关闭链接
func (this *UDPClient) Close() {
	localMsgPrefix := udpClientMsgPrefix+"Close-"

	//关闭链接
	err := this.conn.Close()
	//若读取产生错误,则以异常抛出
	if err!=nil {
		panic(localMsgPrefix+"connCloseErr:"+err.Error())
	}
}
