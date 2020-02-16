package UDPCode1

import (
	"Go-Note/config"
	"fmt"
	"net"
	"strconv"
	"sync"
)

var udpServerMsgPrefix = "UDPCode1_UDPServer-"
var udpServerIsDoing = false

type UDPServer struct {
	doOnce sync.Once
	errChannel chan interface{}
}

func NewUDPServer() *UDPServer {
	obj := new(UDPServer)
	obj.errChannel = make(chan interface{})
	return obj
}

//主入口-开始
func (thisObj *UDPServer) Start(){

	//若已经正在运行了,则不能重复请求
	if udpServerIsDoing {
		fmt.Println(udpServerMsgPrefix+"Listen_Is_Exists!")
		return
	}

	//并发调用时只会被执行一次
	thisObj.doOnce.Do(func(){

		//1.异步开启错误通道监听(该步必须异步执行,否则使用通道会报错)
		go thisObj.errChannelListen()

		//2.准备介入监听
		fmt.Println(udpServerMsgPrefix+"Listen...")

		//3.开启服务端监听(监听是阻塞的)
		thisObj.serverListen()
	})

}

//UDP服务端-开启监听
func (thisObj *UDPServer) serverListen(){
	localMsgPrefix := udpServerMsgPrefix+"serverListen-"

	//开启UDP协议-IP+端口-监听
	listen,listenErr := net.ListenUDP(config.UDP.NetWork,&net.UDPAddr{
		IP	: net.ParseIP(config.UDP.IP),
		Port: config.UDP.Port,
	})
	//若监听产生错误
	if listenErr!=nil {
		panic(localMsgPrefix+"netListenUDP["+config.UDP.NetWork+","+config.UDP.IP+","+strconv.Itoa(config.UDP.Port)+"]Err:"+listenErr.Error())
	}

	//其实下面执行的死循环读取,但是从资源开关上讲,有开还是得有关的
	defer listen.Close()

	//若监听执行成功,则下次不能重复触发该监听
	udpServerIsDoing = true

	//死循环接收每一个客户端请求
	for {
		var data [1024]byte
		//接收数据
		n,addr,readErr := listen.ReadFromUDP(data[:])
		if readErr!=nil {
			go thisObj.errChannelAdd(localMsgPrefix+"listenReadFromUDPErr:"+readErr.Error())
			continue
		}

		//调试输出:
		//fmt.Printf(localMsgPrefix+"data:%v addr:%v count:%v\n", string(data[:n]), addr, n)

		//发送数据
		_,writeErr := listen.WriteToUDP(data[:n],addr)
		if writeErr!=nil {
			go thisObj.errChannelAdd(localMsgPrefix+"listenWriteToUDPErr:"+writeErr.Error())
			continue
		}
	}
}

//UDP服务端-错误通道监听
func (thisObj *UDPServer) errChannelListen(){
	localMsgPrefix := udpServerMsgPrefix+"errChannelListen-"

	//异常捕获
	defer func(){
		if err := recover(); err != nil {
			//暂时先以这种方式直接输出
			fmt.Println(localMsgPrefix+"panicErr:",err)
		}
	}()

	//获取错误通道里的内容
	err := <- thisObj.errChannel

	//暂时先以这种方式直接输出
	fmt.Println(localMsgPrefix+"err:",err)
}

//UDP服务端-添加错误信息到错误通道
func (thisObj *UDPServer) errChannelAdd(errInfo interface{}){
	localMsgPrefix := udpServerMsgPrefix+"errChannelAdd-"

	//异常捕获
	defer func(){
		if err := recover(); err != nil {
			//暂时先以这种方式直接输出
			fmt.Println(localMsgPrefix+"panicErr:",err)
		}
	}()

	//添加错误信息到错误通道
	thisObj.errChannel <- errInfo
}
