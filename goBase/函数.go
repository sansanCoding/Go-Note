package goBase

//######################################################################################################################
//GO-函数Func
//	GO-关于函数Func的说明:
//	  参考文章地址:
//	  http://topgoer.com/%E5%87%BD%E6%95%B0/%E5%8F%82%E6%95%B0.html
//######################################################################################################################

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
//		defer后面的语句在执行的时候，函数调用的参数会被保存起来，但是不执行。也就是复制了一份。
// 		但是并没有说struct这里的this指针如何处理，通过这个例子可以看出go语言并没有把这个明确写出来的this指针当作参数来看待。
//
//		示例:
//		testFunc := func(){
//			for i:=0;i<3;i++ {
//				//注意:等循环执行完毕,函数体结束之前,执行defer的时候,i这个外部变量最后的值就是3(i=3时才退出循环,等到defer执行的时候i就是3了)
//				defer func(){
//					fmt.Println("i:",i)	//注意:这里的i是闭包引用
//				}()
//			}
//		}
//		testFunc()
//		//输出结果:
//		//i: 3
//		//i: 3
//		//i: 3
//
//		//解决defer闭包的外部变量更新:
//@todo //1.一种是匿名函数前,声明临时变量,重置外部变量
//		//示例如下:
//			testFunc1 := func(){
//				for i:=0;i<3;i++ {
//					//一种是临时声明一个变量,刷新匿名函数里的外部变量
//					i:=i
//					defer func(){
//						fmt.Println("i:",i)
//					}()
//				}
//			}
//			testFunc1()
//			//输出结果:
//			//i: 2
//			//i: 1
//			//i: 0
//@todo //2.一种是以实参传递给函数形参处理,这种处理逻辑清晰,容易理解,方便维护
//		//示例如下:
//			testFunc2 := func(){
//				for i:=0;i<3;i++ {
//					//这个是输出是告知当时i是什么值,等defer执行时就是什么值(这个涉及defer陷阱)
//					defer fmt.Println("i:",i)	//注意:这里的i被复制
//					//一种是相当于函数传参,推荐这种,逻辑清晰,容易理解,方便维护
//					defer func(data int){
//						fmt.Println("data:",data)
//					}(i)	//注意:这里的i被复制
//				}
//			}
//			testFunc2()
//			//输出结果:
//			//data: 2
//			//i: 2
//			//data: 1
//			//i: 1
//			//data: 0
//			//i: 0
//
//	e.defer 陷阱
//		(1).defer 与 closure 的传参-复制和传参-闭包引用,造成的结果会不同
//			示例:
//			testFunc := func() (err error) {
//				//这个defer执行时,fmt.Println函数将当时的err复制保存一份,相当于当时传参值是什么,到时候会输出什么
//				defer fmt.Println("err:",err)	//err被复制
//				//这个defer执行时,也是和上面的一样,将当时的err复制保存一份,相当于当时传参值是什么,到时候会输出什么
//				defer func(tempErr error){
//					fmt.Println("funcErr:",tempErr)
//				}(err)	//err被复制
//				//这个defer执行时,由于err是闭包引用的值,等被执行时,err就是最新修改的值
//				defer func(){
//					fmt.Println("deferErr:",err)	//err 闭包引用
//				}()
//
//				err = errors.New("test_error")
//				return
//			}
//			testFunc()
//			//输出结果:
//			//deferErr: test_error
//			//funcErr: <nil>
//			//err: <nil>
//
//		(2).defer 与 return 的具名返回值处理流程 以及 defer在return之后的处理流程
//			总结:
// 			//@todo 1.若函数存在具名返回值,在函数return之前,其次在defer之前,先将return的表达式处理结果赋值给具名返回值;
//			//@todo   若defer存在且也使用了具名返回值时,则是最新赋值操作的结果;
//			//@todo	  若defer存在且也对具名返回值进行赋值操作,则最后return返回的具名返回值就是在defer(注意defer可以执行多次)里的最后一次赋值操作结果;
//			//@todo 2.defer只要在return之前,都会挨个记录,到时候以先进后出的方式,挨个执行;
//			//@todo   但凡遇到return之后的defer,都不会被记录和执行;这一点提现在判断里的返回;
//				  如
//					func(a int){
//						fmt.Println("a:",a)
//						if a==1 {
//							return
//						}
//						defer fmt.Println("我是return之后的defer,我是不会被执行的")
//					}(1)
//
//			示例1:
//			testFunc := func() (actionResult int) {
//				//1.先将具名返回值actionResult赋值为1
//				actionResult = 1
//
//				//2.actionResult是闭包引用的值,所以在return之前如果还有最新赋值操作,则defer的actionResult输出最新赋值结果
//				defer func(){
//					fmt.Println("actionResult:",actionResult)
//				}()
//
//				//3.这里的return返回的表达式,如果没有具名返回值,则是直接返回;
//				//  一旦有具名返回值,会在return之前和defer之前,先将该表达式处理的结果赋值给具名返回值;
//				//	等defer再用到具名返回值是,具名返回值则是最新赋值结果.
//				//return 22	//输出结果 actionResult: 22
//				//return 10*22	//输出结果 actionResult: 220
//				//return actionResult+22 //输出结果 actionResult: 23
//
//				//return actionResult+=22)					//这一种还有下面一种都是语法编译不过去的
//				//return actionResult = actionResult +22
//				//return func() int { return actionResult+22 }()	//当然这种匿名函数写法也是actionResult闭包引用+22返回的结果 //输出结果 actionResult: 23
//
//				//由此可以看出,defer是被程序运行时添加到一种类似排队效果的处理,而不是在语法定义层面就会被记录,如func(),可在函数声明前调用!!!
//				//	简单来说,defer是没遇到return之前,会被记录下来,到时候统一在return之前挨个执行,执行顺序就是先进后出的方式!
//				defer func(){
//					fmt.Println("~~~我这个defer是在return之后,所以是不会被执行的~~~")
//				}()
//				return	//这个return是因为上面加了defer,不写return会形成语法错误,编译不过去
//			}
//			testFunc()
//
//			示例2:
//			//若defer存在且也对具名返回值进行赋值操作,则return返回的具名返回值是defer最后一次赋值操作的结果
//			res := func() (actionResult string) {
//				actionResult = "a"			//1.先将具名返回值赋值为a
//				defer func(){
//					actionResult += "b"		//4.defer先进后出的顺序,将具名返回值拼接d后再隐性赋值会给具名返回值
//				}()
//				defer func(){
//					actionResult += "c"		//3.defer先进后出的顺序,将具名返回值拼接d后再隐性赋值会给具名返回值
//				}()
//
//				return actionResult+"d"		//2.将具名返回值拼接d后再隐性赋值会给具名返回值
//			}()
//			fmt.Println("res:",res)
//			//输出结果：
//			//res: adcb
//
//		(3).defer nil 函数
//			示例:
//			func(){
//				//不过从开发角度讲,应该不会这么干,即使真要真没干,估计要提前判断下
//				var testFunc func() = nil	//testFunc声明数据类型为func(),值为nil
//				defer testFunc()
//				fmt.Println("this is test")	//先输出在抛出异常
//
//				//输出结果:
//				//this is test
//				//紧接着抛出异常:
//				//runtime error: invalid memory address or nil pointer dereference
//			}()
//			改进版:
//			func(){
//				var testFunc func() = nil
//				if testFunc!=nil {	//若不为nil值,则再介入defer执行
//					defer testFunc()
//				}
//				fmt.Println(testFunc,reflect.TypeOf(testFunc).String())	//输出结果:<nil> func()
//				fmt.Println("this is test")	//输出结果:this is test
//			}()
//
//		(4).在错误的位置使用 defer
//			如文章上的示例:
//			//错误的调用
//			res, err := http.Get("http://www.google.com")
//			defer res.Body.Close()
//			if err != nil {
//				return err
//			}
//			//改进错误的调用
//			res, err := http.Get("http://xxxxxxxxxx")
//			if res != nil {
//				defer res.Body.Close()
//			}
//			if err != nil {
//				return err
//			}
//			//-------------------------------------------
//			从上面的文章示例来看,主要是针对资源性的如http请求,mysql数据库链接等;
// 			发起请求后,紧接着使用defer关闭,这样一来就是在未确定是否有错误,defer再执行有错误的资源关闭,则会抛出异常!
//			但是从开发角度上讲,不可能先不判断错误,而直接调用defer关闭资源的说!!!
//			//开发角度修改版-万金油处理
//			res, err := http.Get("http://xxxxxxxxxx")
//			if err != nil {
//				return err
//			}
//			//当然这里加上这个判断也可以,但是想来想去,既然没有err错误返回,res怎么可能还为nil呢?
//			if res != nil {
//				defer res.Body.Close()
//			}
//			//开发角度修改版-正常处理
//			res, err := http.Get("http://xxxxxxxxxx")
//			if err != nil {
//				return err
//			}
//			defer res.Body.Close()
//
//		(5).不检查错误(即defer直接调用的函数,若返回错误,而没有做处理)
//			示例:
//			func(){
//				testFunc := func() error {
//					return errors.New("testError!")
//				}
//				defer testFunc()	//这一步的返回的错误就直接忽略了,当然,这个根据实际情况确定,如果确定不需要处理该错误,忽略就忽略了!
//				fmt.Println("this is test finish!")
//				//输出结果:
//				//this is test finish!
//			}()
//
//		(6).释放相同的资源
//			释放相同的资源,文章上的示例是多个defer但是都是同一个闭包引用的变量进行处理,这样造成的是最后一个资源被释放正确,
// 			但是defer再执行第二次的时候就会报错;解决方案有:
//				第一种 实时传参(即以传参-复制进行处理)
//				第二种 不同的变量命名进行对应释放处理即可
//----------------------------------------------------------------------------------------------------------------------

/* 7.异常处理 */
//----------------------------------------------------------------------------------------------------------------------
//	a.GoLang异常说明
// 		Golang 没有结构化异常，使用 panic 抛出错误，recover 捕获错误。
//
//	b.异常抛出及捕获流程-简单描述:
//		Go抛出一个panic的异常，然后在defer中通过recover捕获这个异常。
//
//	c.异常抛出-panic 与 异常捕获-recover
// 		异常抛出-panic：
//		1、内置函数
//		2、假如函数F中书写了panic语句，会终止其后要执行的代码，在panic所在函数F内如果存在要执行的defer函数列表，按照defer的逆序执行
//		3、返回函数F的调用者G，在G中，调用函数F语句之后的代码不会执行，假如函数G中存在要执行的defer函数列表，按照defer的逆序执行
//		4、直到goroutine整个退出，并报告错误
//
//		异常捕获-recover:
// 		1、内置函数
//    	2、用来控制一个goroutine的panicking行为，捕获panic，从而影响应用的行为
//    	3、一般的调用建议
//        a). 在defer函数中，通过recever来终止一个goroutine的panicking过程，从而恢复正常代码的执行
//        b). 可以获取通过panic传递的error
//
//
//		由于 panic、recover 参数类型为 interface{}，因此可抛出任何类型对象。
//		func panic(v interface{})
//		func recover() interface{}
//
//
//		注意:
//		1.利用recover处理panic指令，defer 必须放在 panic 之前定义，另外 recover 只有在 defer 调用的函数中才有效。否则当panic时，recover无法捕获到panic，无法防止panic扩散。
//		2.recover 处理异常后，逻辑并不会恢复到 panic 那个点去，函数跑到 defer 之后的那个点。
//		3.多个 defer 会形成 defer 栈，后定义的 defer 语句会被最先调用。
//
//	====================================================================================================================
//	1.简单示例:
//		func(){
//			defer func(){
//				if err:=recover();err!=nil {
//					fmt.Println("recoverErr:",reflect.TypeOf(err).String(),err)
//				}
//			}()
//			//抛出string
//			panic("this is panic test!")
//		}()
//
//		func(){
//			defer func(){
//				if err:=recover();err!=nil {
//					fmt.Println("recoverErr:",reflect.TypeOf(err).String(),err)
//				}
//			}()
//			//抛出map等其他的数值类型(不过一般map或string就足够了)
//			panic(map[string]interface{}{
//				"test":"this is panic map!",
//			})
//		}()
//
//		//输出结果:
//		//recoverErr: string this is panic test!
//		//recoverErr: map[string]interface {} map[test:this is panic map!]
//
//	2.异常抛出时,多个defer里,按先进后出的顺序,只要排前defer有异常捕获且也只会被捕获一次
//	示例:
//		func(){
//			defer func(){
//				fmt.Println("fmtPrintRecover3")
//				fmt.Println("fmtPrintRecover:",recover())	//按后进先出,由于之前defer已捕获异常,这个defer则不会捕获到异常
//			}()
//			defer func(){
//				fmt.Println("fmtPrintRecover2")
//			}()
//			defer func(){
//				fmt.Println("fmtPrintRecover1")
//				fmt.Println("fmtPrintRecover:",recover())	//按后进先出,这个defer会捕获到异常
//			}()
//			panic("test panic!")
//		}()
//	//输出结果:
//	//fmtPrintRecover1
//	//fmtPrintRecover: test panic!
//	//fmtPrintRecover2
//	//fmtPrintRecover3
//	//fmtPrintRecover: <nil>
//
//	3.产生异常抛出时,多个defer只要被执行到,就会进入先进后出的顺序,挨个执行;对,不管defer里是否还会抛出异常,多个defer都会按顺序执行!
//	示例:
//		func(){
//			defer func(){
//				fmt.Println("fmtPrintRecover3")
//				fmt.Println("fmtPrintRecover:",recover())
//			}()
//			defer func(){
//				fmt.Println("fmtPrintRecover2")
//			}()
//			defer func(){
//				fmt.Println("fmtPrintRecover1")
//				//fmt.Println("fmtPrintRecover:",recover())
//			}()
//			panic("test panic!")
//		}()
//	//输出结果:
//	//fmtPrintRecover1
//	//fmtPrintRecover2
//	//fmtPrintRecover3
//	//fmtPrintRecover: test panic!
//
//	4.产生异常抛出时,若存在多个defer里,若某一个defer产生了异常,可被后续靠前的defer进行捕获,且只会捕获到最后一个错误
//	示例:
//		func(){
//			defer func(){
//				fmt.Println("fmtPrintRecover3")
//				fmt.Println("fmtPrintRecover:",recover())
//			}()
//			defer func(){
//				fmt.Println("fmtPrintRecover2")
//			}()
//			defer func(){
//				fmt.Println("fmtPrintRecover1")
//				panic("this is fmtPrintRecover1PanicError!")
//			}()
//			panic("test panic!")
//		}()
//	//输出结果:
//	//fmtPrintRecover1
//	//fmtPrintRecover2
//	//fmtPrintRecover3
//	//fmtPrintRecover: this is fmtPrintRecover1PanicError!
//
//	5.产生异常抛出时,若存在多个defer里,只要靠前defer有捕获异常,函数体的异常抛出还是defer的异常抛出,都能捕获到,哪怕defer里再抛出异常,只要后面defer里有捕获异常的,照样进行捕获处理!
//	示例:
//		func(){
//			defer func(){
//				fmt.Println("fmtPrintRecover3")
//				fmt.Println("fmtPrintRecover3:",recover())
//			}()
//			defer func(){
//				fmt.Println("fmtPrintRecover2")
//			}()
//			defer func(){
//				fmt.Println("fmtPrintRecover1")
//				panic("this is fmtPrintRecover1PanicError!")
//			}()
//			defer func(){
//				fmt.Println("fmtPrintRecover0")
//				fmt.Println("fmtPrintRecover0:",recover())
//			}()
//			panic("test panic!")
//		}()
//	//输出结果:
//	//fmtPrintRecover0
//	//fmtPrintRecover0: test panic!
//	//fmtPrintRecover1
//	//fmtPrintRecover2
//	//fmtPrintRecover3
//	//fmtPrintRecover3: this is fmtPrintRecover1PanicError!
//
//	6.捕获函数 recover 只有在延迟调用内直接调用才会终止错误，否则总是返回 nil。任何未捕获的错误都会沿调用堆栈向外传递。
//	示例:
//		func(){
//			defer func(){
//				fmt.Println("recoverError:",recover())	//捕获有效!
//			}()
//
//			defer recover()	//捕获无效!!!
//
//			defer fmt.Println("fmtRecoverError:",recover())	//捕获无效!!!
//
//			defer func(){
//				func(){
//					fmt.Println("funcFuncRecoverError:",recover()) //捕获无效!!!
//				}()
//			}()
//
//			panic("this is test panic!")
//		}()
//	//输出结果:
//	//funcFuncRecoverError: <nil>
//	//fmtRecoverError: <nil>
//	//recoverError: this is test panic!
//----------------------------------------------------------------------------------------------------------------------

/* 8.错误处理 */
//----------------------------------------------------------------------------------------------------------------------
//	a.除用 panic 引发中断性错误外，还可返回 error 类型错误对象来表示函数调用状态。
//	error语法:
//		type error interface {
//			Error() string
//		}
//
//	b.标准库 errors.New 和 fmt.Errorf 函数用于创建实现 error 接口的错误对象。通过判断错误对象实例来确定具体错误类型。
//		示例:
//		func(){
//			err := fmt.Errorf("errorTag => %d",10)
//			fmt.Println("err:",reflect.TypeOf(err).String(),"======",err)
//
//			err1 := errors.New("this is errors!")
//			fmt.Println("err1:",reflect.TypeOf(err1).String(),"======",err1)
//
//			//输出结果:
//			//err: *errors.errorString ====== errorTag => 10
//			//err1: *errors.errorString ====== this is errors!
//		}()
//
//----------------------------------------------------------------------------------------------------------------------

/* 9.Go实现类似 try catch 的异常处理 */
//----------------------------------------------------------------------------------------------------------------------
//	示例:
//		try := func(tryFunc func(),catchFunc func(interface{})){
//			defer func(){
//				if err:=recover(); err!=nil {
//					catchFunc(err)
//				}
//			}()
//			tryFunc()
//		}
//
//		try(func(){
//			fmt.Println("12321")
//			panic("this is try catch panic!")
//		},func(err interface{}){
//			fmt.Println("err:",err)
//		})
//	输出结果:
//	//12321
//	//err: this is try catch panic!
//----------------------------------------------------------------------------------------------------------------------

/* 10.如何区别使用 panic 和 error 两种方式? */
//	惯例是:导致关键流程出现不可修复性错误的使用 panic，其他使用 error。