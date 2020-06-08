package goHighLevel

import (
	"fmt"
	"reflect"
)

//针对结构体强转为nil的处理
//
//总结:
//1.结构体强转为nil必须为指针结构体才可以!
//2.指针结构体强转nil后,直接与nil匹配,返回结果是true!
//3.指针结构体强转nil后,再传入给函数形参,然后函数形参接收类型是interface{}时,与nil比较相等的结果是false!
//	个人估计是interface{}值和类型都是nil,凡是跟interface{}比较,不但比较值还比较类型.
//
//命令行-输入:{"optTag":"GoHighLevel","optParams":{"methodName":"Nil_StructToNil"}}
func (thisObj *goHighLevelStruct) Nil_StructToNil(){

	//如下2种写法都是表示强转 将 xxx 强转为 (*goHighLevelStruct) 该值和该类型
	fmt.Println( (*goHighLevelStruct)(&goHighLevelStruct{}) ) //这个强转的值是&{}(即*goHighLevelStruct结构体),类型是 *goHighLevelStruct
	fmt.Println( (*goHighLevelStruct)(nil) )	//这个强转的值为nil,但是类型是 *goHighLevelStruct
	//指针结构体强转nil后直接与nil匹配,返回结果是true
	fmt.Println( (*goHighLevelStruct)(nil)==nil )
	//但是 指针结构体强转nil后 再传入给函数形参,然后函数形参接收类型是interface{}时,与nil比较相等的结果是false
	//	个人估计是interface{}值和类型都是nil,凡是跟interface{}比较,不但比较值还比较类型.
	thisObj.Nil_StructToNilFunc( (*goHighLevelStruct)(nil) )

	fmt.Println("----------------- 如下显示对应的值和数据类型 -----------------")

	fmt.Println("指针结构体示例:")
	fmt.Println( "(*goHighLevelStruct)(&goHighLevelStruct{}) => ",
		(*goHighLevelStruct)(&goHighLevelStruct{}),
		fmt.Sprint(reflect.TypeOf((*goHighLevelStruct)(&goHighLevelStruct{}))) )
	//输出结果:(*goHighLevelStruct)(&goHighLevelStruct{}) =>  &{} *goHighLevel.goHighLevelStruct

	fmt.Println( "(*goHighLevelStruct)(nil) => ",
		(*goHighLevelStruct)(nil),
		fmt.Sprint(reflect.TypeOf((*goHighLevelStruct)(nil))) )
	//输出结果:(*goHighLevelStruct)(nil) =>  <nil> *goHighLevel.goHighLevelStruct

	fmt.Println("值结构体示例:")
	fmt.Println( "(goHighLevelStruct)(goHighLevelStruct{}) => ",
		(goHighLevelStruct)(goHighLevelStruct{}),
		fmt.Sprint(reflect.TypeOf((goHighLevelStruct)(goHighLevelStruct{}))) )
	//输出结果:(goHighLevelStruct)(goHighLevelStruct{}) =>  {} goHighLevel.goHighLevelStruct

	//如下代码注释是因为打开后,编译都无法通过,报错提示如下:
	//	# Go-Note/goHighLevel
	//	goHighLevel\nil.go:39:22: cannot convert nil to type goHighLevelStruct
	//	goHighLevel\nil.go:40:48: cannot convert nil to type goHighLevelStruct
	//
	//fmt.Println( "(goHighLevelStruct)(nil) => ",
	//	(goHighLevelStruct)(nil),
	//	fmt.Sprint(reflect.TypeOf((goHighLevelStruct)(nil))) )

}
func (thisObj *goHighLevelStruct) Nil_StructToNilFunc(x interface{}){
	if x==nil {
		fmt.Println("the x is nil:",x,fmt.Sprint(reflect.TypeOf(x)))
	}else{
		fmt.Println("the x is not nil:",x,fmt.Sprint(reflect.TypeOf(x)))
	}
}
