package goBase

import (
	"Go-Note/util"
	"fmt"
	"reflect"
)

//######################################################################################################################
//GO-自定义类型
//	参考文章地址:
//		http://topgoer.com/go%E5%9F%BA%E7%A1%80/%E7%BB%93%E6%9E%84%E4%BD%93.html
//----------------------------------------------------------------------------------------------------------------------
//
//	1.使用type关键字来定义自定义类型。
//	  自定义类型是定义了一个全新的类型。可以基于内置的基本类型定义，也可以通过struct定义。
//
//	2.类型别名是Go1.9版本添加的新功能。
//	  类型别名规定：TypeAlias只是Type的别名，本质上TypeAlias与Type是同一个类型。
//		示例:
//		//类型定义
//		type NewInt int
//
//		//类型别名
//		type MyInt = int
//
//		func main() {
//		   var a NewInt
//		   var b MyInt
//
//		   fmt.Printf("type of a:%T\n", a) //type of a:main.NewInt
//		   fmt.Printf("type of b:%T\n", b) //type of b:int
//		}
//		结果显示a的类型是main.NewInt，表示main包下定义的NewInt类型。b的类型是int。MyInt类型只会在代码中存在，编译完成时并不会有MyInt类型。
//######################################################################################################################

type customType struct {

}

var CustomType *customType

func init(){

}

func NewCustomType() *customType {
	return &customType{

	}
}

//执行入口
func (thisObj *customType) Do(params map[string]interface{}){
	//传参必须-方法名
	methodName := params["methodName"].(string)

	//CallMethodReflect调试:
	res,resOk := util.Helper.CallMethodReflect(thisObj,methodName,[]interface{}{})

	//输出结果:
	fmt.Println(res,resOk)
	for k,v := range res {
		fmt.Println("CallMethodReflectRes:",k,v)
	}
}

//入口
//命令行-输入:{"optTag":"CustomType","optParams":{"methodName":"Index"}}
func (thisObj *customType) Index(){

	type MyInt int
	var testInt MyInt
	testInt = 1
	fmt.Println("testInt:",testInt)

	type MyStr string
	var testStr MyStr
	testStr = "is string"
	fmt.Println("testStr:",testStr)
}

//类型别名
//命令行-输入:{"optTag":"CustomType","optParams":{"methodName":"TypeAlias"}}
func (thisObj *customType) TypeAlias(){

	//类型定义
	type NewInt int
	//类型别名
	type MyInt = int
	var testNewInt NewInt
	var testMyInt MyInt
	testNewInt = 1
	testMyInt = 1
	//输出调试:
	fmt.Println("testNewInt:",testNewInt,reflect.TypeOf(testNewInt))
	fmt.Println("testMyInt:",testMyInt,reflect.TypeOf(testMyInt))
	//输出结果:
	//testNewInt: 1 goBase.NewInt
	//testMyInt: 1 int
}
