package main

import (
	"Go-Note/TCP/TCPCode1"
	"Go-Note/UDP/UDPCode1"
	"Go-Note/goBase"
	"Go-Note/goHighLevel"
	"Go-Note/goTestCode"
	"Go-Note/util"
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"runtime/debug"
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
			//异常信息输出
			fmt.Println("inputCallPanicErr:",err)
			//异常时输出错误栈
			fmt.Println(string(debug.Stack()))
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

	//传参必须-操作标记
	optTag,_ 		:= util.InterfaceToStr(inputJson["optTag"])
	//传参必须-操作参数(根据操作参数执行不同的操作)
	optParams 		:= inputJson["optParams"].(map[string]interface{})

	switch optTag {
	//第1版-简易版TCP代码
	case "TCPCode1":
		{
			//	1.TCPCode1的TCP服务端开启命令-命令行输入:{"optTag":"TCPCode1","optParams":{"doTag":"serverStart"}}
			//	2.TCPCode1的TCP客户端调用命令-命令行输入:{"optTag":"TCPCode1","optParams":{"doTag":"clientStart","sendMsg":"Test"}}
			TCPCode1.NewTCPCommon().Do(optParams)
		}
	//第1版-简易版UDP代码
	case "UDPCode1":
		{
			//	1.UDPCode1的UDP服务端开启命令-命令行输入:{"optTag":"UDPCode1","optParams":{"doTag":"serverStart"}}
			//	2.UDPCode1的UDP客户端调用命令-命令行输入:{"optTag":"UDPCode1","optParams":{"doTag":"clientStart","sendMsg":"Test"}}
			UDPCode1.NewUDPCommon().Do(optParams)
		}
	//字符串高级
	case "StringAdvanced":
		{
			//命令行输入:{"optTag":"StringAdvanced","optParams":{"methodName":"对应的方法名称"}}
			goBase.StringAdvanced.Do(optParams)
		}
	//数组Array
	case "Array":
		{
			//命令行输入:{"optTag":"Array","optParams":{"methodName":"对应的方法名称"}}
			goBase.Array.Do(optParams)
		}
	//切片Slice
	case "Slice":
		{
			//命令行输入:{"optTag":"Slice","optParams":{"methodName":"对应的方法名称"}}
			goBase.Slice.Do(optParams)
		}
	//指针
	case "指针":
		{
			//小测试-命令行输入:{"optTag":"指针","optParams":{"doTag":"exam"}}
			goBase.Pointer.Do(optParams)
		}
	//map数据
	case "MapData":
		{
			//命令行输入:{"optTag":"MapData","optParams":{"methodName":"对应的方法名称"}}
			goBase.MapData.Do(optParams)
		}
	//自定义类型
	case "CustomType":
		{
			//命令行输入:{"optTag":"CustomType","optParams":{"methodName":"对应的方法名称"}}
			goBase.CustomType.Do(optParams)
		}
	//struct结构体
	case "StructType":
		{
			//命令行输入:{"optTag":"StructType","optParams":{"methodName":"对应的方法名称"}}
			goBase.StructType.Do(optParams)
		}
	//流程控制
	case "Lckz":
		{
			//命令行输入:{"optTag":"Lckz","optParams":{"methodName":"对应的方法名称"}}
			goBase.Lckz.Do(goBase.Lckz,optParams)
		}
	//util
	case "Util":
		{
			//命令行输入:{"optTag":"Util","optParams":{"methodName":"对应的方法名称"}}
			goTestCode.UtilTestCode.Do(optParams)
		}
	//单元测试
	case "FuncTest":
		{
			//命令行输入:{"optTag":"FuncTest","optParams":{"methodName":"对应的方法名称"}}
			//具体代码参考:goBase/函数-单元测试-funcTest.go 和 函数-压力测试-goTest.go
		}
	//go高级
	case "GoHighLevel":
		{
			//命令行输入:{"optTag":"GoHighLevel","optParams":{"methodName":"对应的方法名称"}}
			goHighLevel.GoHighLevelStruct.Do(optParams)
		}
	}
}
