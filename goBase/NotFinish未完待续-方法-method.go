package goBase

//######################################################################################################################
//GO-方法Method
//	GO-关于方法Method的说明:
//	  参考文章地址:
//	  http://topgoer.com/%E6%96%B9%E6%B3%95/%E6%96%B9%E6%B3%95%E5%AE%9A%E4%B9%89.html
//######################################################################################################################

/* 1. 方法定义 */
//----------------------------------------------------------------------------------------------------------------------
//	Golang 方法总是绑定对象实例，并隐式将实例作为第一实参 (receiver)。
//
//	• 只能为当前包内命名类型定义方法。
//	• 参数 receiver 可任意命名。如方法中未曾使用 ，可省略参数名。
//	• 参数 receiver 类型可以是 T 或 *T。基类型 T 不能是接口或指针。
//	• 不支持方法重载，receiver 只是参数签名的组成部分。
//	• 可用实例 value 或 pointer 调用全部方法，编译器自动转换。
//
//	一个方法就是一个包含了接受者的函数，接受者可以是命名类型或者结构体类型的一个值或者是一个指针。
//
//	所有给定类型的方法属于该类型的方法集。
//----------------------------------------------------------------------------------------------------------------------

/* 1.1.1. 方法定义： */
//----------------------------------------------------------------------------------------------------------------------
//
//	语法示例:
//	func (recevier type) methodName(参数列表)(返回值列表){}
//
//	#参数和返回值可以省略
//
//	================================================================================================================
// 	@todo 1.当接受者不是一个指针时，该方法操作对应接受者的值的副本。
// 			意思就是即使使用了指针调用函数，但是函数的接受者是值类型，所以函数内部操作还是对副本的操作，而不是指针操作；
//			函数内部对 【副本】 的操作，不会像指针那样影响到外部使用。
//
//	@todo 2.当接受者是一个指针时,该方法操作对应接受者的值的指针。
//			即不管是 值调用函数 还是 指针调用函数，函数内部操作是对指针的操作；
//			函数内部对 【指针】 的操作，会像指针那样影响到外部使用。
//
//	3.当接受者不管是值类型还是指针类型，若是在外部调用函数且针对外部成员属性进行修改时，方法内部再使用该成员属性是被外部修改后的值。
//		即对象的成员属性是公用时,不管接受者是值类型还是指针类型,外部对该公用成员属性修改,对象内部方法再使用该公用成员属性时,则是被修改后的值;
//		只有当接受者是值类型时,对象内部对该公用成员属性的修改,外部再使用时,则是不受影响,还是原值（因为是修改的是接受者值的【副本】）;
//		当接收者是指针类型时,对象内部对该公用成员属性的修改,外部再使用时,则是受影响的,即是内部方法修改后的值（因为是修改的接受者值的【指针】）。
//
//	================================================================================================================
//	示例:
//	1.创建method目录,进入到method目录,创建example.go文件并写入内容如下:
//		type User struct {
//			Name string
//			Age int
//		}
//
//		//方法接受者是值类型-外部调用方法
//		//	当方法接受者是值类型时，该方法操作对应 接受者的值 是 接受者值的【副本】操作；
//		//	在方法里操作 接受者值的【副本】，不会影响到外部。
//		func (u User) Echo(){
//			fmt.Println("User.Name-before:",u.Name)
//			u.Name = "方法接受者是值类型，所以这个字符串不会在外部被使用到"
//			fmt.Println("User.Name:",u.Name)
//			fmt.Println("User.Age:",u.Age)
//		}
//
//		//方法接受者是指针类型-外部调用方法
//		//	当方法接受者是指针类型时，该方法操作对应 接受者的值 是 接受者值的【指针】操作；
//		//	在方法里操作 接受者值的【指针】，会影响到外部。
//		func (u *User) Echo2(){
//			fmt.Println("User.Name-before:",u.Name)
//			u.Name = "方法接受者是指针类型，所以这个字符串会在外部显示"
//			fmt.Println("User.Name:",u.Name)
//			fmt.Println("User.Age:",u.Age)
//		}
//
//	--- A.方法接受者是值类型-示例 ---
//	A-1.准备调试代码:
//		//值类型调用方法
//		u1 := method.User{"test1",21}
//		u1.Name = "test111"
//		u1.Echo()
//
//		fmt.Println("")
//		fmt.Println( "u1.Name:",u1.Name )
//		fmt.Println("")
//
//		//指针类型调用方法
//		u2 := &method.User{"test2",22}
//		u2.Name = "test222"
//		u2.Echo()
//
//		fmt.Println("")
//		fmt.Println( "u2.Name:",u2.Name )
//		fmt.Println("")
//
//	A-2.输出结果:
//		User.Name-before: test111
//		User.Name: 方法接受者是值类型，所以这个字符串不会在外部被使用到
//		User.Age: 21
//
//		u1.Name: test111
//
//		User.Name-before: test222
//		User.Name: 方法接受者是值类型，所以这个字符串不会在外部被使用到
//		User.Age: 22
//
//		u2.Name: test222
//	~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
//
//	--- B.方法接受者是指针类型-示例 ---
//	A-1.准备调试代码:
//		//值类型调用方法
//		u1 := method.User{"test1",21}
//		u1.Name = "test111"
//		u1.Echo2()
//
//		fmt.Println("")
//		fmt.Println( "u1.Name:",u1.Name )
//		fmt.Println("")
//
//		//指针类型调用方法
//		u2 := &method.User{"test2",22}
//		u2.Name = "test222"
//		u2.Echo2()
//
//		fmt.Println("")
//		fmt.Println( "u2.Name:",u2.Name )
//		fmt.Println("")
//
//	A-2.输出结果:
//		User.Name-before: test111
//		User.Name: 方法接受者是指针类型，所以这个字符串会在外部显示
//		User.Age: 21
//
//		u1.Name: 方法接受者是指针类型，所以这个字符串会在外部显示
//
//		User.Name-before: test222
//		User.Name: 方法接受者是指针类型，所以这个字符串会在外部显示
//		User.Age: 22
//
//		u2.Name: 方法接受者是指针类型，所以这个字符串会在外部显示
//----------------------------------------------------------------------------------------------------------------------

/* 1.1.2. 普通函数与方法的区别 */
//----------------------------------------------------------------------------------------------------------------------
//1.对于普通函数，形参接收者为【值类型】时，不能将指针类型的数据直接传递；形参接收者为【指针类型】时，不能将值类型的数据直接传递。
//		即普通函数的形参会指定是【值类型】还是【指针类型】传递,
//			如果指定【值类型】则只能传递【值类型】的数据；
//			如果指定【指针类型】则只能传递【指针类型】的数据；
//
//2.对于方法（如struct的方法），接收者为【值类型】时，可以直接用 【调用函数为指针类型的变量】 调用方法；
// 	接收者为【指针类型】时，则可以用【调用函数为值类型的变量】 调用方法。
//		即struct里的方法，不管接收者为【值类型】 或是 【指针类型】，只要是公用的方法，外部都可以被调用；
//		区别在与，接收者为【值类型】时，内部方法修改操作的值，外部都不会收到影响(外部调用函数不管是【值类型】或是【指针类型】的变量)；
//
//----------------------------------------------------------------------------------------------------------------------

/* 2. 匿名字段 */
//----------------------------------------------------------------------------------------------------------------------
//
//	Golang匿名字段 ：可以像字段成员那样访问匿名字段方法，编译器负责查找。
//
//	================================================================================================================
//	示例:
//	1.进入method目录,创建methodAnonymousField.go文件并写入如下内容
//		type MethodAnonymousField struct {
//			Name string
//			Tag int
//		}
//
//		type TestMethodAnonymousField struct {
//			MethodAnonymousField
//		}
//
//		func (thisObj *MethodAnonymousField) Echo(){
//			res := fmt.Sprintf("%p-%v",thisObj,thisObj)
//			fmt.Println(res)
//		}
//
//	2.外部调用示例
//		obj := method.TestMethodAnonymousField{method.MethodAnonymousField{"test",123}}
//		res := fmt.Sprintf("%p-%v",&obj,&obj)
//		fmt.Println(res)
//		fmt.Println("~~~~~~~~~~~~~~~")
//		obj.Echo()
//
//	3.输出结果
//		0xc0000047a0-&{{test 123}}
//		~~~~~~~~~~~~~~~
//		0xc0000047a0-&{test 123}
//----------------------------------------------------------------------------------------------------------------------
//@todo 未完待续