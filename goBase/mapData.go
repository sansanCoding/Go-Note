package goBase

import (
	"Go-Note/util"
	"fmt"
	"reflect"
)

//######################################################################################################################
//GO-map
//	参考文章地址:
//		http://topgoer.com/go%E5%9F%BA%E7%A1%80/Map.html
//
//----------------------------------------------------------------------------------------------------------------------
//@todo map是一种无序的基于key-value的数据结构，Go语言中的map是引用类型，必须初始化才能使用。
//----------------------------------------------------------------------------------------------------------------------
//
//	1.map类型的变量默认初始值为nil，需要使用make()函数来分配内存。语法为：
//		make(map[KeyType]ValueType, [cap])
//	其中cap表示map的容量，该参数虽然不是必须的，但是我们应该在初始化map的时候就为其指定一个合适的容量。(这个视具体情况而定,可加可不加)
//
//######################################################################################################################

type mapData struct {

}

var MapData *mapData

func init(){
	MapData = NewMapData()
}

func NewMapData() *mapData {
	return &mapData{

	}
}

//执行入口
func (thisObj *mapData) Do(params map[string]interface{}){
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

//入门
//命令行-输入:{"optTag":"MapData","optParams":{"methodName":"Index"}}
func (thisObj *mapData) Index(){
	//声明,未初始化(即未分配内存地址)
	var testData map[string]int
	//未分配内存地址,赋值会报错
	//testData["test"] = 1
	//调试输出:
	fmt.Println("testData:",testData,reflect.TypeOf(testData))
	//输出结果:
	//testData: map[] map[string]int

	//声明并初始化赋值(即分配内存地址)
	testData1 := make(map[string]int)
	//已分配内存地址,可被赋值
	testData1["test"] = 1
	//调试输出:
	fmt.Println("testData1:",testData1,reflect.TypeOf(testData1))

	//声明并初始化赋值(即分配内存地址)
	testData2 := map[string]int{}
	//已分配内存地址,可被赋值
	testData2["test"] = 2
	//调试输出:
	fmt.Println("testData2:",testData2,reflect.TypeOf(testData2))

	//声明并初始化赋值(即分配内存地址)
	var testData3 = map[string]int{}
	//已分配内存地址,可被赋值
	testData3["test"] = 3
	//调试输出:
	fmt.Println("testData3:",testData3,reflect.TypeOf(testData3))
}

//查找key是否存在
//命令行-输入:{"optTag":"MapData","optParams":{"methodName":"FindKeyExists"}}
func (thisObj *mapData) FindKeyExists(){
	testData := make(map[string]int)
	testData["test"] = 1
	//如果key存在ok为true,v为对应的值；不存在ok为false,v为值类型的零值
	findVal,ok := testData["test"]
	if ok {
		fmt.Println("findVal is exists!the value is ",findVal)
	}else{
		fmt.Println("findVal is not exists!")
	}
}

//Map遍历
//命令行-输入:{"optTag":"MapData","optParams":{"methodName":"ForRange"}}
func (thisObj *mapData) ForRange(){
	testData := make(map[string]int)
	testData["test1"] = 1
	testData["test2"] = 2
	testData["test3"] = 3
	testData["test4"] = 4

	//遍历key和value
	for key,value := range testData {
		fmt.Println("key:",key,"value:",value)
	}
	fmt.Println("------------------ MapForRange Key And Value ------------------")

	//只遍历key,不需要value(如果在不熟悉这种写法的意思,会误认为只给一个变量赋值的是value而不是key,即使名称叫做key)
	for key := range testData {
		fmt.Println("key:",key)
	}
	fmt.Println("------------------ MapForRange Key 1 ------------------")

	//另一种只遍历key,不需要value的写法(这种写法,更加明晰是使用的key还是使用的value)
	for key,_ := range testData {
		fmt.Println("key:",key)
	}
	fmt.Println("------------------ MapForRange Key 2 ------------------")
}

//Map删除元素
//命令行-输入:{"optTag":"MapData","optParams":{"methodName":"Delete"}}
func (thisObj *mapData) Delete(){
	testData := make(map[string]int)
	testData["test1"] = 1
	testData["test2"] = 2
	testData["test3"] = 3

	//删除指定map中的某个key(即阐述了该元素)
	//	由于map是引用类型,删除是针对该map生效
	delete(testData,"test1")

	for key,value := range testData {
		fmt.Println("key:",key,"value:",value)
	}
}
