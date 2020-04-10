package goBase

/* 1. 压力测试 */
/* 1.1.1. Go怎么写测试用例 */
//----------------------------------------------------------------------------------------------------------------------
//	开发程序其中很重要的一点是测试:
// 		如何保证代码的质量;
// 		如何保证每个函数是可运行，且运行结果是正确的;
// 		如何保证写出来的代码性能是好的;
//
// 	单元测试的重点在于发现程序设计或实现的逻辑错误，使问题及早暴露，便于问题的定位解决。
//
// 	性能测试的重点在于发现程序设计上的一些问题，让线上的程序能够在高并发的情况下还能保持稳定。
//
//	Go语言中自带有一个轻量级的测试框架testing和自带的go test命令来实现单元测试和性能测试，
//
// 	testing框架和其他语言中的测试框架类似，可以基于这个框架写针对相应函数的测试用例，也可以基于该框架写相应的压力测试用例。
//
//	另外建议安装gotests插件自动生成测试代码:
//		go get -u -v github.com/cweill/gotests/...
//----------------------------------------------------------------------------------------------------------------------

/* 1.1.2. 如何编写测试用例 */
//----------------------------------------------------------------------------------------------------------------------
//	由于go test命令只能在一个相应的目录下执行所有文件，所以接下来需要新建一个目录【goTest】, 这样所有的代码和测试代码都在这个目录下。
//
//	接下来在该目录下面创建 【goTest.go】和【goTest_test.go】 的go文件。
//
//	1.进入【goTest】目录,写入【goTest.go】内容如下:
//		//除法-浮点数
//		func DivisionByFloat64(num1,num2 float64) (float64,error) {
//			if num2==0 {
//				return 0,errors.New("num2不能为0")
//			}
//
//			return num1 / num2,nil
//		}
//
//	2.进入【goTest】目录,【goTest_test.go】是单元测试文件,但是要记住下面的这些原则:
//		{
//			文件名必须是_test.go结尾的，这样在执行go test的时候才会执行到相应的代码。
//				go test 默认只输出失败的单元测试用例
//				go test -v 输出全部的单元测试用例（不管成功或者失败）
//
//			必须import testing这个包。
//
//			所有的测试用例函数必须是Test开头。
//
//			测试用例会按照源代码中写的顺序依次执行。
//
//			测试函数TestXxx()的参数是testing.T，可以使用该类型来记录错误或者是测试状态。
//
//			测试格式：func TestXxx(t *testing.T),Xxx部分可以为任意的字母数字的组合，
// 			但是首字母不能是小写字母[a-z]，例如Testintdiv是错误的函数名。
//
//			函数中通过调用testing.T的Error, Errorf, FailNow, Fatal, FatalIf方法，说明测试不通过，调用Log方法用来记录测试的信息。
//		}
//
//		#【goTest_test.go】写入内容如下:
//		func Test_DivisionByFloat64_1(t *testing.T) {
//			localMsgPrefix := "Test_DivisionByFloat64_1-"
//
//			res,err := DivisionByFloat64(10,1)
//			if err!=nil {
//				t.Error(localMsgPrefix+"err:",err)	//如果不是预期,若有错误,则提示不通过
//				return
//			}
//
//			t.Log(localMsgPrefix+"res:",res)	//如果正确,记录下有用的信息或期望的信息
//		}
//
//		func Test_DivisionByFloat64_2(t *testing.T) {
//			localMsgPrefix := "Test_DivisionByFloat64_2-"
//
//			t.Log(localMsgPrefix+"res:","this is test!")
//		}
//
//	3.输入命令
//		#先进入到funcTest包目录下
//		cd xxx/src/Go-Note/goBase/goTest
//
//		#输入命令如下:
//		xxx/src/Go-Note/goBase/goTest>go test -v
//
//		#输出结果:
//		=== RUN   Test_DivisionByFloat64_1
//		--- PASS: Test_DivisionByFloat64_1 (0.00s)
//		goTest_test.go:14: Test_DivisionByFloat64_1-res: 10
//		=== RUN   Test_DivisionByFloat64_2
//		--- PASS: Test_DivisionByFloat64_2 (0.00s)
//		goTest_test.go:20: Test_DivisionByFloat64_2-res: this is test!
//		PASS
//		ok      Go-Note/goBase/goTest   0.304s
//----------------------------------------------------------------------------------------------------------------------

/* 1.1.3. 如何编写压力测试 */
//----------------------------------------------------------------------------------------------------------------------
//
//	压力测试用来检测函数(方法）的性能，和编写单元功能测试的方法类似，但需要注意以下几点：
//
//		压力测试用例必须遵循如下格式，其中XXX可以是任意字母数字的组合，但是首字母不能是小写字母
//			func BenchmarkXXX(b *testing.B) { ... }
//
//	go test不会默认执行压力测试的函数，如果要执行压力测试需要带上参数-test.bench，语法:-test.bench="test_name_regex",
// 	例如 go test -test.bench=".*" 或者 go test -bench=".*" 表示测试全部的压力测试函数
//
//	在压力测试用例中,循环体内使用testing.B.N,以使测试可以正常的运行,文件名也必须以_test.go结尾。
//
//	1.进入【goTest】目录,创建并写入【goBenchmark_test.go】内容如下:
//		func Benchmark_DivisionByFloat64(b *testing.B) {
//			for i:=0;i<b.N;i++ {	//使用b.N循环
//				DivisionByFloat64(4,5)
//			}
//		}
//
//		func Benchmark_TestConsumeTime(b *testing.B) {
//			b.StopTimer() //调用该函数停止压力测试的时间计数
//
//			//做一些初始化的工作,例如读取文件数据,数据库连接之类的,
//			//这样这些时间不影响测试函数本身的性能
//
//			b.StartTimer() //重新开始时间
//			for i := 0; i < b.N; i++ {
//				DivisionByFloat64(4, 5)
//			}
//		}
//
//	2.输入命令:
//		#先进入到funcTest包目录下
//		cd xxx/src/Go-Note/goBase/goTest
//
//		#输入命令如下:
//		xxx/src/Go-Note/goBase/goTest>go test -v
//
//		#输出结果:
//		goos: windows
//		goarch: amd64
//		pkg: Go-Note/goBase/goTest
//		Benchmark_DivisionByFloat64-4           1000000000               0.701 ns/op           0 B/op          0 allocs/op
//		Benchmark_TestConsumeTime-4             1000000000               0.381 ns/op           0 B/op          0 allocs/op
//		PASS
//		ok      Go-Note/goBase/goTest   1.626s
//
//	3.输出结果说明
//		上面的结果显示没有执行任何TestXXX的单元测试函数，显示的结果只执行了压力测试函数:
// 		第一条Benchmark_DivisionByFloat64-4执行了1000000000次，
// 			0.701 ns/op:每次执行平均时间是0.701纳秒，
//			0 B/op:每次分配的内存大小，
//			0 allocs/op:每次内存分配操作次数。
//
// 		第二条Benchmark_TestConsumeTime-4执行了1000000000次，
// 			0.381 ns/op:每次的平均执行时间是0.381纳秒，
//			0 B/op:每次分配的内存大小，
//			0 allocs/op:每次内存分配操作次数。
//
//		最后一条显示总共的执行时间。
//----------------------------------------------------------------------------------------------------------------------