package GoBase

import "fmt"

//######################################################################################################################
//GO-指针
//	GO-关于指针的说明:
//	  参考文章地址:
//		http://topgoer.com/go%E5%9F%BA%E7%A1%80/%E6%8C%87%E9%92%88.html
//
//	@todo 总结： 取地址操作符&和取值操作符*是一对互补操作符，
// 	@todo 		&取出地址，*根据地址取出地址指向的值。
//
//	@todo 变量、指针地址、指针变量、取地址、取值的相互关系和特性如下：
//  @todo   1.对变量进行取地址（&）操作，可以获得这个变量的指针变量。
//  @todo   2.指针变量的值是指针地址。
//  @todo   3.对指针变量进行取值（*）操作，可以获得指针变量指向的原变量的值。

//@todo 关于new和make,Go语言中new和make是内建的两个函数，主要用来分配内存!
//
//	new是一个内置的函数，它的函数签名如下：
//    func new(Type) *Type
//	其中，
//  @todo   1.Type表示类型，new函数只接受一个参数，这个参数是一个类型
//  @todo   2.*Type表示类型指针，new函数返回一个指向该类型内存地址的指针。
//	new函数不太常用，使用new函数得到的是一个类型的指针，并且该指针对应的值为该类型的零值。
//	示例:
//    var a *int		//只声明了指针变量,未分配内存地址
//    a = new(int)		//使用new对该指针变量分配内存地址
//    *a = 10			//再进行赋值,注意赋值是以*a而不是a,如果是a的话,就得要有地址赋给a
//    fmt.Println(*a)	//输出为10
//
//	@todo make也是用于内存分配的，区别于new，它只用于slice、map以及chan的内存创建，
// 	@todo 而且它返回的类型就是这三个类型本身，而不是他们的指针类型，因为这三种类型就是引用类型，所以就没有必要返回他们的指针了。
// 		make函数的函数签名如下：
//		func make(t Type, size ...IntegerType) Type
//	make函数是无可替代的，我们在使用slice、map以及channel的时候，都需要使用make进行初始化，然后才可以对它们进行操作。
//	示例:
//	  var b map[string]int			//声明一个map,但是没有分配内存地址
//    b = make(map[string]int)		//使用make进行内存地址分配
//    b["测试"] = 100				//进行相关赋值操作
//    fmt.Println(b)				//输出为map[测试:100]
//	最直接的map赋值方式:
//		b := make(map[string]int)
//	要么是直接声明并赋值
//		b := map[string]int{
// 			"1":2,
// 		}
//
//@todo new与make的区别
//    1.二者都是用来做内存分配的。
//    2.make只用于slice、map以及channel的初始化，返回的还是这三个引用类型本身；
//    3.而new用于类型的内存分配，并且内存对应的值为类型零值，返回的是指向类型的指针。
//######################################################################################################################


type pointer struct {

}

var Pointer *pointer

func init(){
	Pointer = NewPointer()
}

func NewPointer() *pointer {
	return &pointer{

	}
}

//执行入口
func (this *pointer) Do(params map[string]interface{}){
	//传参必须-标记
	doTag := params["doTag"].(string)
	switch doTag {
	//小测试
	case "exam":
		{
			this.Exam()
		}
	}
}

//小测试
//指针小练习
//程序定义一个int变量num的地址并打印
//将num的地址赋给指针ptr，并通过ptr去修改num的值
func (this *pointer) Exam(){
	var num int
	var ptr *int
	fmt.Println(&num) //输出如:0xc000090040
	ptr = &num	//从开发角度上讲,很少是这样需要引用地址处理,除非是数据量大的情况下
	*ptr = 100
	fmt.Println(num) //输出如:100
}
