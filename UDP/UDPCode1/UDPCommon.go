package UDPCode1

import (
	"Go-Note/util"
	"fmt"
)

type UDPCommon struct {

}

func NewUDPCommon() *UDPCommon {
	obj := new(UDPCommon)
	return obj
}

//执行入口
func (thisObj *UDPCommon) Do(params map[string]interface{}) {
	//传参必须-标记
	doTag,_ := util.InterfaceToStr(params["doTag"])

	switch doTag {
	//UDP服务端开启
	case "serverStart":
		{
			NewUDPServer().Start()
		}
	//UDP客户端调用
	case "clientStart":
		{
			//客户端调用时,必须有发送消息存在
			sendMsg := ""
			if params["sendMsg"]!=nil {
				sendMsg = params["sendMsg"].(string)
			}

			//创建客户端对象
			client := NewUDPClient()
			//最后关闭链接
			defer client.Close()
			//发送消息并获取响应
			response := client.Link().Send([]byte(sendMsg)).Response()
			fmt.Println("UDPCode1_UDPClient_Response:",response)
		}
	}
}