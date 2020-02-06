package TCPCode1

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"sync"
)

var tcpServerMsgPrefix = "TCPCode1_TCPServer-"
var tcpServerIsDoing = false

type TCPServer struct {
	doOnce sync.Once
	errChannel chan interface{}
}

func NewTCPServer() *TCPServer {
	obj := new(TCPServer)
	obj.errChannel = make(chan interface{})
	return obj
}

//主入口-开始
func (this *TCPServer) Start(){

	//若已经正在运行了,则不能重复请求
	if tcpServerIsDoing {
		fmt.Println(tcpServerMsgPrefix+"Listen_Is_Exists!")
		return
	}

	//并发调用时只会被执行一次
	this.doOnce.Do(func(){

		//1.异步开启错误通道监听(该步必须异步执行,否则使用通道会报错)
		go this.errChannelListen()

		//2.准备介入监听
		fmt.Println(tcpServerMsgPrefix+"Listen...")

		//3.开启服务端监听(监听是阻塞的)
		this.serverListen()
	})

}

//TCP服务端-开启监听
func (this *TCPServer) serverListen(){
	localMsgPrefix := tcpServerMsgPrefix+"serverListen-"

	//开启TCP协议-地址+端口-监听
	listen,listenErr := net.Listen(TCPNetWork,TCPAddress)
	//若监听产生错误
	if listenErr!=nil {
		panic(localMsgPrefix+"netListen["+TCPNetWork+","+TCPAddress+"]Err:"+listenErr.Error())
	}

	//若监听执行成功,则下次不能重复触发该监听
	tcpServerIsDoing = true

	//死循环获取每一个客户端请求
	for {
		//建立链接
		conn,connErr := listen.Accept()
		if connErr!=nil {
			go this.errChannelAdd(localMsgPrefix+"listenAcceptErr:"+connErr.Error())
			continue
		}
		//调用每一个链接处理
		go this.serverAccept(conn)
	}
}

//TCP服务端-处理每一个链接
func (this *TCPServer) serverAccept(conn net.Conn){
	localMsgPrefix := tcpServerMsgPrefix+"serverAccept-"

	//处理完毕后就关闭链接
	defer conn.Close()

	for {
		//缓存读取
		readConn := bufio.NewReader(conn)
		//以128字节读取数据
		var bytes [128]byte
		n,err := readConn.Read(bytes[:])
		//这个判断一般针对文件有用
		if err==io.EOF {
			go this.errChannelAdd(localMsgPrefix+"readConnReadEOF!")
			break
		//若是其他错误
		}else if err!=nil {
			go this.errChannelAdd(localMsgPrefix+"readConnReadErr:"+err.Error())
			break
		}
		//将最开头获取到指定n的位置字节数据(包含n位置的字节)
		revStr := string(bytes[:n])
		//调试输出
		//fmt.Println("客户端发送的数据是:",revStr)
		//输出给客户端
		conn.Write([]byte(revStr))
	}
}

//TCP服务端-错误通道监听
func (this *TCPServer) errChannelListen(){
	localMsgPrefix := tcpServerMsgPrefix+"errChannelListen-"

	//异常捕获
	defer func(){
		if err := recover(); err != nil {
			//暂时先以这种方式直接输出
			fmt.Println(localMsgPrefix+"panicErr:",err)
		}
	}()

	//获取错误通道里的内容
	err := <- this.errChannel

	//暂时先以这种方式直接输出
	fmt.Println(localMsgPrefix+"err:",err)
}

//TCP服务端-添加错误信息到错误通道
func (this *TCPServer) errChannelAdd(errInfo interface{}){
	localMsgPrefix := tcpServerMsgPrefix+"errChannelAdd-"

	//异常捕获
	defer func(){
		if err := recover(); err != nil {
			//暂时先以这种方式直接输出
			fmt.Println(localMsgPrefix+"panicErr:",err)
		}
	}()

	//添加错误信息到错误通道
	this.errChannel <- errInfo
}
