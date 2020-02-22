package goBase

import (
	"Go-Note/util"
	"encoding/json"
	"fmt"
)

//######################################################################################################################
//GO-map
//	参考文章地址:
//		http://topgoer.com/go%E5%9F%BA%E7%A1%80/%E7%BB%93%E6%9E%84%E4%BD%93.html
//----------------------------------------------------------------------------------------------------------------------
//1-0.使用type和struct关键字来定义结构体，具体代码格式如下：
//		type 类型名 struct {
//			字段名 字段类型
//			字段名 字段类型
//			…
//		}
//
//		a.类型名：标识自定义结构体的名称，在同一个包内不能重复。
//		b.字段名：表示结构体字段名。结构体中的字段名必须唯一。
//		c.字段类型：表示结构体字段的具体类型。
//
//		示例:
//		//结构体声明:
//		type person struct {
//			name 	string
//			address string
//			age 	int
//		}
//		//或者 同样的类型字段写在一行
//		type person struct {
//			name,address string
//			age int
//		}
//
//1-1-1.结构体初始化
//		示例:
//		type test struct {
//			name string
//			age int
//		}
//		var Test test	//值-结构体初始化
//		var Test *test	//指针-结构体初始化
//
//1-1-2.匿名结构体
//		示例:
//		var User struct{name string;age int}
//		User.name = "test"
//		User.age = 21
//
//1-2.使用键值对初始化
//		示例:
//		type test struct {
//			name string
//			age int
//		}
//
//		//值-键值对初始化
//		Test := test{
//			name:"test2"
// 			age:22,,
// 		}
//		//指针-键值对初始化
//		Test := &test{
//			name:"test3",
//			//age:"23",	//初始化时,某些字段值没有赋值时,则值为该数据类型的零值
// 		}
//
//1-3.使用值的列表初始化
//		示例:
//		type test struct {
//			name string
//			age int
//		}
//		//值-结构体初始化
//		Test := test{
// 			"test",
//			21,
// 		}
//		//指针-结构体初始化
//		Test := &test{
// 			"test",
//			21,
// 		}
//		使用这种格式初始化时，需要注意：
//    		a.必须初始化结构体的所有字段。
//    		b.初始值的填充顺序必须与字段在结构体中的声明顺序一致。
//    		c.该方式不能和键值初始化方式混用。
//
//----------------------------------------------------------------------------------------------------------------------
//2.只有当结构体实例化时，才会真正地分配内存。也就是必须实例化后才能使用结构体的字段。
//	结构体本身也是一种类型，我们可以像声明内置类型一样使用var关键字声明结构体类型。
//		语法:
//		var 结构体实例 结构体类型
//
//		基本实例化示例:
//		type test struct {
//			name string
//			age int
//		}
//		var Test test
//		Test.name 	= "test"
//		Test.age 	= 21
//----------------------------------------------------------------------------------------------------------------------
//3.创建指针类型结构体(即使用new实例化)
//		示例:
//		type test struct {
//			name string
//			age int
//		}
//		testObj := new(test)	//创建的是结构体指针
//		testObj.name = "test"	//Go语言中支持对 结构体指针 直接使用.来访问结构体的成员。
//		testObj.age = 21
//
//		//调试输出
//		fmt.Printf("%v",testObj)
//		fmt.Println("")
//		fmt.Printf("%#v",testObj)
//		fmt.Println("")
//		fmt.Printf("%+v",testObj)
//		fmt.Println("")
//
//		//输出结果:
//		//&{test 21}
//		//&goBase.test{name:"test", age:21}
//		//&{name:test age:21}
//----------------------------------------------------------------------------------------------------------------------
//4.取结构体的地址实例化(即使用&实例化)
//	使用&对结构体进行取地址操作相当于对该结构体类型进行了一次new实例化操作。
//	这是Go语言的语法糖。
//		示例:
//		//未赋值的结构体字段,则值为数据类型的零值
//		testObj := &test{}
//
//		//调试输出:
//		fmt.Printf("%v",testObj)
//		fmt.Println("")
//		fmt.Printf("%#v",testObj)
//		fmt.Println("")
//		fmt.Printf("%+v",testObj)
//		fmt.Println("")
//		//输出结果:
//		//&{ 0}
//		//&goBase.test{name:"", age:0}
//		//&{name: age:0}
//----------------------------------------------------------------------------------------------------------------------
//5.构造函数
//	因为struct是值类型，如果结构体比较复杂的话，值拷贝性能开销会比较大，所以该构造函数返回的是结构体指针类型。
//	示例:
//		type test struct {
//			name string
//			age int
//		}
//		func NewTest(params map[string]interface{}) *test {
//			nameStr := params["name"].(string)
//			ageInt 	:= params["age"].(int)
//			return &test{
//				name:nameStr,
//				age:ageInt,
//			}
//		}
//----------------------------------------------------------------------------------------------------------------------
//6.方法和接收者(即声明一个属于该对象里的方法)
//	(1).方法定义格式如下
//		//接收者变量就类似于其他语言中的this 或者 self。
//		func (接收者变量 接收者类型) 方法名(参数列表) (返回参数) {
//			函数体
//		}
//
//		a.接收者变量：接收者中的参数变量名在命名时，官方建议使用接收者类型名的第一个小写字母，而不是self、this之类的命名。例如，Person类型的接收者变量应该命名为 p，Connector类型的接收者变量应该命名为c等。
//	 	b.接收者类型：接收者类型和参数类似，可以是指针类型和非指针类型。
//	 	c.方法名、参数列表、返回参数：具体格式与函数定义相同。
//
//	(2).关于 指针类型的接收者 和 值类型的接收者 不同之处
//		a.指针类型的接收者,若修改接收者里任一成员变量,在其他地方再调用该被修改的成员变量时,都是修改后的效果!
//		b.值类型的接收者,Go语言会在代码运行时将接收者的值复制一份。在值类型接收者的方法中可以获取接收者的成员值，但修改操作只是针对副本，无法修改接收者变量本身。
//		c.注意:
//			1.只要接收者是值类型,在所有【值-接收者方法】里面对任何字段进行修改,在外部调用后无效!
//			2.【值-接收者】结构体外部修改任何字段,会对所有【值-接收者方法】里面获取任何字段的值都是修改后的!
//			3.总结就是
// 					内部-值类型接收者,在方法里面针对字段修改是以值接收者的副本进行修改,作用域在该方法里,外部调用或内部其他方法调用都是无效!
//					外部-值类型结构体,在外部针对字段修改,值类型接收者方法里面获取的是修改后的值!
//
//		指针类型接收者示例:
//		func (t *test) Index(ageInt int) {
//			t.age = ageInt
//		}
//		值类型接收者示例:
//		func (t test) Index(ageInt int) {
//			t.age = ageInt
//		}
//
//	(3).什么时候应该使用指针类型接收者
//		a.需要修改接收者中的值。
//    	b.接收者是拷贝代价比较大的大对象。
//    	c.保证一致性，如果有某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者。
//		--------------------------
//		总结一下:
//		如果跟其他语言的对象一样,就使用指针接收者.
//----------------------------------------------------------------------------------------------------------------------
//7.任意类型添加方法
//	a.接收者的类型可以是任何类型，不仅仅是结构体，任何类型(包含自定义类型)都可以拥有方法。
//	b.注意事项： 非本地类型不能定义方法，也就是说我们不能给别的包的类型定义方法。
//
//	示例:
//		type MyInt int
//
//		//针对任意类型添加的方法:
// 		//	1.接收者是针对该任意类型本身,目前测试的效果是修改该任意类型值,再调用该任意类型添加的方法,获取该任意类型的值是修改后的。
//		//	2.只要是该任意类型有添加方法,不影响该任意类型值的使用,如type MyInt int,定义的整型值,还是继续按定义的整型值使用。
//		func (m MyInt) Run(){
//			fmt.Println("this is MyInt Run Func,m is ",m)
//		}
//
//		func (thisObj *structType) AnyTypeBindFunc(){
//			var test MyInt
//			test.Run()
//			test = 100
//			test.Run()
//			//输出结果:
//			//this is MyInt Run Func,m is  0
//			//this is MyInt Run Func,m is  100
//		}
//----------------------------------------------------------------------------------------------------------------------
//8.结构体的匿名字段
//	a.结构体允许其成员字段在声明时没有字段名而只有类型，这种没有名字的字段就称为匿名字段。
//	b.匿名字段默认采用类型名作为字段名，结构体要求字段名称必须唯一，因此一个结构体中同种类型的匿名字段只能有一个。
//	c.该方式无法被外部包直接所使用,直接使用会报错如下:
//		.\main.go:133:5: implicit assignment of unexported field 'string' in goBase.AnonymousField literal
//		.\main.go:134:5: implicit assignment of unexported field 'int' in goBase.AnonymousField literal
//
//	示例:
//	type test struct {
//		string
//		int
//	}
//
//	func (t test) Index(){
//		fmt.Println(t.string)
//		fmt.Println(t.int)
//	}
//
//	var Test test
//	Test.Index()
//	Test.string = "Test"
//	Test.int = 21
//	Test.Index()
//----------------------------------------------------------------------------------------------------------------------
//9.嵌套结构体
//	a.一个结构体中可以嵌套包含另一个结构体或结构体指针。
//
//	示例:
//	type UserExtension struct {
//		Address string
//	}
//	type User struct {
//		Name 			string
//		Age				int
//		UserExtension 	UserExtension		//指明 某个字段 的值为一个 值-结构体(也可以声明成 指针-结构体)
//	}
//
//	参考方法:
// 	func (thisObj *structType) NestStruct(){ ... }
//----------------------------------------------------------------------------------------------------------------------
//10.嵌套匿名结构体
//	a.当访问结构体成员时会先在结构体中查找该字段，找不到再去匿名结构体中查找。
//
//	示例:
//	type TestExt struct {
//		Address string
//	}
//	type Test struct {
//		Name string
//		Age int
//		TestExt				//匿名结构体,按 匿名结构体.字段名 或 直接匿名结构体的字段名 获取字段值
//	}
//
//	参考方法:
// 	func (thisObj *structType) NestAnonymousStruct(){ ... }
//----------------------------------------------------------------------------------------------------------------------
//11.嵌套结构体的字段名冲突
//	a.嵌套结构体内部可能存在相同的字段名。这个时候为了避免歧义需要指定具体的内嵌结构体的字段。
//
//	示例:
//	type Test1 struct {
//		Address string
//	}
//	type Test2 struct {
//		Address string
//	}
//	type Test struct {
//		Name string
//		Test1
//		Test2
//	参考方法:
//	func (thisObj *structType) NestAnonymousStructUnq(){ ... }
//----------------------------------------------------------------------------------------------------------------------
//12.结构体的"继承"
//	参考方法:
//	func (thisObj *structType) Inherit() { ... }
//----------------------------------------------------------------------------------------------------------------------
//13.结构体字段的可见性
//	a.结构体中字段大写开头表示可公开访问，小写表示私有（仅在定义当前结构体的包中可访问）。
//
//	示例:
//	type test struct {
//		Name string		//外部包可调用该Name字段
//		age int			//当前包可调用该age字段
//	}
//----------------------------------------------------------------------------------------------------------------------
//14.结构体与JSON序列化
// 	//JSON序列化：结构体-->JSON格式的字符串
//	//	json序列化是默认使用字段名作为key
//  //data, err := json.Marshal(结构体)
//----------------------------------------------------------------------------------------------------------------------
//15.结构体标签（Tag）
//	a.结构体标签由一个或多个键值对组成。键与值使用冒号分隔，值用双引号括起来。键值对之间使用一个空格分隔。
// 		注意事项： 	为结构体编写Tag时，必须严格遵守键值对的规则。
// 					结构体标签的解析代码的容错能力很差，一旦格式写错，编译和运行时都不会提示任何错误，通过反射也无法正确取值。
// 					例如不要在key和value之间添加空格。
//
// 	Tag在结构体字段的后方定义，由一对反引号包裹起来，具体的格式如下：
//    `key1:"value1" key2:"value2"`
//
//	示例:
//	type test struct {
//		Name string `json:"name"`	//通过指定tag实现json序列化该字段时的key,多个重复的以第一个为准
//		Address string 				//json序列化是默认使用字段名作为key
//		age int 					//私有不能被json包访问
//
//	}
//	testObj := test{
//		Name:"test",
//		age:21,
//	}
//	testJsonByte,testJsonErr := json.Marshal(testObj)
//
//	//输出调试:
//	fmt.Println("testJsonStr:",string(testJsonByte))
//	fmt.Println("testJsonErr:",testJsonErr)
//	//输出结果:
//	//testJsonStr: {"name":"test"}
//	//testJsonErr: <nil>
//######################################################################################################################

type structType struct {

}

var StructType *structType

func init(){
	StructType = NewStructType()
}

func NewStructType() *structType {
	return &structType{

	}
}

//执行入口
func (thisObj *structType) Do(params map[string]interface{}){
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
//命令行-输入:{"optTag":"StructType","optParams":{"methodName":"Index"}}
func (thisObj *structType) Index(){
	//结构体声明
	type test struct {
		name string
		age int
	}

	//值-结构体初始化
	{
		var Test test
		Test.name 	= "test"
		Test.age 	= 21

		//按 值的默认格式 输出
		fmt.Println(Test)

		//值的默认格式表示
		fmt.Printf("%v",Test)
		fmt.Println("")

		//值的Go语法表示
		fmt.Printf("%#v",Test)
		fmt.Println("")

		//类似%v，但输出结构体时会添加字段名
		fmt.Printf("%+v",Test)
		fmt.Println("")

		//输出结果:
		//{test 21}
		//{test 21}
		//goBase.test{name:"test", age:21}
		//{name:test age:21}
	}

	//指针-结构体初始化
	{
		var Test *test
		Test = new(test)
		Test.name = "test2"
		Test.age = 22

		//调试输出:
		fmt.Println(Test)
		fmt.Printf("%v",Test)
		fmt.Println("")
		fmt.Printf("%#v",Test)
		fmt.Println("")
		fmt.Printf("%+v",Test)
		fmt.Println("")
		//输出结果:
		//&{test2 22}
		//&{test2 22}
		//&goBase.test{name:"test2", age:22}
		//&{name:test2 age:22}
	}

	//使用值的列表初始化
	{
		//值-结构体初始化
		//Test := test{
		//	"test",
		//	21,
		//}
		//指针-结构体初始化
		Test := &test{
			"test",
			21,
		}
		fmt.Printf("%v---%#v---%+v",Test,Test,Test)
		fmt.Println("")
		//值-结构体初始化-输出结果:
		//{test 21}---goBase.test{name:"test", age:21}---{name:test age:21}
		//指针-结构体初始化-输出结果:
		//&{test 21}---&goBase.test{name:"test", age:21}---&{name:test age:21}
	}
}

//匿名结构体
//命令行-输入:{"optTag":"StructType","optParams":{"methodName":"Anonymous"}}
func (thisObj *structType) Anonymous(){
	var User struct{name string;age int}
	User.name = "test"
	User.age = 21

	//调试输出:
	fmt.Printf("%v",User)
	fmt.Println("")
	fmt.Printf("%#v",User)
	fmt.Println("")
	fmt.Printf("%+v",User)
	fmt.Println("")

	//输出结果:
	//{test 21}
	//struct { name string; age int }{name:"test", age:21}
	//{name:test age:21}
}

//指针类型结构体
//命令行-输入:{"optTag":"StructType","optParams":{"methodName":"Pointer"}}
func (thisObj *structType) Pointer(){
	type test struct {
		name string
		age int
	}
	testObj := new(test)
	testObj.name = "test"
	testObj.age = 21

	//调试输出:
	fmt.Printf("%v",testObj)
	fmt.Println("")
	fmt.Printf("%#v",testObj)
	fmt.Println("")
	fmt.Printf("%+v",testObj)
	fmt.Println("")

	//输出结果:
	//&{test 21}
	//&goBase.test{name:"test", age:21}
	//&{name:test age:21}
}

//取结构体的地址实例化
//命令行-输入:{"optTag":"StructType","optParams":{"methodName":"Pointer2"}}
func (thisObj *structType) Pointer2(){
	type test struct {
		name string
		age int
	}

	//注意:这里的语法 花括号 包起来的作用域!!!

	//实例化1:
	{
		//未赋值的结构体字段,则值为数据类型的零值
		testObj := &test{}

		//调试输出:
		fmt.Printf("%v",testObj)
		fmt.Println("")
		fmt.Printf("%#v",testObj)
		fmt.Println("")
		fmt.Printf("%+v",testObj)
		fmt.Println("")
		//输出结果:
		//&{ 0}
		//&goBase.test{name:"", age:0}
		//&{name: age:0}
	}

	//实例化2:
	{
		testObj := &test{
			name:"test",
			age:21,
		}

		//调试输出:
		fmt.Printf("%v",testObj)
		fmt.Println("")
		fmt.Printf("%#v",testObj)
		fmt.Println("")
		fmt.Printf("%+v",testObj)
		fmt.Println("")
		//输出结果:
		//&{test 21}
		//&goBase.test{name:"test", age:21}
		//&{name:test age:21}
	}
}

//面试题
//命令行-输入:{"optTag":"StructType","optParams":{"methodName":"CaseInterview"}}
func (thisObj *structType) CaseInterview(){
	type student struct {
		name string
		age  int
	}

	m := make(map[string]*student)
	stus := []student{
	   {name: "pprof.cn", age: 18},
	   {name: "测试", age: 23},
	   {name: "博客", age: 28},
	}

	for _, stu := range stus {
		//引址造成的错误原因:
		//	这里由于是取址操作,循环处理会将最后一个值赋给stu,所以整个m的值都是stus最后一个值
		//修正:
		//stus改成如下:
		//	stus := []*student{
		//	   {name: "pprof.cn", age: 18},
		//	   {name: "测试", age: 23},
		//	   {name: "博客", age: 28},
		//	}
		//m[stu.name]的存储改为 m[stu.name] = stu
	   m[stu.name] = &stu
	}

	for k, v := range m {
	   fmt.Println(k, "=>", v.name)
	}
}

//任意类型绑定方法
//命令行-输入:{"optTag":"StructType","optParams":{"methodName":"AnyTypeBindFunc"}}
type myInt int

func (m myInt) Run(){
	fmt.Println("this is MyInt Run Func,m is ",m)
}

func (thisObj *structType) AnyTypeBindFunc(){
	var test myInt
	test.Run()
	test = 100
	test.Run()
	//输出结果:
	//this is MyInt Run Func,m is  0
	//this is MyInt Run Func,m is  100
}

//匿名字段结构体
//命令行-输入:{"optTag":"StructType","optParams":{"methodName":"AnonymousField"}}
func (thisObj *structType) AnonymousField(){
	type test struct {
		string
		int
	}

	//值类型-结构体赋值
	{
		testObj := test{
			"Test",
			21,
		}

		//调试输出:
		fmt.Println(testObj)
		fmt.Println(testObj.string)
		fmt.Println(testObj.int)
		//输出结果:
		//{Test 21}
		//Test
		//21
	}

	//指针类型-结构体赋值
	{
		testObj := &test{
			"Test",
			21,
		}

		//调试输出:
		fmt.Println(testObj)
		fmt.Println(testObj.string)
		fmt.Println(testObj.int)
		//输出结果:
		//&{Test 21}
		//Test
		//21
	}
}

//匿名字段结构体-绑定方法
type anonymousFieldBindFunc struct {
	string
	int
}

func (a anonymousFieldBindFunc) Index(){
	fmt.Println(a)
	fmt.Println(a.string)
	fmt.Println(a.int)
	//注意:
	// 	1.只要接收者是值类型,在所有【值-接收者方法】里面对任何字段进行修改,在外部调用后无效!
	//	2.【值-接收者】结构体外部修改任何字段,会对所有【值-接收者方法】里面获取任何字段的值都是修改后的!
	a.string = "test2222"
	a.int = 22
}

func (a anonymousFieldBindFunc) Second(){
	fmt.Println("Second~~~~:")
	fmt.Println(a)
	fmt.Println(a.string)
	fmt.Println(a.int)
	fmt.Println("Second~~~~!")
}

//命令行-输入:{"optTag":"StructType","optParams":{"methodName":"AnonymousFieldBindFunc"}}
func (thisObj *structType) AnonymousFieldBindFunc(){
	//值类型-结构体
	{
		var test anonymousFieldBindFunc
		test.Index()
		test.string = "test"
		test.int = 21
		test.Index()
		fmt.Println("test.string:",test.string)
		fmt.Println("test.int:",test.int)
		test.Second()
	}

	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~")

	//指针类型-结构体
	{
		var test *anonymousFieldBindFunc
		test = new(anonymousFieldBindFunc)
		test.Index()
		test.string = "test"
		test.int = 21
		test.Index()
		fmt.Println("test.string:",test.string)
		fmt.Println("test.int:",test.int)
		test.Second()
	}

	//输出结果:
	//{ 0}
	//
	//0
	//{test 21}
	//test
	//21
	//test.string: test
	//test.int: 21
	//Second~~~~:
	//{test 21}
	//test
	//21
	//Second~~~~!
	//~~~~~~~~~~~~~~~~~~~~~~
	//{ 0}
	//
	//0
	//{test 21}
	//test
	//21
	//test.string: test
	//test.int: 21
	//Second~~~~:
	//{test 21}
	//test
	//21
	//Second~~~~!
}

//嵌套结构体
//命令行-输入:{"optTag":"StructType","optParams":{"methodName":"NestStruct"}}
func (thisObj *structType) NestStruct(){
	type UserExtension struct {
		Address string
	}

	//嵌套 值-结构体
	{
		type User struct {
			Name 			string
			Age				int
			UserExtension 	UserExtension	//指明 某个字段值 是一个 值-结构体
		}

		var Test User
		Test.UserExtension = UserExtension{
			Address:"www.baidu.com",
		}

		//输出调试:
		fmt.Printf("%v",Test)
		fmt.Println("")
		fmt.Printf("%#v",Test)
		fmt.Println("")
		fmt.Printf("%+v",Test)
		fmt.Println("")

		//输出结果:
		//{ 0 {www.baidu.com}}
		//goBase.User{Name:"", Age:0, UserExtension:goBase.UserExtension{Address:"www.baidu.com"}}
		//{Name: Age:0 UserExtension:{Address:www.baidu.com}}
	}

	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")

	//嵌套 指针-结构体
	{
		type User struct {
			Name 			string
			Age				int
			UserExtension 	*UserExtension	//指明 某个字段值 是一个 值-结构体
		}

		var Test User
		Test.UserExtension = &UserExtension{
			Address:"www.baidu.com",
		}

		//输出调试:
		fmt.Printf("%#v",Test)
		fmt.Println("")
		fmt.Println(Test.UserExtension.Address)
		fmt.Println("")
		//输出结果:
		//goBase.User{Name:"", Age:0, UserExtension:(*goBase.UserExtension)(0xc0000e00f0)}
		//www.baidu.com
	}
}

//嵌套匿名结构体
//命令行-输入:{"optTag":"StructType","optParams":{"methodName":"NestAnonymousStruct"}}
func (thisObj *structType) NestAnonymousStruct(){
	type TestExt struct {
		Address string
	}
	type Test struct {
		Name string
		Age int
		TestExt
	}
	var test Test
	test.Address = "www.baidu.com"

	//输出调试:
	fmt.Println("test.TestExt.Address:",test.TestExt.Address)
	fmt.Println("test.Address:",test.Address)
	fmt.Println("----------------------")
	fmt.Printf("%v",test)
	fmt.Println("")
	fmt.Printf("%+v",test)
	fmt.Println("")
	fmt.Printf("%#v",test)
	fmt.Println("")
	//输出结果:
	//test.TestExt.Address: www.baidu.com
	//test.Address: www.baidu.com
	//----------------------
	//{ 0 {www.baidu.com}}
	//{Name: Age:0 TestExt:{Address:www.baidu.com}}
	//goBase.Test{Name:"", Age:0, TestExt:goBase.TestExt{Address:"www.baidu.com"}}
}

//嵌套匿名结构体-重复字段名冲突
//命令行-输入:{"optTag":"StructType","optParams":{"methodName":"NestAnonymousStructUnq"}}
func (thisObj *structType) NestAnonymousStructUnq(){
	type Test1 struct {
		Address string
	}
	type Test2 struct {
		Address string
	}
	type Test struct {
		Name string
		Test1
		Test2
	}

	var test Test
	//test.Address = "www.baidu.com" 		//报错:Ambiguous reference 'Address'
	test.Test1.Address = "www.baidu.com1"	//必须指明哪个结构体的字段名
	test.Test2.Address = "www.baidu.com2"	//必须指明哪个结构体的字段名

	//调试输出:
	fmt.Println("test.Test-#:",fmt.Sprintf("%#v",test))
	fmt.Println("test.Test-+:",fmt.Sprintf("%+v",test))
	//输出结果:
	//test.Test-#: goBase.Test{Name:"", Test1:goBase.Test1{Address:"www.baidu.com1"}, Test2:goBase.Test2{Address:"www.baidu.com2"}}
	//test.Test-+: {Name: Test1:{Address:www.baidu.com1} Test2:{Address:www.baidu.com2}}
}

//结构体继承-1
type testInherit struct {
	Name string
}
func (t *testInherit) Test(){
	fmt.Println("this is testInherit.Test!")
	fmt.Println("t.Name:",t.Name)
}
type testTest struct {
	*testInherit	//这种方式继承还需要额外创建一次指针-结构体
}
func (t *testTest) TestTest(){
	t.Name = "testTest111222333"
}
func (t *testTest) TestTest2(){
	fmt.Println("t.Name2:",t.Name)
}
//命令行-输入:{"optTag":"StructType","optParams":{"methodName":"Inherit"}}
func (thisObj *structType) Inherit() {
	testObj := new(testTest)
	testObj.testInherit = &testInherit{
		Name: "testqqqqqq",
	}
	testObj.Test()
	testObj.TestTest2()

	testObj.TestTest()

	testObj.Test()
	testObj.TestTest2()

	//输出结果:
	//this is testInherit.Test!
	//t.Name: testqqqqqq
	//t.Name2: testqqqqqq
	//this is testInherit.Test!
	//t.Name: testTest111222333
	//t.Name2: testTest111222333
}

//结构体继承-2
type testInherit2 struct {
	Name string
}
func (t *testInherit2) Test(){
	fmt.Println("this is testInherit2.Test!")
	fmt.Println("t.Name:",t.Name)
}
type testTest2 struct {
	testInherit2		//这种方式的继承,只需要new(testTest2)就包含了testInherit2 一并 指针-实例化,这个方式的继承更适合其他语言对象的继承逻辑!
}
func (t *testTest2) TestTest(){
	t.Name = "test456"
}
func (t *testTest2) TestTest2(){
	fmt.Println("t.Name2:",t.Name)
}
//命令行-输入:{"optTag":"StructType","optParams":{"methodName":"Inherit2"}}
func (thisObj *structType) Inherit2(){
	testTestObj := new(testTest2)
	testTestObj.TestTest()
	testTestObj.Test()
	testTestObj.Name = "test123"
	testTestObj.TestTest2()
	//输出结果:
	//this is testInherit2.Test!
	//t.Name: test456
	//t.Name2: test123
}

//json序列化
//命令行-输入:{"optTag":"StructType","optParams":{"methodName":"JsonEncode"}}
func (thisObj *structType) JsonEncode(){
	type test struct {
		Name string `json:"name"`
		age int
	}
	testObj := test{
		Name:"test",
		age:21,
	}
	testJsonByte,testJsonErr := json.Marshal(testObj)

	//输出调试:
	fmt.Println("testJsonStr:",string(testJsonByte))
	fmt.Println("testJsonErr:",testJsonErr)
	//输出结果:
	//testJsonStr: {"name":"test"}
	//testJsonErr: <nil>
}

//小练习
//命令行-输入:{"optTag":"StructType","optParams":{"methodName":"SmallExercise"}}
func (thisObj *structType) SmallExercise(){
	//主要考察:切片和map都是引用类型,以值方式传递给方法中的形参(相当于指针传递),方法中进行修改,外部再调用时也是修改后的值!

	//切片-小练习
	{
		testArr := []int{1,2,3}
		testFunc := func(testArr []int){
			testArr[1] = 999
		}

		//输出调试:
		fmt.Println("testArr_before:",testArr)
		testFunc(testArr)
		fmt.Println("testArr_after:",testArr)
		//输出结果：
		//testArr_before: [1 2 3]
		//testArr_after: [1 999 3]
	}

	//切片值结构体-小练习
	{
		type testStruct struct {
			Name string
			Age int
		}

		testArr := []testStruct{
			{Name:"test1",Age:21},
			{Name:"test2",Age:22},
			{Name:"test3",Age:23},
		}
		testFunc := func(testArr []testStruct){
			testArr[1].Age = 999	//@todo 这里有个疑问,为什么切片直接使用 struct(值结构体) 可以修改,而map是不行的
		}

		//输出调试:
		fmt.Println("testArr_before:",testArr)
		testFunc(testArr)
		fmt.Println("testArr_after:",testArr)
		//输出结果:
		//testArr_before: [{test1 21} {test2 22} {test3 23}]
		//testArr_after: [{test1 21} {test2 999} {test3 23}]
	}


	fmt.Println("~~~~~~~~~~~~ 切片指针结构体-小练习 start ~~~~~~~~~~~~")
	{
		type testStruct struct {
			Name string `json:"name"`
			Age int	`json:"age"`
		}
		testArr := []*testStruct{
			{Name:"test1",Age:21,},
			{Name:"test2",Age:22,},
			{Name:"test3",Age:23,},
		}
		testFunc := func(testArr []*testStruct){
			testArr[1].Age = 999
		}
		//输出调试:
		testArrByte,_ := json.Marshal(testArr)
		fmt.Println("testArr_before:",string(testArrByte))

		testFunc(testArr)

		testArrByte,_ = json.Marshal(testArr)
		fmt.Println("testArr_after:",string(testArrByte))
		//输出结果:
		//~~~~~~~~~~~~ 切片指针结构体-小练习 start ~~~~~~~~~~~~
		//testArr_before: [{"name":"test1","age":21},{"name":"test2","age":22},{"name":"test3","age":23}]
		//testArr_after: [{"name":"test1","age":21},{"name":"test2","age":999},{"name":"test3","age":23}]
		//~~~~~~~~~~~~ 切片指针结构体-小练习 end ~~~~~~~~~~~~
	}
	fmt.Println("~~~~~~~~~~~~ 切片指针结构体-小练习 end ~~~~~~~~~~~~")


	//map-小练习
	{
		testMap := map[string]int{
			"test1":21,
			"test2":22,
			"test3":23,
		}
		testFunc := func(testMap map[string]int){
			testMap["test2"] = 999
		}

		//输出调试:
		fmt.Println("testMap_before:",testMap)
		testFunc(testMap)
		fmt.Println("testMap_after:",testMap)
		//输出结果:
		//testMap_before: map[test1:21 test2:22 test3:23]
		//testMap_after: map[test1:21 test2:999 test3:23]
	}

	//map值结构体-小练习
	{
		type testStruct struct {
			Name string
			Age int
		}

		testMap := map[string]testStruct{
			"test1":{Name:"test1",Age:21},
			"test2":{Name:"test2",Age:22},
			"test3":{Name:"test2",Age:23},
		}
		testFunc := func(testMap map[string]testStruct){
			//如下这样的赋值操作是无法进行的,原因:分配的左侧不是可寻址的值.
			//	1.map[xxx]试图直接修改struct属性
			// 	2.interface{}断言里试图直接修改struct属性
			//	@todo 3.针对结构体最好是以指针形式进行创建和修改

			//无法进行赋值操作
			//testMap["test2"].Age = 999	//@todo 注意:这里由于不是指针结构体,不能直接map[xxx].Age直接修改struct字段值

			//这样先获取,确定好可寻址,再进行修改完毕后覆盖回去
			testStructObj := testMap["test2"]
			testStructObj.Age = 999
			testMap["test2"] = testStructObj
		}

		//输出调试:
		fmt.Println("testMap_before:",testMap)
		testFunc(testMap)
		fmt.Println("testMap_after:",testMap)
		//输出结果:
		//testMap_before: map[test1:{test1 21} test2:{test2 22} test3:{test2 23}]
		//testMap_after: map[test1:{test1 21} test2:{test2 999} test3:{test2 23}]
	}

	fmt.Println("~~~~~~~~~~~~ map指针结构体-小练习 start ~~~~~~~~~~~~")
	{
		type testStruct struct {
			Name string `json:"__name__"`
			Age int	`json:"__age__"`
		}

		testMap := map[string]*testStruct{
			"test1":&testStruct{Name:"test1",Age:21,},	//不用加&testStruct,如果map声明时是指针结构体,初始化赋值时就是以指针结构体赋值了
			"test2":{Name:"test2",Age:22,},
			"test3":{Name:"test3",Age:23,},
		}
		testFunc := func(testMap map[string]*testStruct){
			testMap["test2"].Age = 999	//@todo 注意:这里因为是指针结构体,map[xxx].Age可以直接修改struct字段值
		}

		//输出调试:
		testMapByte,_ := json.Marshal(testMap)
		fmt.Println("testMap_before:",string(testMapByte))

		testFunc(testMap)

		testMapByte,_ = json.Marshal(testMap)
		fmt.Println("testMap_after:",string(testMapByte))

		//输出结果:
		//~~~~~~~~~~~~ map指针结构体-小练习 start ~~~~~~~~~~~~
		//testMap_before: {"test1":{"__name__":"test1","__age__":21},"test2":{"__name__":"test2","__age__":22},"test3":{"__name__":"test3","__age__":23}}
		//testMap_after: {"test1":{"__name__":"test1","__age__":21},"test2":{"__name__":"test2","__age__":999},"test3":{"__name__":"test3","__age__":23}}
		//~~~~~~~~~~~~ map指针结构体-小练习 end ~~~~~~~~~~~~
	}
	fmt.Println("~~~~~~~~~~~~ map指针结构体-小练习 end ~~~~~~~~~~~~")

}