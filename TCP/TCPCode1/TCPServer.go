package TCPCode1

import (
	"Go-Note/config"
	"bufio"
	"fmt"
	"io"
	"net"
	"reflect"
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

	//开启TCP协议-地址(包含ip+端口)-监听
	listen,listenErr := net.Listen(config.TCP.NetWork,config.TCP.Address)
	//若监听产生错误
	if listenErr!=nil {
		panic(localMsgPrefix+"netListen["+config.TCP.NetWork+","+config.TCP.Address+"]Err:"+listenErr.Error())
	}

	//其实下面执行的死循环读取,但是从资源开关上讲,有开还是得有关的
	defer listen.Close()

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
	//localMsgPrefix := tcpServerMsgPrefix+"serverAccept-"

	//处理完毕后就关闭链接
	defer conn.Close()

	//缓存读取
	readConn := bufio.NewReader(conn)

	for {
		//以128字节读取数据
		var bytes [128]byte
		n,err := readConn.Read(bytes[:])
		//错误检测
		if this.serverAcceptReadErrorCheck(err) {
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

//TCP服务端-处理每一个链接-读取错误检测处理
//@return bool true 代表是有错误(也有部分不是错误,如读到输入流结尾,客户端端断开链接),false 代表是无错误
func (this *TCPServer) serverAcceptReadErrorCheck(err error) bool {
	localMsgPrefix := tcpServerMsgPrefix+"serverAcceptReadErrorCheck-"

	//若是读到结尾,则代表没有了
	if err==io.EOF {
		go this.errChannelAdd(localMsgPrefix+"readConnReadEOF!")
		return true
	//若是其他错误
	}else if err!=nil {
		go this.errChannelAdd(localMsgPrefix+"readConnReadErr[ErrType:"+fmt.Sprint(reflect.TypeOf(err))+"]:"+err.Error())
		return true
	}

	return false
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
