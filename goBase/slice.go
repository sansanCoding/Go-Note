package goBase

import (
	"Go-Note/util"
	"fmt"
	"reflect"
)

//######################################################################################################################
//GO-切片Slice
//	GO-关于切片Slice的说明:
//	  参考文章地址:
//		http://topgoer.com/go%E5%9F%BA%E7%A1%80/%E5%88%87%E7%89%87Slice.html
//
//@todo 需要说明，slice 并不是数组或数组指针。它通过内部指针和相关属性引用数组片段，以实现变长方案。
// 	  1. 切片：切片是数组的一个引用，因此切片是引用类型。但自身是结构体，值拷贝传递。
//    2. 切片的长度可以改变，因此，切片是一个可变的数组。
//    3. 切片遍历方式和数组一样，可以用len()求长度。表示可用元素数量，读写操作不能超过该限制。
//    4. cap可以求出slice最大扩张容量，不能超出数组限制。0 <= len(slice) <= len(array)，其中array是slice引用的数组。
//    5. 切片的定义：var 变量名 []类型，比如 var str []string  var arr []int。
//    6. 如果 slice == nil，那么 len、cap 结果都等于 0。
//######################################################################################################################

//关于切片的:冒号说明:
//	a.冒号之前是起始位,获取值时不包含该起始位.
//	b.冒号之后是截止位,获取值时包含该截止位.
//	简单示例:
//		arr := []int{1,2,3,4,5,6,7,8,9,10,}
//		arr[0:1] 输出的是1
//		arr[6:8] 输出的是[7 8]

type slice struct {

}

var Slice *slice

func init(){
	Slice = NewSlice()
}

func NewSlice() *slice {
	return &slice{

	}
}

//执行入口
func (thisObj *slice) Do(params map[string]interface{}){
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

//切片复制
//copy ：函数 copy 在两个 slice 间复制数据，复制长度以 len 小的为准。两个 slice 可指向同一底层数组，允许元素区间重叠。
//总结:
//	1.copy()函数是将第二个切片 复制到 第一个切片 上，复制长度以 len 小的为准。-- 参考示例1
//	2.若是同一个切片，切分到多个变量上，则多个变量之间互相copy()也会影响原切片上。-- 参考示例2
//@todo 应及时将所需数据 copy 到较小的 slice，以便释放超大号底层数组内存。
//调试-命令行输入:
//	{"optTag":"Slice","optParams":{"methodName":"Copy"}}
func (thisObj *slice) Copy(){
	//示例1:
	fmt.Println("------------------示例1:------------------")
	arr := []int{4,5,6,}
	arr1 := []int{1,2,}

	//将arr的值复制给arr1,由于arr1的长度只有2位,所以arr1的1和2,被4和5覆盖了
	copy(arr1,arr)
	//调试输出
	fmt.Println("arr:",arr)
	fmt.Println("arr1:",arr1)
	//输出:
	//arr: [4 5 6]
	//arr1: [4 5]

	//示例2:
	fmt.Println("------------------示例2:------------------")
	{
		data := []int{00,11,22,33,44,55,66,77,88,99}
		arr1 := data[:3]
		arr2 := data[3:]

		fmt.Println("copy前:")
		fmt.Println("arr1:",arr1)
		fmt.Println("arr2:",arr2)

		copy(arr1,arr2)

		fmt.Println("copy后:")
		fmt.Println("arr1:",arr1)
		fmt.Println("arr2:",arr2)
		fmt.Println("data:",data)	//两个同源的切片，互相copy()后，也会影响到原切片的值

		//输出结果:
		//------------------示例2:------------------
		//copy前:
		//arr1: [0 11 22]
		//arr2: [33 44 55 66 77 88 99]
		//copy后:
		//arr1: [33 44 55]
		//arr2: [33 44 55 66 77 88 99]
		//data: [33 44 55 33 44 55 66 77 88 99]
	}

}

//数组和切片的引用关系
//命令行-输入:{"optTag":"Slice","optParams":{"methodName":"ArrayAndSliceRelation"}}
func (thisObj *slice) ArrayAndSliceRelation(){
	fmt.Println("~~~~~~~~~~~~~~~~~~~~ 示例1 start ~~~~~~~~~~~~~~~~~~~~")
	{
		var s []int
		a := [9]int{1,2,3,4,5,6,7,8,9}
		s = a[2:4]

		a[2] = 33
		a[3] = 44
		fmt.Println("s:",s)
		fmt.Println("a:",a)

		s[0] = 133
		s[1] = 244
		fmt.Println("s:",s)
		fmt.Println("a:",a)

		//输出结果:
		//s: [33 44]
		//a: [1 2 33 44 5 6 7 8 9]
		//s: [133 244]
		//a: [1 2 133 244 5 6 7 8 9]
	}
	fmt.Println("~~~~~~~~~~~~~~~~~~~~ 示例1 end ~~~~~~~~~~~~~~~~~~~~")

	fmt.Println("~~~~~~~~~~~~~~~~~~~~ 示例2 start ~~~~~~~~~~~~~~~~~~~~")
	{
		a := [...]int{1,2,3}

		testFunc := func(a []int){
			a[1] = 99
		}

		fmt.Println("a:",a)

		fmt.Println(reflect.TypeOf(a[0:2]))

		//数组先进行元素截取,这样会形成切片,再进行函数传参,修改后的值也会影响原数组的值
		testFunc(a[0:2])

		fmt.Println("a:",a)

		//输出结果:
		//a: [1 2 3]
		//a: [1 99 3]
	}
	fmt.Println("~~~~~~~~~~~~~~~~~~~~ 示例2 end ~~~~~~~~~~~~~~~~~~~~")

	fmt.Println("~~~~~~~~~~~~~~~~~~~~ 示例3 start ~~~~~~~~~~~~~~~~~~~~")
	{
		a := [...]int{1,2,3}

		//数组是值类型,赋值和传参是复制整个数组
		testFunc := func(a [3]int){
			//这里的修改针对的是变量副本,不是变量本身
			a[1] = 99
		}

		fmt.Println("a:",a)
		testFunc(a)
		fmt.Println("a:",a)

		//输出结果:
		//a: [1 2 3]
		//a: [1 2 3]
	}
	fmt.Println("~~~~~~~~~~~~~~~~~~~~ 示例3 end ~~~~~~~~~~~~~~~~~~~~")

}
