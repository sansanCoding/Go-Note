package goBase

//defer讲解:
//1.defer是按执行到defer代码才会被执行,若提前停止代码执行,没有到defer代码是不会触发的
//	示例如下:
//	func main() {
//		fmt.Println(1)
//
//		if true {
//			return	//return 或是 panic("test") 等,都是未执行到如下defer的代码,如下defer没有被执行
//		}
//
//		defer func() {
//			fmt.Println(123123)
//		}()
//	}
//	输出结果:
//	1

//2.defer是按栈的方式先进后出,后进先出顺序列执行
//	示例如下:
//	func main() {
//		fmt.Println(1)
//
//		defer func(){
//			fmt.Println(22)	//22其后输出
//		}()
//
//		defer func(){
//			fmt.Println(33)	//33先被输出
//		}()
//
//		if true {
//			panic("test")
//		}
//	}
//	输出结果:
//	1
//	panic: test
//	33
//
//	22