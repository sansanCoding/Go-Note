package goHighLevel

import (
	"Go-Note/util"
	"fmt"
)

//初始化入口

type goHighLevelStruct struct {

}

//这里声明的变量名称以软件包名称开头,该方式只作笔记形式声明,不建议使用该方式命名!
var GoHighLevelStruct *goHighLevelStruct

func init(){
	GoHighLevelStruct = NewGoHighLevelStruct()
}

func NewGoHighLevelStruct() *goHighLevelStruct {
	return &goHighLevelStruct{

	}
}

//执行入口
func (thisObj *goHighLevelStruct) Do(params map[string]interface{}){
	//传参必须-方法名
	methodName := params["methodName"].(string)

	//CallMethodReflect调试:
	res,resOk := util.Helper.CallMethodReflect(thisObj,methodName,[]interface{}{})

	//输出结果:
	fmt.Println("---------------- 反射执行方法完毕,输出方法返回结果: ----------------")
	fmt.Println(res,resOk)
	for k,v := range res {
		fmt.Println("CallMethodReflectRes:",k,v)
	}
}