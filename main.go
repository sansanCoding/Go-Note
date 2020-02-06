package main

import (
	"Go-Note/TCP/TCPCode1"
	"Go-Note/UDP/UDPCode1"
	"Go-Note/util"
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main(){

	//提示语
	fmt.Println("请输入对应调试代码编号,输入exit则退出:")

	//开启命令行输入
	stdInReader := bufio.NewReader(os.Stdin)
	//死循环处理每一个命令行输入内容
	for {
		//监听命令行输入内容
		input,inputErr := stdInReader.ReadString('\n')
		//若输入产生错误,则直接抛出异常停止继续执行
		if inputErr!=nil {
			panic("stdInReader.ReadString.Error:"+inputErr.Error())
		}
		input = strings.Trim(input,"\r\n")
		if input=="" {
			fmt.Println("请输入对应调试代码编号!")
			continue
		}
		if input=="exit" {
			fmt.Println("已退出调试")
			return
		}

		//调试输出:
		//fmt.Println("input:",input)
		go inputCall(input)
	}

}

//输入调用
func inputCall(input string){
	//捕获异常
	defer func(){
		if err := recover();err != nil {
			fmt.Println("inputCallPanicErr:",err)
		}
	}()

	//命令行输入的内容类似如:{"optTag":"TCPCode1","optParams":{"doTag":"serverStart"}}
	var inputJson map[string]interface{}
	jsonErr := json.Unmarshal([]byte(input),&inputJson)
	if jsonErr!=nil {
		panic("input_jsonUnmarshal_err:"+jsonErr.Error())
	}

	//调试输出:
	//fmt.Println("inputJson:",inputJson)
	//return
	//输出结果:
	//inputJson: map[optParams:map[doTag:serverStart] optTag:TCPCode1]

	//必须-操作标记
	optTag,_ 		:= util.InterfaceToStr(inputJson["optTag"])
	//必须-操作参数
	optParams 		:= inputJson["optParams"].(map[string]interface{})

	switch optTag {
	//第1版-简易版TCP代码
	case "TCPCode1":
		{
			//根据操作参数执行不同的操作
			//	1.TCPCode1的TCP服务端开启命令-命令行输入:{"optTag":"TCPCode1","optParams":{"doTag":"serverStart"}}
			//	2.TCPCode1的TCP客户端调用命令-命令行输入:{"optTag":"TCPCode1","optParams":{"doTag":"clientStart","sendMsg":"Test"}}
			TCPCode1.NewTCPCommon().Do(optParams)
		}
	//第1版-简易版UDP代码
	case "UDPCode1":
		{
			//根据操作参数执行不同的操作
			//	1.UDPCode1的TCP服务端开启命令-命令行输入:{"optTag":"UDPCode1","optParams":{"doTag":"serverStart"}}
			//	2.UDPCode1的TCP客户端调用命令-命令行输入:{"optTag":"UDPCode1","optParams":{"doTag":"clientStart","sendMsg":"Test"}}
			UDPCode1.NewUDPCommon().Do(optParams)
		}
	}
}