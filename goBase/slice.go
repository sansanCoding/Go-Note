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
//	a.冒号之前是起始位,获取值时包含该起始位.
//	b.冒号之后是截止位,获取值时不包含该截止位.
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

//append追加元素
//append ：向 slice 尾部添加数据，返回新的 slice 对象。
//总结:
//	如果想让原切片变量赋予给新切片变量后，新切片变量的任何修改操作都不影响到原切片的话，
//	目前只好是用原切片进行循环遍历，挨个将值赋予给新切片变量存储即可!
//调试-命令行输入:
//	{"optTag":"Slice","optParams":{"methodName":"Append"}}
func (thisObj *slice) Append(){
	arr := make([]int,0)
	arr = append(arr,1,2,3,4,5)
	//arr1 := append(arr,11,22,33,44,55)	//如果这里append是多个值,修改arr1的值不会影响到arr的值
	arr1 := append(arr,[]int{111}...)		//如果这里append是1个值,修改arr1的值会影响到arr的值
	arr1[0] = 888	//注意这里，arr1是由arr进行append后分配的内存地址，但是append只有1个值后，修改arr1的值也会影响到arr的值

	//输出调试:
	fmt.Println("arr:",arr,fmt.Sprintf("%p",&arr) )
	fmt.Println("arr1:",arr1,fmt.Sprintf("%p",&arr1) )
	//输出结果:
	//arr: [888 2 3 4 5] 0xc000004740
	//arr1: [888 2 3 4 5 111] 0xc000004760

	//@todo 这里感觉有个bug,append追加2个元素以上后赋予给新变量，新变量的任何操作才不会影响原切片....
	fmt.Println("直接获取原切片所有的值赋予给新变量，新变量的任何操作不会影响原切片-示例:")
	arr11 := append(arr,10000,20000)
	arr11[2] = 1111
	//输出调试:
	fmt.Println("arr:",arr,fmt.Sprintf("%p",&arr) )
	fmt.Println("arr11:",arr11,fmt.Sprintf("%p",&arr11) )
	//输出结果:
	//直接获取原切片所有的值赋予给新变量，新变量的任何操作不会影响原切片-示例:
	//arr: [888 2 3 4 5] 0xc000004740
	//arr11: [888 2 1111 4 5 10000 20000] 0xc0000047c0
	//-------------------------------------------------------------------
	fmt.Println("-------------------------------------------------------------------")

	arr2 := arr[:]
	arr3 := arr
	arr4 := arr[:2]
	arr5 := arr[2:]
	arr5[0] = 999	//注意这里，虽然arr5是新分配了内存地址，但是由于是arr切分出来的，切片底层依然会影响到arr的对应下标的值

	//输出示例:
	fmt.Println("arr:",arr,fmt.Sprintf("%p",&arr) )
	fmt.Println("arr1:",arr1,fmt.Sprintf("%p",&arr1) )
	fmt.Println("arr2:",arr2,fmt.Sprintf("%p",&arr2) )
	fmt.Println("arr3:",arr3,fmt.Sprintf("%p",&arr3) )
	fmt.Println("arr4:",arr4,fmt.Sprintf("%p",&arr4) )
	fmt.Println("arr5:",arr5,fmt.Sprintf("%p",&arr5) )
	//输出结果:
	//arr: [888 2 999 4 5] 0xc000004740
	//arr1: [888 2 999 4 5 111] 0xc000004760
	//arr2: [888 2 999 4 5] 0xc000004820
	//arr3: [888 2 999 4 5] 0xc000004840
	//arr4: [888 2] 0xc000004860
	//arr5: [999 4 5] 0xc000004880

	fmt.Println("TestEnd!")
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
