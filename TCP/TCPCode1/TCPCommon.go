package TCPCode1

import (
	"Go-Note/util"
	"fmt"
)

type TCPCommon struct {

}

func NewTCPCommon() *TCPCommon {
	obj := new(TCPCommon)
	return obj
}

//执行入口
func (thisObj *TCPCommon) Do(params map[string]interface{}){
	//传参必须-标记
	doTag,_ := util.InterfaceToStr(params["doTag"])

	switch doTag {
	//TCP服务端开启
	case "serverStart":
		{
			NewTCPServer().Start()
		}
	//TCP客户端调用
	case "clientStart":
		{
			//客户端调用时,必须有发送消息存在
			sendMsg := ""
			if params["sendMsg"]!=nil {
				sendMsg = params["sendMsg"].(string)
			}

			//创建TCP客户端对象
			tcpClient := NewTCPClient()
			//最后关闭TCP链接
			defer tcpClient.Close()
			//发送消息并获取响应
			response := tcpClient.Link().Send([]byte(sendMsg)).Response()
			fmt.Println("TCPCode1_TCPClient_Response:",response)
		}
	}
}