package goBase

import (
	"Go-Note/util"
	"fmt"
	"strconv"
)


//######################################################################################################################
//GO-数组Array
//	GO-关于数组Array的说明:
//	  参考文章地址:
//	  http://topgoer.com/go%E5%9F%BA%E7%A1%80/%E6%95%B0%E7%BB%84Array.html
//
// 	  1. 数组：是同一种数据类型的固定长度的序列。
//    2. 数组定义：var a [len]int，比如：var a [5]int，数组长度必须是常量，且是类型的组成部分。一旦定义，长度不能变。
//    3. 长度是数组类型的一部分，因此，var a[5] int和var a[10]int是不同的类型。
//    4. 数组可以通过下标进行访问，下标是从0开始，最后一个元素下标是：len-1
//    for i := 0; i < len(a); i++ {
//    }
//    for index, v := range a {
//    }
//    5. 访问越界，如果下标在数组合法范围之外，则触发访问越界，会panic
//@todo 6. 数组是值类型，赋值和传参会复制整个数组，而不是指针。因此改变副本的值，不会改变本身的值。
//    7.支持 "=="、"!=" 操作符，因为内存总是被初始化过的。
//    8.指针数组 [n]*T，数组指针 *[n]T。

//@todo 值拷贝行为会造成性能问题，通常会建议使用 slice，或数组指针。

//######################################################################################################################


type array struct {

}

var Array *array

func init(){
	Array = NewArray()
}

func NewArray() *array {
	return &array{

	}
}

//执行入口
func (thisObj *array) Do(params map[string]interface{}){
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

//一维数组-测试
//命令行-输入:{"optTag":"Array","optParams":{"methodName":"OneTest"}}
func (thisObj *array) OneTest(){
	var arr [5]int	//未初始化元素值为 0。
	arr[1] = 1
	fmt.Println(arr,len(arr))	//输出:[0 1 0 0 0] 5

	arr = [5]int{1,2}	//未初始化元素值为 0。
	fmt.Println(arr,len(arr))	//输出:[1 2 0 0 0] 5

	arr = [...]int{1,2,3,4,5}	//通过初始化值确定数组长度。
	fmt.Println(arr,len(arr))	//输出:[1 2 3 4 5] 5

	arr = [5]int{1:100,4:200}	//使用引号初始化元素。//未初始化元素值为 0。
	fmt.Println(arr,len(arr))	//输出:[0 100 0 0 200] 5

	//带有结构体的数组
	arrStruct := [...]struct{
		name string
		age int
	}{
		{"test1",20},
		{"test2",30},
	}
	fmt.Println(arrStruct,len(arrStruct))	//输出:[{test1 20} {test2 30}] 2
}

//多维数组-测试
//命令行-输入:{"optTag":"Array","optParams":{"methodName":"MultiTest"}}
func (thisObj *array) MultiTest(){
	//定义 长度为2个的数组,值都是整型数组,每个值(即整型数组)拥有3个整型元素
	arr := [2][3]int{
		{1,2},
		{3,4,5},
	}
	fmt.Println(arr,len(arr))	//输出:[[1 2 0] [3 4 5]] 2

	//定义 通过初始化值确定数组长度,每一个值都是长度为3的整型数组
	arr1 := [...][3]int{	//第2维度不能使用"...",即第2维度要有确定值
		{1,2,3},
		{4,5,6},
		{7,8,9},
		{0,0,0},
		{1,1,1},
	}
	//最终输出的是整个数组长度为5,每一个值都是整型数组 且 拥有3个整型元素
	fmt.Println(arr1,len(arr1))	//输出:[[1 2 3] [4 5 6] [7 8 9] [0 0 0] [1 1 1]] 5

	//第2维度使用...的错误示例:
	//arr2 := [...][...]int{
	//	{1,2,3,4,5},
	//	{1,2,3,4,5},
	//}
	//fmt.Println(arr2,len(arr2))
	//第2维度使用...后,程序报错:
	//# Go-Note/GoBase
	//GoBase\array.go:98:15: use of [...] array outside of array literal

	//第3维度使用示例:
	//	第3维度讲解说明:
	//	1.[...][5][2]int 	第1维度定义:由初始化值确定数组长度 且 初始化值时,每一个初始化值是长度为5的数组
	//	2.[5][2]int			第2维度定义:每一个值是2位长度的整型数组,该维度的数组长度是5位
	//	3.[2]int 			第3维度定义:每一个值是整型,该维度的数组长度是2位
	//	简单来讲就跟剥洋葱一样,一层剥下一层,每层有多少个,值又是什么
	arr3 := [...][5][2]int {//第1维度-初始化声明值有2个,所以第1维度的数组长度为2位
		{//第2维度-声明长度有5个,第2维度的数组长度为5位,但是初始化时只赋值了4个,剩余的1个就是以0为值
			{//第3维度-声明长度有2个,第3维度的数组长度为2位
				1,2,
			},
			{
				1,2,
			},
			{
				1,2,
			},{
				1,2,
			},
		},
		{
			{1,2},{1,2},{1,2},{1,2},
		},
	}
	//这里修改的是第1维度的最后一位元素的值
	arr3[len(arr3)-1] = [5][2]int{
		{3,4},{3,4},{3,4},{3,4},{3,4},
	}
	fmt.Println(arr3,len(arr3))	//输出:[[[1 2] [1 2] [1 2] [1 2] [0 0]] [[3 4] [3 4] [3 4] [3 4] [3 4]]] 2
}

//小测试(小测试方法命名,就以Exam+数字命名即可)
//	找出数组中和为给定值的两个元素的下标，例如数组[1,3,5,8,7]，
//	找出两个元素之和等于8的下标分别是（0，4）和（1，2）
//命令行-输入:{"optTag":"Array","optParams":{"methodName":"Exam1"}}
func (thisObj *array) Exam1(){

	arr := [5]int{1,3,5,8,7}

	//i:当前元素下标
	i := 0
	//j:下一个元素下标
	j := 0
	//循环标记
	forTag 	:= 0
	//页码,循环标记达到一轮后,该页码加1;当页码达到数组长度就停止循环
	page 	:= 0
	//数组长度
	arrLen 	:= len(arr)

	//存储结果
	res := make([]string,0)
	for {
		//下一个元素下标
		j+=1

		//下一个元素下标,不超过数组长度,计算和值比较
		if j<arrLen {
			sum := arr[i]+ arr[j]
			if sum==8 {
				res = append(res,strconv.Itoa(i)+","+strconv.Itoa(j))
				//fmt.Println("~~~~~~~~~~~~~~~~")
				//fmt.Println("两个元素之和等于8的下标分别是:",i,j)
				//fmt.Println("~~~~~~~~~~~~~~~~")
			}
		}

		//每轮循环计数
		forTag++
		//若达到循环计数,进行相关计数和重置
		if forTag>=arrLen {
			//页码达到一轮就计数加1
			page++
			//每轮循环计数完毕,从下一轮开始计数
			forTag = page
			//第一个元素和第二个元素下标介入到下一轮
			i = page
			j = i
		}

		//若页码达到全部元素比较完毕,则退出循环
		if page>=arrLen {
			break
		}
	}

	//循环输出结果:
	//~~~~~~~~~~~~~~~~
	//两个元素之和等于8的下标分别是: 0 4
	//~~~~~~~~~~~~~~~~
	//~~~~~~~~~~~~~~~~
	//两个元素之和等于8的下标分别是: 1 2
	//~~~~~~~~~~~~~~~~

	//调试输出:
	fmt.Println("两个元素之和等于8的下标分别是:",res)
	//输出结果：
	//两个元素之和等于8的下标分别是: [0,4 1,2]
}