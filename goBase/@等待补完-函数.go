package goBase

//######################################################################################################################
//GO-函数Func
//	GO-关于函数Func的说明:
//	  参考文章地址:
//	  http://topgoer.com/%E5%87%BD%E6%95%B0/%E5%8F%82%E6%95%B0.html
//----------------------------------------------------------------------------------------------------------------------

/* 1.golang函数特点 */
//----------------------------------------------------------------------------------------------------------------------
//		a.无需声明原型。
// 			直接定义和直接调用,也是不区分代码位置.
//			示例:
//			//函数定义前-直接调用
//			test()
//			//定义函数
// 			func test()int{
// 				return 123
// 			};
// 			//函数定义后-直接调用
// 			test()
//
//		b.支持不定变参。
// 			示例:
// 			func test(args ...interface{}){
// 				fmt.Println("args:",args)
// 			}
//			//传参示例1:
//			test(1,"2",[]int{3},map[int]int{4:4})
//			//传参示例2:
//			// 示例2中的如果按切片值打散传入,切片的数据类型要统一,另外也不可在切片打散语法前后再添加参数传参
//			test([]interface{}{1,2,3,4}...)	//正确传参
//			//test(111,222,[]interface{}{1,2,3,4}...,333,444)	//错误传参
//
//			//输出结果:
//			//args: [1 2 [3] map[4:4]]
//			//args: [1 2 3 4]
//
//		c.支持多返回值。
// 			定义函数返回值时可设置多个值返回,且返回值数据类型不限制统一。
// 			示例:
// 			func test()(int,error){
// 				return 123,errors.New("test")
// 			}
//
//		d.支持命名返回参数。
//			1.返回参数可声明变量,相当于返回参数变量声明,但对部分数据类型没有执行初始化开辟内存空间存储的操作!
//			2.一旦声明返回参数变量,函数体到最后执行完毕时要有return存在.
//			示例:
//			func test() (testMap map[string]interface{}){
//				testMap = make(map[string]interface{})	//针对需要map之类的需要初始化操作!
//				testMap["test"] = 1
//				return	//声明返回参数变量后要有return
//			}
//
//			res := test()
//			fmt.Println(res)
//			//输出结果:
//			//map[test:1]
//
//		e.支持匿名函数和闭包。
//			1.匿名函数和匿名结构体一样
//				示例:
//				func(){
// 					fmt.Println("this is 匿名函数")
// 				}()
//			2.闭包:简而言之，闭包的作用就是在test函数执行完并返回后，闭包使得垃圾回收机制不会收回test函数所占用的资源，因为test函数的内部函数的执行需要依赖test函数中的变量a。
//				func test() func() int {
//					a := 0
//					return func() int {
//						a++
//						return a
//					}
//				}
//				test1 := test()
//				fmt.Println(test1())
//				fmt.Println(test1())
//				//输出结果:
//				//1
//				//2
//
//		f.函数也是一种类型，一个函数可以赋值给变量。
//			示例1:
//			var testFunc func(tag int)int
//			testFunc = func(tag int)int{
//				tag++
//				return tag
//			}
//			res := testFunc(1)
//			fmt.Println(res)	//输出:2
//
//			示例2:
//			testFunc := func(tag int)int{
//				tag++
//				return tag
//			}
//			res := testFunc(1)
//			fmt.Println(res)	//输出:2
//
//		g.不支持 嵌套 (nested) 一个包不能有两个名字一样的函数。
//		h.不支持 重载 (overload)
//		i.不支持 默认参数 (default parameter)。
//----------------------------------------------------------------------------------------------------------------------

/* 2.函数参数 */
//----------------------------------------------------------------------------------------------------------------------
//	1.传递方式说明:
//		(1).2种传递方式
//			值传递：指在调用函数时将实际参数复制一份传递到函数中，这样在函数中如果对参数进行修改，将不会影响到实际参数。
//			引用传递：是指在调用函数时将实际参数的地址传递到函数中，那么在函数中对参数所进行的修改，将影响到实际参数。
//
//		(2).默认情况下，Go 语言使用的是值传递，即在调用过程中不会影响到实际参数。
//
//		(3).注意:
//				a.无论是值传递，还是引用传递，传递给函数的都是变量的副本。
//				  不过，值传递是值的拷贝，引用传递是地址的拷贝。一般来说，地址拷贝更为高效。而值拷贝取决于拷贝的对象大小，对象越大，则性能越低。
//		@todo   b.map、slice、chan、指针、interface默认以引用的方式传递。(interface有可能有歧义,测试代码是没有测出interface是有引用效果!)
//
//		(4).不定参数传值 就是函数的参数不是固定的，后面的类型是固定的。（可变参数）
//			要么是统一数据类型的多个参数一起传,要么就指定多个不同数据类型传参。
//
//		(5).Golang 可变参数本质上就是 slice。只能有一个，且必须是最后一个。
//			如 func test(name string,age int,args ...int){
// 				args[0] //第一个参数值
// 				args[1] //第二个参数值
// 			}
//
//		(6).在参数赋值时可以不用用一个一个的赋值，可以直接传递一个数组或者切片，特别注意的是在参数后加上"…"即可。
//			一个一个的赋值是针对多个指定的不同数据类型,如果是统一数据类型可以以可变参数传入。
//			其实一个map[string]interface{]就解决大部分传参难题
//
//		(7).任意类型的不定参数： 就是函数的参数和每个参数的类型都不是固定的。
// 			用interface{}传递任意类型数据是Go语言的惯例用法，而且interface{}是类型安全的。
//			如 func test(name string,age int,args ...interface{}){
// 					args[0].(string) 	//第一个参数值数据类型是字符串
//					args[1].(int) 		//第二个参数值数据类型是整型
//					...
//					args[len(args)-2].(*testStrict).Age 		//倒数第二个参数值数据类型是指针结构体
//					args[len(args)-1].(map[string]interface{}) 	//最后一个参数值数据类型是map
// 			}
//
//		(8).使用 slice 做变参时,如果是直接以 slice 为一个整体参数传入,就不需要...展开传入;
//			如果想让 slice 每个值都作为一个独立参数传入,就使用...展开传入.
//			示例:
//			testSlice := []interface{}{1,2,3}
//			testFunc := func(args ...interface{}){
//				fmt.Println("args:",args,"argsLen:",len(args))
//			}
//			testFunc(testSlice)		//slice 自身作为【一个】独立参数传入
//			testFunc(testSlice...)	//slice 【每个】值作为一个独立参数传入
//			//输出结果:
//			//args: [[1 2 3]] argsLen: 1
//			//args: [1 2 3] argsLen: 3
//----------------------------------------------------------------------------------------------------------------------

/* 3.函数返回值 */
//----------------------------------------------------------------------------------------------------------------------
//	挑重点说明下:
//
//		@todo return //没有表达式(凡是有值的都叫表达式)在return右边的,叫隐式返回
//		@todo return xxx 或 return 变量+1 或 return func(){}...等 都叫显示返回
//		@todo 注意:这里有一个特殊性,显示返回时,若有指定命名返回参数的修改操作,则会针对命名返回参数代码进行执行再return或者defer执行完毕后return!
//
//		a.命名返回参数可被同名局部变量遮蔽，此时需要显式返回。
//			示例:
//			testFunc := func()(res int){
//				{
//					var res = 1
//					//return			//直接使用return返回,会报错res redeclared in this block
//					return res			//必须这里使用显示(即指定变量)返回
//
//				}
//				return res				//上面花括号已经使用return返回了,这里不会执行到
//			}
//
//			fmt.Println( testFunc() )	//输出是1,而不是0
//
//		b.命名返回参数允许 defer 延迟调用通过闭包读取和修改。
//			示例1:
//			testFunc := func() (res int) {
//
//				defer func(){		//第二步:在return之前,执行defer里的方法体;此时res为1,+=100后,res为101
//					res += 100
//				}()
//
//				res++				//第一步:函数体内先执行res++;此时res默认为0,++后,res为1
//
//				return				//第三步:返回res;此时res为101
//			}
//
//			fmt.Println( testFunc() )	//输出:101
//
//			示例2:@todo 注意这个示例,关于函数返回值+defer处理流程
//			testFunc := func() (res int) {
//
//				fmt.Println("func-body:",res)				//第一步:输出 func-body: 0
//
//				defer func(){
//					fmt.Println("defer-res-before:",res)	//第三步:输出 并 res+=100
//
//					res += 100
//
//					fmt.Println("defer-res-after:",res)
//				}()
//
//				return	res+1								//第二步:return之前进行代码运算res+1,将计算后的值重新赋给res,此时res值为1;
// 															//再接着进行第三步defer,第三步defer执行完毕后;
// 															//第四步最终将res的值返回给外部!
//			}
//
//			fmt.Println( testFunc() )
//			//输出结果:
//			//func-body: 0
//			//defer-res-before: 1
//			//defer-res-after: 101
//			//101
//
//		c.显式 return 返回前，会先修改命名返回参数。
//			示例:
//			testFunc := func() (res int,res1 map[string]interface{}) {
//
//				defer func(){
//					fmt.Println("this is testFunc() defer:")
//					fmt.Println(res,res1)
//				}()
//
//				//参考:return z + 200 // 执行顺序: (z = z + 200) -> (call defer) -> (return)
//				return res+2,map[string]interface{}{	//@todo 注意:这里的return会针对res,res1对应命名返回参数进行最后结果修改
//					"test":"this is res1",
//				}
//			}
//
//			res,res1 := testFunc()
//
//			//输出调试:
//			fmt.Println("this is testFunc() return result:")
//			fmt.Println(res,res1)
//			//输出结果:
//			//this is testFunc() defer:
//			//2 map[test:this is res1]
//			//this is testFunc() return result:
//			//2 map[test:this is res1]
//----------------------------------------------------------------------------------------------------------------------

/* 4.匿名函数 */
//----------------------------------------------------------------------------------------------------------------------
//	@todo 匿名函数可赋值给变量,slice,map,结构字段,或者在 channel 里传送。
//
//		a.匿名函数赋给变量
//			示例:
//			testFunc := func(){
//				fmt.Println("123")
//			}
//			testFunc() //输出:123
//
//		b.匿名函数作为slice值
//			示例:
//			testFunc := func(tag int,name string) (resTag int,resName string, resData map[string]interface{}){
//				fmt.Println("testFunc-tag:",tag,"name:",name)
//				return
//			}
//
//			//这里声明的时候指定不指定 参数变量名、返回参数变量名 都没有约束力度,在声明方法的参数变量名时,变量名不一样也没问题
//			//testFuncArr := []func(int,string)(int,string,map[string]interface{}){}
//			//上面与下面的声明效果一致!
//			testFuncArr := []func(tag11 int,name22 string)(resTag int,resName string,resData map[string]interface{}){}
//
//			testFuncArr = append(testFuncArr,testFunc)
//
//			for k,v := range testFuncArr {
//				resTag,resName,resData := v(k,"test")
//				fmt.Println("forRange-resTag:",resTag,"resName:",resName,"resData:",resData)
//			}
//
//			//输出结果:
//			//testFunc-tag: 0 name: test
//			//forRange-resTag: 0 resName:  resData: map[]
//
//		c.匿名函数作为map值
//			示例:
//			testFuncMap := map[string]func(int,string)(int,string,map[string]interface{}){
//				"testFunc":func(tag int,name string)(resTag int,resName string,resData map[string]interface{}){
//					fmt.Println("tag:",tag,"name:",name)
//					return
//				},
//			}
//			testFuncMap["testFunc"](21,"TestTest")
//			//输出结果:
//			//tag: 21 name: TestTest
//
//		d.匿名函数作为 结构字段 值
//			示例:
//			type testStruct struct {
//				testFunc func(int,string)(int,string,map[string]interface{})
//			}
//
//			//var TestStruct testStruct
//			//TestStruct.testFunc = func(age int,name string)(resAge int,resName string,resData map[string]interface{}){
//			//	fmt.Println("age:",age,"name:",name)
//			//	return
//			//}
//			TestStruct := testStruct{
//			testFunc:func(age int,name string)(resAge int,resName string,resData map[string]interface{}){
//					fmt.Println("age:",age,"name:",name)
//					return
//				},
//			}
//			TestStruct.testFunc(21,"aaa")
//
//			//输出结果:
//			//age: 21 name: aaa
//
//		e.匿名函数作为 channel 值
//			示例：
//			//创建一个通道,该通道是缓冲的(注:若是无缓冲的通道,则会形成通道死锁!)
//			c := make(chan func(int,string)(int,string,map[string]interface{}),1)
//
//			//写入通道值
//			c <- func(age int,name string) (resAge int,resName string,resData map[string]interface{}) {
//				fmt.Println("age:",age,"name:",name)
//				return
//			}
//
//			//获取通道值
//			testFunc := <- c
//
//			testFunc(21,"asd")
//
//			//输出结果:
//			//age: 21 name: asd
//----------------------------------------------------------------------------------------------------------------------

/* 5.闭包、递归 */
//----------------------------------------------------------------------------------------------------------------------
//	参考文章 《闭包、递归》 说明.
//
//	a.闭包
//		示例-外部引用函数参数局部变量:
//		testFunc := func(age int) func(int) int {
//			return func(i int) int {
//				age += i
//				return age
//			}
//		}
//		subFunc := testFunc(5)
//		fmt.Println( subFunc(1) )	//输出:6
//		fmt.Println( subFunc(1) )	//输出:7
//
//	b.递归-函数内部调用自己
//		示例:
//		func testFunc(i int){
//			fmt.Println("i:",i)
//
//			if i--;i>0 {
//				testFunc(i)
//			}else{
//				fmt.Println("--------------------")
//			}
//
//			fmt.Println("i:",i+1)
//		}
//		testFunc(5)
//		//输出结果:
//		//i: 5
//		//i: 4
//		//i: 3
//		//i: 2
//		//i: 1
//		//--------------------
//		//i: 1
//		//i: 2
//		//i: 3
//		//i: 4
//		//i: 5
//----------------------------------------------------------------------------------------------------------------------

/* 6.defer延迟调用 */
//----------------------------------------------------------------------------------------------------------------------
//	a.defer特性：
//    	(1).关键字 defer 用于注册延迟调用。
//    	(2).这些调用直到 return 前才被执。因此，可以用来做资源清理。
//    	(3).多个defer语句，按先进后出的方式执行。
//    	(4).defer语句中的变量，在defer声明时就决定了。
//
//	b.defer用途:
//		(1).关闭文件句柄
//		(2).锁资源释放
//		(3).数据库连接释放
//		(4).异常捕获
//
//	c.defer是先进后出 且是 代码执行到defer才会被注册
//		注意:后面的语句会依赖前面的资源,因此如果先前面的资源先释放了,后面的语句就没法执行了。
//
//		示例:
//		func testFunc(i int){
//			if i<=0 {
//				fmt.Println("this is i elt 0!")
//				return
//			}
//			//这个defer没有被执行,是因为上面的return
//			defer func(){
//				fmt.Println("this is defer!")
//			}()
//		}
//		testFunc(0)
//		//输出结果:
//		//this is i elt 0!
//
//	d.defer 碰上闭包
//	//@todo 等待补完!
//----------------------------------------------------------------------------------------------------------------------