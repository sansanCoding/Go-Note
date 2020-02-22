package goBase

import "fmt"

//流程控制

type lckz struct {
	Base
}

var Lckz *lckz

func init(){
	Lckz = NewLckz()
}

func NewLckz() *lckz {
	return &lckz{

	}
}

//switch流程控制
//命令行-输入:{"optTag":"Lckz","optParams":{"methodName":"Switch"}}
func (thisObj *lckz) Switch(){
	test := 0
	switch test {
	case 0:
		fmt.Println("this is 0")
		fallthrough	//fallthrough强制执行后面的case代码
	case 1:
		fmt.Println("this is 1")
	case 2:
		fmt.Println("this is 2")
	default:
		fmt.Println("this is default")
	}

	//输出结果:
	//this is 0
	//this is 1
}
