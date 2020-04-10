package goBase

//######################################################################################################################
//GO-函数-单元测试
//	GO-函数-单元测试的说明:
//	  参考文章地址:
//	  http://topgoer.com/%E5%87%BD%E6%95%B0/%E5%8D%95%E5%85%83%E6%B5%8B%E8%AF%95.html
//######################################################################################################################

/* 1. 单元测试 */
//	TDD（Test Driven Development）:测试驱动开发

/* 1.1. go test工具 */
//----------------------------------------------------------------------------------------------------------------------
//	a.Go语言中的测试依赖go test命令
//		go test命令是一个按照一定约定和组织的测试代码的驱动程序。
// 		在包目录内，所有以_test.go为后缀名的源代码文件都是go test测试的一部分，不会被go build编译到最终的可执行文件中。
//
//	b.在*_test.go文件中有三种类型的函数，单元测试函数、基准测试函数和示例函数
//		类型			格式							作用
//		测试函数		函数名前缀为Test				测试程序的一些逻辑行为是否正确
//		基准函数		函数名前缀为Benchmark			测试函数的性能
//		示例函数		函数名前缀为Example			为文档提供示例文档
//
//		go test命令会遍历所有的*_test.go文件中符合上述命名规则的函数，然后生成一个临时的main包用于调用相应的测试函数，
// 		然后构建并运行、报告测试结果，最后清理测试中生成的临时文件。
//
//	c.Golang单元测试对文件名和方法名，参数都有很严格的要求
//		1、文件名必须以xx_test.go命名
//		2、方法必须是Test[^a-z]开头
//		3、方法参数必须 t *testing.T
//		4、使用go test执行单元测试
//
//	d.go test的参数解读
//		go test是go语言自带的测试工具，其中包含的是两类，单元测试和性能测试
//
//		通过go help test可以看到go test的使用说明：
//
//		格式形如： go test [-c] [-i] [build flags] [packages] [flags for test binary]
//
//		参数解读：
//
//		-c : 编译go test成为可执行的二进制文件，但是不运行测试。
//
//		-i : 安装测试包依赖的package，但是不运行测试。
//
//		关于build flags，调用go help build，这些是编译运行过程中需要使用到的参数，一般设置为空
//
//		关于packages，调用go help packages，这些是关于包的管理，一般设置为空
//
//		关于flags for test binary，调用go help testflag，这些是go test过程中经常使用到的参数
//
//		-test.v : 是否输出全部的单元测试用例（不管成功或者失败），默认没有加上，所以只输出失败的单元测试用例。
//
//		-test.run pattern: 只跑哪些单元测试用例
//
//		-test.bench patten: 只跑那些性能测试用例
//
//		-test.benchmem : 是否在性能测试的时候输出内存情况
//
//		-test.benchtime t : 性能测试运行的时间，默认是1s
//
//		-test.cpuprofile cpu.out : 是否输出cpu性能分析文件
//
//		-test.memprofile mem.out : 是否输出内存性能分析文件
//
//		-test.blockprofile block.out : 是否输出内部goroutine阻塞的性能分析文件
//
//		-test.memprofilerate n : 内存性能分析的时候有一个分配了多少的时候才打点记录的问题。这个参数就是设置打点的内存分配间隔，也就是profile中一个sample代表的内存大小。默认是设置为512 * 1024的。如果你将它设置为1，则每分配一个内存块就会在profile中有个打点，那么生成的profile的sample就会非常多。如果你设置为0，那就是不做打点了。
//
//		你可以通过设置memprofilerate=1和GOGC=off来关闭内存回收，并且对每个内存块的分配进行观察。
//
//		-test.blockprofilerate n: 基本同上，控制的是goroutine阻塞时候打点的纳秒数。默认不设置就相当于-test.blockprofilerate=1，每一纳秒都打点记录一下
//
//		-test.parallel n : 性能测试的程序并行cpu数，默认等于GOMAXPROCS。
//
//		-test.timeout t : 如果测试用例运行时间超过t，则抛出panic
//
//		-test.cpu 1,2,4 : 程序运行在哪些CPU上面，使用二进制的1所在位代表，和nginx的nginx_worker_cpu_affinity是一个道理
//
//		-test.short : 将那些运行时间较长的测试用例运行时间缩短
//
//		目录结构：
//
//		test
//		|
//		—— calc.go
//		|
//		—— calc_test.go
//----------------------------------------------------------------------------------------------------------------------

/* 1.2. 测试函数 */
/* 1.2.1. 测试函数的格式 */
//----------------------------------------------------------------------------------------------------------------------
//	a.每个测试函数必须导入testing包，测试函数的基本格式（签名）如下：
//		func TestName(t *testing.T){
//    		// ...
//		}
//
//	b.测试函数的名字必须以Test开头，可选的后缀名必须以大写字母开头，举几个例子:
//		func TestAdd(t *testing.T){ ... }
//		func TestSum(t *testing.T){ ... }
//		func TestLog(t *testing.T){ ... }
//
//	c.其中参数t用于报告测试失败和附加的日志信息。 testing.T的拥有的方法如下：
//		func (c *T) Error(args ...interface{})
//		func (c *T) Errorf(format string, args ...interface{})
//		func (c *T) Fail()
//		func (c *T) FailNow()
//		func (c *T) Failed() bool
//		func (c *T) Fatal(args ...interface{})
//		func (c *T) Fatalf(format string, args ...interface{})
//		func (c *T) Log(args ...interface{})
//		func (c *T) Logf(format string, args ...interface{})
//		func (c *T) Name() string
//		func (t *T) Parallel()
//		func (t *T) Run(name string, f func(t *T)) bool
//		func (c *T) Skip(args ...interface{})
//		func (c *T) SkipNow()
//		func (c *T) Skipf(format string, args ...interface{})
//		func (c *T) Skipped() bool
//----------------------------------------------------------------------------------------------------------------------

/* 1.2.2. 测试函数示例 */
/* 单元测试-简单示例 */
//----------------------------------------------------------------------------------------------------------------------
//	(1).先定义一个funcTest包,创建simple.go文件,在simple.go文件中声明一个被调用的函数SimpleJoinString()
//		package funcTest
//
//		import "strings"
//
//		//单元测试-简单示例
//
//		//简单-拼接字符串
//		func SimpleJoinString(strArr []string,joinTag string) string {
//			str := ""
//			for _,v := range strArr {
//				str += v+joinTag
//			}
//			return strings.Trim(str,joinTag)
//		}
//
//	(2).在funcTest包当前下创建simple_test.go测试文件并在该simple_test.go测试文件中定义一个测试函数
//		package funcTest
//
//		import (
//			"fmt"
//			"testing"
//		)
//
//		//单元测试:简单-拼接字符串
//		func TestSimpleJoinString(t *testing.T){
//			joinRes := SimpleJoinString([]string{"1","2","3"},",")
//			fmt.Println("joinRes:",joinRes)
//		}
//
//	(3).进入到funcTest包目录下,输入命令go test
//		#先进入到funcTest包目录下
//		cd xxx/src/Go-Note/goBase/funcTest
//
//		#输入命令如下:
//		xxx/src/Go-Note/goBase/funcTest>go test
//
//		#直接使用go test命令执行后,输出如下:
//		joinRes: 1,2,3
//		PASS
//		ok      Go-Note/goBase/funcTest 0.326s
//
//
//	上面的简单示例,只有一个测试用例,接下来再新增一个测试用例:
//
//
//	(4).在simple_test.go测试文件中再新增一个测试用例,该新增的测试用例通过对比将会提示错误信息
//		//单元测试:简单-拼接字符串对比
//		func TestSimpleJoinStringCompare(t *testing.T){
//			joinRes1 := SimpleJoinString([]string{"1","2","3"},",")
//			joinRes2 := SimpleJoinString([]string{"1","2","3","4"},",")
//			if reflect.DeepEqual(joinRes1,joinRes2) {
//				fmt.Printf("TestSimpleJoinStringCompareRes: joinRes1:%v eq joinRes2:%v",joinRes1,joinRes2)
//			}else{
//				//使用t.Errorf()输出测试用例失败的信息
//				t.Errorf("TestSimpleJoinStringCompareRes: joinRes1:%v neq joinRes2:%v",joinRes1,joinRes2)
//			}
//		}
//
//	(5).在funcTest包目录下,输入命令go test,输出如下:
//		#输入命令如下:
//		xxx/src/Go-Note/goBase/funcTest>go test
//
//		#输出如下:
//		joinRes: 1,2,3
//		--- FAIL: TestSimpleJoinStringCompare (0.00s)
//		simple_test.go:22: TestSimpleJoinStringCompareRes: joinRes1:1,2,3 neq joinRes2:1,2,3,4
//		FAIL
//		exit status 1
//		FAIL    Go-Note/goBase/funcTest 0.278s
//
//	(6).在funcTest包目录下,输入命令go test -v,查看测试函数名称和运行时间:
//		#输入命令如下:
//		xxx/src/Go-Note/goBase/funcTest>go test -v
//
//		#输出如下:
//		=== RUN   TestSimpleJoinString
//		joinRes: 1,2,3
//		--- PASS: TestSimpleJoinString (0.00s)
//		=== RUN   TestSimpleJoinStringCompare
//		--- FAIL: TestSimpleJoinStringCompare (0.00s)
//		simple_test.go:22: TestSimpleJoinStringCompareRes: joinRes1:1,2,3 neq joinRes2:1,2,3,4
//		FAIL
//		exit status 1
//		FAIL    Go-Note/goBase/funcTest 0.294s
//
//	(7).在funcTest包目录下,输入命令go test -v -run="xxx";-run的值对应一个正则表达式，只有函数名匹配上的测试函数才会被go test命令执行。
//		#输入命令如下:
//		xxx/src/Go-Note/goBase/funcTest>go test -v -run="Compare"
//
//		#输出如下:
//		=== RUN   TestSimpleJoinStringCompare
//		--- FAIL: TestSimpleJoinStringCompare (0.00s)
//		simple_test.go:22: TestSimpleJoinStringCompareRes: joinRes1:1,2,3 neq joinRes2:1,2,3,4
//		FAIL
//		exit status 1
//		FAIL    Go-Note/goBase/funcTest 0.275s
//
//	(8).新增一个修复的simple_test.go测试文件中产生错误的测试用例
//		//单元测试:简单-拼接字符串对比2(修正错误)
//		func TestSimpleJoinStringCompare2(t *testing.T){
//			joinRes1 := SimpleJoinString([]string{"1","2","3"},",")
//			joinRes2 := SimpleJoinString([]string{"1","2","3"},",")
//			if reflect.DeepEqual(joinRes1,joinRes2) {
//				fmt.Printf("TestSimpleJoinStringCompareRes: joinRes1:%v eq joinRes2:%v",joinRes1,joinRes2)
//			}else{
//				t.Errorf("TestSimpleJoinStringCompareRes: joinRes1:%v neq joinRes2:%v",joinRes1,joinRes2)
//			}
//		}
//
//	(9).在funcTest包目录下,输入命令go test -v,查看修复的后结果
//		#输入命令如下:
//		xxx/src/Go-Note/goBase/funcTest>go test -v
//
//		#输出如下:
//		=== RUN   TestSimpleJoinString
//		joinRes: 1,2,3
//		--- PASS: TestSimpleJoinString (0.00s)
//		=== RUN   TestSimpleJoinStringCompare
//		--- FAIL: TestSimpleJoinStringCompare (0.00s)
//		simple_test.go:22: TestSimpleJoinStringCompareRes: joinRes1:1,2,3 neq joinRes2:1,2,3,4
//		=== RUN   TestSimpleJoinStringCompare2
//		TestSimpleJoinStringCompareRes: joinRes1:1,2,3 eq joinRes2:1,2,3
//		--- PASS: TestSimpleJoinStringCompare2 (0.00s)
//		FAIL
//		exit status 1
//		FAIL    Go-Note/goBase/funcTest 0.318s
//
//	(10).将simple_test.go测试文件中屏蔽产生错误的测试用例,测试用例完整通过示例如下:
//		#输入命令如下:
//		xxx/src/Go-Note/goBase/funcTest>go test -v
//
//		#输出如下:
//		=== RUN   TestSimpleJoinString
//		joinRes: 1,2,3
//		--- PASS: TestSimpleJoinString (0.00s)
//		=== RUN   TestSimpleJoinStringCompare2
//		TestSimpleJoinStringCompareRes: joinRes1:1,2,3 eq joinRes2:1,2,3
//		--- PASS: TestSimpleJoinStringCompare2 (0.00s)
//		PASS
//		ok      Go-Note/goBase/funcTest 0.309s
//----------------------------------------------------------------------------------------------------------------------

/* 1.3. 测试组 */
//	测试组可以以一个切片或者map,进行遍历后,传入被测试的函数不同参数进行测试!
//	可使用 %#v 的格式化方式,格式化切片或map输出!
//	示例:
//	t.Errorf("%#v", []string{"1","2"})

/* 1.4. 子测试 */
//	语法示例:
//	func TestXxx(t *testing.T){
//		t.Run("子测试名称", func(t *testing.T) { // 使用t.Run()执行子测试
//			//执行测试函数
//			//xxx()
//			//若有错误,则输出测试失败的信息
//			//t.Errorf("%#v",xxxx)
//		})
//	}
//
//	测试示例:
//	//单元测试-子测试
//	func TestSimpleJoinStringCompareRun(t *testing.T){
//		testList := map[string]map[string][]string{
//			"test1":{"compare1":[]string{"1","2","3"},"compare2":[]string{"1","2","3"}},
//			"test2":{"compare1":[]string{"1","2","3"},"compare2":[]string{"11","2","3"}},
//			"test3":{"compare1":[]string{"1","2","3"},"compare2":[]string{"1","2","3"}},
//		}
//		for name,testMap := range testList {
//			t.Run(name,func(t *testing.T){
//				compare1Str := SimpleJoinString(testMap["compare1"],",")
//				compare2Str := SimpleJoinString(testMap["compare2"],",")
//				if reflect.DeepEqual(compare1Str,compare2Str) {
//					fmt.Printf("TestSimpleJoinStringCompareRun: %#v eq %#v \r\n",testMap["compare1"],testMap["compare2"])
//				}else{
//					t.Errorf("TestSimpleJoinStringCompareRun: %#v neq %#v",testMap["compare1"],testMap["compare2"])
//				}
//			})
//		}
//	}
//
//
//	#输入命令:
//	# 查看全部的测试用例
//	go test -v
//
//	#输出结果:
//	=== RUN   TestSimpleJoinStringCompareRun
//	=== RUN   TestSimpleJoinStringCompareRun/test1
//	TestSimpleJoinStringCompareRun: []string{"1", "2", "3"} eq []string{"1", "2", "3"}
//	=== RUN   TestSimpleJoinStringCompareRun/test2
//	=== RUN   TestSimpleJoinStringCompareRun/test3
//	TestSimpleJoinStringCompareRun: []string{"1", "2", "3"} eq []string{"1", "2", "3"}
//	--- FAIL: TestSimpleJoinStringCompareRun (0.00s)
//	--- PASS: TestSimpleJoinStringCompareRun/test1 (0.00s)
//	--- FAIL: TestSimpleJoinStringCompareRun/test2 (0.00s)
//	simple_test.go:51: TestSimpleJoinStringCompareRun: []string{"1", "2", "3"} neq []string{"11", "2", "3"}
//	--- PASS: TestSimpleJoinStringCompareRun/test3 (0.00s)
//	FAIL
//	exit status 1
//	FAIL    Go-Note/goBase/funcTest 0.284s
//
//
//	#输入命令:
//	# 指定子测试用例
//	go test -v -run="TestSimpleJoinStringCompareRun/test2"
//
//	#输出结果：
//	=== RUN   TestSimpleJoinStringCompareRun
//	=== RUN   TestSimpleJoinStringCompareRun/test2
//	--- FAIL: TestSimpleJoinStringCompareRun (0.00s)
//	--- FAIL: TestSimpleJoinStringCompareRun/test2 (0.00s)
//	simple_test.go:51: TestSimpleJoinStringCompareRun: []string{"1", "2", "3"} neq []string{"11", "2", "3"}
//	FAIL
//	exit status 1
//	FAIL    Go-Note/goBase/funcTest 0.428s

/* 1.5. 测试覆盖率 */
//	测试覆盖率是代码被测试套件覆盖的百分比。通常我们使用的都是语句的覆盖率，也就是在测试中至少被运行一次的代码占总代码的比例。
//
//	a.-cover 查看测试覆盖率:
// 		go test -cover
//
//	b.-coverprofile 将覆盖率相关的记录信息输出到一个文件,存放于【当前文件夹】下(即当前所在的目录下):
//		go test -cover -coverprofile=test.out
//
//	c.go tool cover -html=xxx 使用cover工具来处理生成的记录信息，该命令会打开本地的浏览器窗口生成一个HTML报告。
//		go tool cover -html=test.out	//注意:有可能需要test.out文件提前存在
//
//	==================================================================================================================
//	输入命令:
//		go test -cover
//
//	输出如下:
//		TestSimpleJoinStringCompareRun: []string{"1", "2", "3"} eq []string{"1", "2", "3"}
//		TestSimpleJoinStringCompareRun: []string{"1", "2", "3"} eq []string{"1", "2", "3"}
//		--- FAIL: TestSimpleJoinStringCompareRun (0.00s)
//		--- FAIL: TestSimpleJoinStringCompareRun/test2 (0.00s)
//		simple_test.go:51: TestSimpleJoinStringCompareRun: []string{"1", "2", "3"} neq []string{"11", "2", "3"}
//		FAIL
//		coverage: 100.0% of statements			//测试用例覆盖了100%的代码
//		exit status 1
//		FAIL    Go-Note/goBase/funcTest 0.362s

/* 1.6. 基准测试 */
/* 1.6.1. 基准测试函数格式 */
//----------------------------------------------------------------------------------------------------------------------
//	基准测试就是在一定的工作负载之下检测程序性能的一种方法。
//
//	a.基准测试的基本格式如下：
//		func BenchmarkName(b *testing.B){
//			// ...
//		}
//
//	b.基准测试以Benchmark为前缀，需要一个*testing.B类型的参数b，基准测试必须要执行b.N次，这样的测试才有对照性，
// 	  b.N的值是系统根据实际情况去调整的，从而保证测试的稳定性。 testing.B拥有的方法如下：
//		func (c *B) Error(args ...interface{})
//		func (c *B) Errorf(format string, args ...interface{})
//		func (c *B) Fail()
//		func (c *B) FailNow()
//		func (c *B) Failed() bool
//		func (c *B) Fatal(args ...interface{})
//		func (c *B) Fatalf(format string, args ...interface{})
//		func (c *B) Log(args ...interface{})
//		func (c *B) Logf(format string, args ...interface{})
//		func (c *B) Name() string
//		func (b *B) ReportAllocs()
//		func (b *B) ResetTimer()
//		func (b *B) Run(name string, f func(b *B)) bool
//		func (b *B) RunParallel(body func(*PB))
//		func (b *B) SetBytes(n int64)
//		func (b *B) SetParallelism(p int)
//		func (c *B) Skip(args ...interface{})
//		func (c *B) SkipNow()
//		func (c *B) Skipf(format string, args ...interface{})
//		func (c *B) Skipped() bool
//		func (b *B) StartTimer()
//		func (b *B) StopTimer()
//----------------------------------------------------------------------------------------------------------------------

/* 1.6.2. 基准测试示例 */
//----------------------------------------------------------------------------------------------------------------------
//	(1).先定义一个funcTest包(若该包存在,则不需要创建),创建simpleBenchmark_test.go,在simpleBenchmark_test.go文件中声明一个基准函数BenchmarkSimpleJoinString()
//		//基准函数-简单-拼接字符串
//		func BenchmarkSimpleJoinString(b *testing.B) {
//			for i:=0;i<b.N;i++ {
//				SimpleJoinString([]string{"1","2","3"},",")
//			}
//		}
//
//	(2).进入到funcTest包目录下,输入命令go test -bench=SimpleJoinString
//		#先进入到funcTest包目录下
//		cd xxx/src/Go-Note/goBase/funcTest
//
//		#输入命令如下:
//		xxx/src/Go-Note/goBase/funcTest>go test -bench=SimpleJoinString
//
//		#输出如下:
//		goos: windows
//		goarch: amd64
//		pkg: Go-Note/goBase/funcTest
//		BenchmarkSimpleJoinString-4      3795404               302 ns/op
//		PASS
//		ok      Go-Note/goBase/funcTest 1.764s
//
//		~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
//		针对上面输出结果的字段说明:
//		goos => 目标平台的操作系统:darwin 苹果系统、linux 、windows
//		goarch => 处理器架构之类的:amd64
//		pkg => 包名
//		BenchmarkSimpleJoinString-4 =>
// 			表示针对SimpleJoinString进行的基准测试;
// 			数字4表示GOMAXPROCS的值,这个对于并发基准测试很重要;
//			3795404和302 ns/op表示每次调用SimpleJoinString函数耗时302ns，这个结果是3795404次调用的平均值。
//		~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
//
//	(3).基准测试添加-benchmem参数,可获得内存分配的统计数据
//		#先进入到funcTest包目录下
//		cd xxx/src/Go-Note/goBase/funcTest
//
//		#输入命令如下:
//		xxx/src/Go-Note/goBase/funcTest>go test -bench=SimpleJoinString -benchmem
//
//		#输出如下:
//		goos: windows
//		goarch: amd64
//		pkg: Go-Note/goBase/funcTest
//		BenchmarkSimpleJoinString-4      3628057               316 ns/op              48 B/op          4 allocs/op
//		PASS
//		ok      Go-Note/goBase/funcTest 1.943s
//
//		~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
//		48 B/op => 表示每次操作内存分配了48字节
// 		4 allocs/op => allocs/op则表示每次操作进行了4次内存分配
//		~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
//
//		====================================================================================================
//
//		为了实现性能对比效果,在simple.go文件中添加一个SimpleJoinString2()优化性能版:
//		//测试函数-简单-拼接字符串 优化性能版
//		func SimpleJoinString2(strArr []string,joinTag string) string {
//			var strBuilder strings.Builder
//			//strBuilder.Grow(10000)
//			for _,v := range strArr {
//				strBuilder.WriteString(v)
//				strBuilder.WriteString(joinTag)
//			}
//			return strings.Trim(strBuilder.String(),joinTag)
//		}
//
//
//		#输入命令如下:
//		go test -bench=SimpleJoinString -benchmem
//
//		#输出如下:
//		goos: windows
//		goarch: amd64
//		pkg: Go-Note/goBase/funcTest
//		BenchmarkSimpleJoinString-4      5005947               236 ns/op              48 B/op          4 allocs/op
//		BenchmarkSimpleJoinString2-4     6539142               156 ns/op              40 B/op          2 allocs/op
//		PASS
//		ok      Go-Note/goBase/funcTest 4.957s
//----------------------------------------------------------------------------------------------------------------------

/* 1.6.3. 性能比较函数 */
//----------------------------------------------------------------------------------------------------------------------
//	使用如 go test -bench=. 或 go test -bench=xxx(相同前缀)。
//
//	默认情况下，每个基准测试至少运行1秒。如果在Benchmark函数返回时没有到1秒，则b.N的值会按1,2,5,10,20,50，…增加，并且函数再次运行。
//
//	可以使用-benchtime标志增加最小基准时间，以产生更准确的结果。
//		命令示例:go test -bench=xxx基准函数名 -benchtime=20s
//
//
//	@todo 使用性能比较函数做测试的时候,b.N不可作为输入的大小,只做循环次数大小
//		示例如下:
//		func baseBenchmarkFib(b *testing.B,n int){
//			//这里b.N只作循环次数大小,不作输入参数的大小
//			for i:=0;i<b.N;i++ {
//				xxx(n)
//			}
//		}
//
//	~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
//
//	多个函数性能比较-示例:
//		1.在simple.go中添加一个Fib()测试函数
//			//测试函数-一个计算第n个斐波那契数的函数
//			func Fib(n int) int {
//				if n<2 {
//					return n
//				}
//				return Fib(n-1) + Fib(n-2)
//			}
//
//		2.在simpleBenchmark_test.go中添加多个需要比较性能的函数
//			//基准函数-一个计算第n个斐波那契数的函数
//			func baseBenchmarkFib(b *testing.B,n int){
//				for i:=0;i<b.N;i++ {
//					Fib(n)
//				}
//			}
//			//多个函数性能对比检测
//			func BenchmarkFib1(b *testing.B) { baseBenchmarkFib(b,1) }
//			func BenchmarkFib2(b *testing.B) { baseBenchmarkFib(b,2) }
//			func BenchmarkFib4(b *testing.B) { baseBenchmarkFib(b,4) }
//			func BenchmarkFib6(b *testing.B) { baseBenchmarkFib(b,6) }
//			func BenchmarkFib10(b *testing.B) { baseBenchmarkFib(b,10) }
//			func BenchmarkFib20(b *testing.B) { baseBenchmarkFib(b,20) }
//			func BenchmarkFib40(b *testing.B) { baseBenchmarkFib(b,40) }
//			func BenchmarkFib60(b *testing.B) { baseBenchmarkFib(b,60) }
//
//		3.进入性能测试
//			#进入到funcTest包目录下
//			cd xxx/src/Go-Note/goBase/funcTest
//
//			#输入命令如下:
//			xxx/src/Go-Note/goBase/funcTest>go test -bench=Fib
//
//			#输出结果如下:
//			goos: windows
//			goarch: amd64
//			pkg: Go-Note/goBase/funcTest
//			BenchmarkFib1-4         434111779                2.72 ns/op
//			BenchmarkFib2-4         179502624                6.57 ns/op
//			BenchmarkFib4-4         58211734                21.2 ns/op
//			BenchmarkFib6-4         21108996                56.9 ns/op
//			BenchmarkFib10-4         2891614               410 ns/op
//			BenchmarkFib20-4           23247             51293 ns/op
//			BenchmarkFib40-4               2         763251200 ns/op
//			*** Test killed: ran too long (11m0s).
//			exit status 1
//			FAIL    Go-Note/goBase/funcTest 660.290s
//
//			~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
//			上面的BenchmarkFib60-4 没有被执行出来是因为执行超时了...
//			~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
//
//		4.屏蔽simpleBenchmark_test.go中的BenchmarkFib6-4与BenchmarkFib60-4,再进入性能测试
//			#进入simpleBenchmark_test.go文件中
//			//基准函数-一个计算第n个斐波那契数的函数
//			func baseBenchmarkFib(b *testing.B,n int){
//				for i:=0;i<b.N;i++ {
//					Fib(n)
//				}
//			}
//			//多个函数性能对比检测
//			func BenchmarkFib1(b *testing.B) { baseBenchmarkFib(b,1) }
//			func BenchmarkFib2(b *testing.B) { baseBenchmarkFib(b,2) }
//			func BenchmarkFib4(b *testing.B) { baseBenchmarkFib(b,4) }
//			//func BenchmarkFib6(b *testing.B) { baseBenchmarkFib(b,6) }
//			func BenchmarkFib10(b *testing.B) { baseBenchmarkFib(b,10) }
//			func BenchmarkFib20(b *testing.B) { baseBenchmarkFib(b,20) }
//			func BenchmarkFib40(b *testing.B) { baseBenchmarkFib(b,40) }
//			//func BenchmarkFib60(b *testing.B) { baseBenchmarkFib(b,60) }
//
//			~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
//
//			#进入到funcTest包目录下
//			cd xxx/src/Go-Note/goBase/funcTest
//
//			#输入命令如下:
//			xxx/src/Go-Note/goBase/funcTest>go test -bench=Fib
//
//			#输出结果如下:
//			goos: windows
//			goarch: amd64
//			pkg: Go-Note/goBase/funcTest
//			BenchmarkFib1-4         496873095                2.49 ns/op
//			BenchmarkFib2-4         170287377                6.83 ns/op
//			BenchmarkFib4-4         54618195                22.3 ns/op
//			BenchmarkFib10-4         2791472               436 ns/op
//			BenchmarkFib20-4           22591             53640 ns/op
//			BenchmarkFib40-4               2         816816500 ns/op
//			PASS
//			ok      Go-Note/goBase/funcTest 12.898s
//
//
//		5.使用-benchtime标志增加最小基准时间，以产生更准确的结果
//			#输入命令如下:
//			# 这里指定60s
//			xxx/src/Go-Note/goBase/funcTest>go test -bench=Fib40 -benchtime=60s
//
//			#输出结果如下:
//			goos: windows
//			goarch: amd64
//			pkg: Go-Note/goBase/funcTest
//			BenchmarkFib40-4              86         844786623 ns/op
//			PASS
//			ok      Go-Note/goBase/funcTest 133.898s
//----------------------------------------------------------------------------------------------------------------------

/* 1.6.4. 重置时间 */
//----------------------------------------------------------------------------------------------------------------------
//
//	b.ResetTimer之前的处理不会放到执行时间里，也不会输出到报告中，所以可以在之前做一些不计划作为测试报告的操作。
//
//	~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
//
//	分别测试无重置时间和有重置时间的效果示例如下:
//
//	1.在simpleBenchmark_test.go中写入如下测试函数:
//		//基准函数-没有重置时间测试
//		func BenchmarkNotResetTimerTest(b *testing.B){
//			//time.Sleep(5*time.Second)
//			//b.ResetTimer() //重置计时器
//			for i:=0;i<b.N;i++ {
//				SimpleJoinString2([]string{"1","2","3"},",")
//			}
//		}
//
//		//基准函数-有重置时间测试
//		func BenchmarkHasResetTimerTest(b *testing.B){
//			time.Sleep(5*time.Second)
//			b.ResetTimer() //重置计时器
//			for i:=0;i<b.N;i++ {
//				SimpleJoinString2([]string{"1","2","3"},",")
//			}
//		}
//
//	2.输入无重置时间的命令
//		#进入到funcTest包目录下
//		cd xxx/src/Go-Note/goBase/funcTest
//
//		#输入命令如下:
//		xxx/src/Go-Note/goBase/funcTest>go test -bench=NotResetTimerTest -benchmem
//
//		#输出结果:
//		goos: windows
//		goarch: amd64
//		pkg: Go-Note/goBase/funcTest
//		BenchmarkNotResetTimerTest-4     5926956               213 ns/op              40 B/op          2 allocs/op
//		PASS
//		ok      Go-Note/goBase/funcTest 1.787s
//
//	3.输入有重置时间的命令
//		#进入到funcTest包目录下
//		cd xxx/src/Go-Note/goBase/funcTest
//
//		#输入命令如下:
//		xxx/src/Go-Note/goBase/funcTest>go test -bench=HasResetTimerTest -benchmem
//
//		#输出结果:
//		goos: windows
//		goarch: amd64
//		pkg: Go-Note/goBase/funcTest
//		BenchmarkHasResetTimerTest-4     7469322               169 ns/op              40 B/op          2 allocs/op
//		PASS
//		ok      Go-Note/goBase/funcTest 32.759s
//
//----------------------------------------------------------------------------------------------------------------------

/* 1.6.5. 并行测试 */
//----------------------------------------------------------------------------------------------------------------------
//
//	func (b B) RunParallel(body func(PB))会以并行的方式执行给定的基准测试。
//
//	RunParallel会创建出多个goroutine，并将b.N分配给这些goroutine执行， 其中goroutine数量的默认值为GOMAXPROCS。
// 	用户如果想要增加非CPU受限（non-CPU-bound）基准测试的并行性， 那么可以在RunParallel之前调用SetParallelism。
// 	RunParallel通常会与-cpu标志一同使用。
//
//	~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
//
//	测试用例:
//	1.在simpleBenchmark_test.go中写入如下测试函数:
//		//基准函数-并行测试
//		func BenchmarkTestParallel(b *testing.B){
//			// b.SetParallelism(1) // 设置使用的CPU数(通常不需要设置该值;虽然有基于windows系统的测试,效果貌似区别不大,或许测试用例不对吧!)
//			b.RunParallel(func(pb *testing.PB){
//				for pb.Next() {
//					SimpleJoinString2([]string{"1","2","3"},",")
//				}
//			})
//		}
//
//
//	2.输入命令
// 		#进入到funcTest包目录下
//		cd xxx/src/Go-Note/goBase/funcTest
//
//		#输入命令如下:
//		xxx/src/Go-Note/goBase/funcTest>go test -bench=TestParallel -benchmem
//
//		#输出结果:
//		goos: windows
//		goarch: amd64
//		pkg: Go-Note/goBase/funcTest
//		BenchmarkTestParallel-4         10372520               120 ns/op              40 B/op          2 allocs/op
//		PASS
//		ok      Go-Note/goBase/funcTest 1.650s
//
//	3.添加 -cpu xxx 参数来指定使用的CPU数量
// 		#进入到funcTest包目录下
//		cd xxx/src/Go-Note/goBase/funcTest
//
//		#输入命令如下:
//		xxx/src/Go-Note/goBase/funcTest>go test -bench=TestParallel -benchmem -cpu 2
//
//		#输出结果:
//		goos: windows
//		goarch: amd64
//		pkg: Go-Note/goBase/funcTest
//		BenchmarkTestParallel-2          8355684               171 ns/op              40 B/op          2 allocs/op
//		PASS
//		ok      Go-Note/goBase/funcTest 1.911s
//
//	4.也可以添加 -benchtime=xxx 参数
// 		#进入到funcTest包目录下
//		cd xxx/src/Go-Note/goBase/funcTest
//
//		#输入命令如下:
//		xxx/src/Go-Note/goBase/funcTest>go test -bench=TestParallel -benchmem -cpu 2 -benchtime=60s
//
//		#输出结果:
//		goos: windows
//		goarch: amd64
//		pkg: Go-Note/goBase/funcTest
//		BenchmarkTestParallel-2         680887506              121 ns/op              40 B/op          2 allocs/op
//		PASS
//		ok      Go-Note/goBase/funcTest 93.230s
//
//	======================================================================================================
//
//	并行测试与非并行测试一起基准测试用例:
//	1.在simpleBenchmark_test.go中写入如下测试函数:
//		//基准函数-非并行测试
//		func BenchmarkNotParallel(b *testing.B){
//			for i:=0;i<b.N;i++ {
//				SimpleJoinString2([]string{"1","2","3"},",")
//			}
//		}
//		//基准函数-并行测试
//		func BenchmarkTestParallel(b *testing.B){
//			// b.SetParallelism(1) // 设置使用的CPU数(通常不需要设置该值;虽然有基于windows系统的测试,效果貌似区别不大,或许测试用例不对吧!)
//			b.RunParallel(func(pb *testing.PB){
//				for pb.Next() {
//					SimpleJoinString2([]string{"1","2","3"},",")
//				}
//			})
//		}
//
//	2.输入命令
// 		#进入到funcTest包目录下
//		cd xxx/src/Go-Note/goBase/funcTest
//
//		#输入命令如下:
//		xxx/src/Go-Note/goBase/funcTest>go test -bench=Parallel -benchmem
//
//		#输出结果:
//		goos: windows
//		goarch: amd64
//		pkg: Go-Note/goBase/funcTest
//		BenchmarkNotParallel-4           6218290               192 ns/op              40 B/op          2 allocs/op
//		BenchmarkTestParallel-4         12665455               103 ns/op              40 B/op          2 allocs/op
//		PASS
//		ok      Go-Note/goBase/funcTest 5.277s
//
//	======================================================================================================
//
//----------------------------------------------------------------------------------------------------------------------

/* 1.7. Setup与TearDown */
//----------------------------------------------------------------------------------------------------------------------
//	测试程序有时需要在测试之前进行额外的设置（setup）或在测试之后进行拆卸（teardown）。
//----------------------------------------------------------------------------------------------------------------------

/* 1.7.1. TestMain */
//----------------------------------------------------------------------------------------------------------------------
//
//	通过在*_test.go文件中定义TestMain函数来可以在测试之前进行额外的设置（setup）或在测试之后进行拆卸（teardown）操作。
//
//	如果测试文件包含函数:func TestMain(m *testing.M)那么生成的测试会先调用 TestMain(m)，然后再运行具体测试。
// 	TestMain运行在主goroutine中, 可以在调用 m.Run前后做任何设置（setup）和拆卸（teardown）。
// 	退出测试的时候应该使用m.Run的返回值作为参数调用os.Exit。
//
//	~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
//
//	TestMain语法示例:
//		func TestMain(m *testing.M) {
//			fmt.Println("write setup code here...") // 测试之前的做一些设置
//			// 如果 TestMain 使用了 flags，这里应该加上flag.Parse()
//			retCode := m.Run()                         // 执行测试
//			fmt.Println("write teardown code here...") // 测试之后做一些拆卸工作
//			os.Exit(retCode)                           // 退出测试
//		}
//
//	@todo 需要注意的是：
//  @todo 在调用TestMain时, flag.Parse并没有被调用。
// 	@todo 所以如果TestMain 依赖于command-line标志 (包括 testing 包的标记), 则应该显示的调用flag.Parse。
//
//	~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
//
//	TestMain测试用例:
//	1.创建funcTest_SetupAndTearDown包,并在funcTest_SetupAndTearDown包内创建2个文件:
//		#TestMain的函数所在主文件
//		testMain_test.go
//		#需要测试的函数所在子文件(其实并没有主文件和子文件之分,只是TestMain函数在testMain_test.go文件里就以该文件为主文件而已)
//		testSub_test.go
//
//	2.在testMain_test.go中写入如下:
//		func TestMain(m *testing.M) {
//			fmt.Println("我是测试之前的一些设置...") 	// 测试之前的做一些设置
//			//如果 TestMain 使用了 flags，这里应该加上flag.Parse()
//			retCode := m.Run()                         	// 执行测试
//			fmt.Println("我是测试之后的一些设置...") 	// 测试之后做一些拆卸工作
//			os.Exit(retCode)                           	// 退出测试
//		}
//
//	3.在testSub_test.go中写入如下:
//		func TestSub(t *testing.T){
//			fmt.Println("我是被测试的函数执行过程....")
//		}
//
//	4.输入命令
// 		#进入到funcTest包目录下
//		cd xxx/src/Go-Note/goBase/funcTest_SetupAndTearDown
//
//		#输入命令如下:
//		xxx/src/Go-Note/goBase/funcTest_SetupAndTearDown>go test -v
//
//		#输出结果:
//		我是测试之前的一些设置...
//		=== RUN   TestSub
//		我是被测试的函数执行过程....
//		--- PASS: TestSub (0.00s)
//		PASS
//		我是测试之后的一些设置...
//		ok      Go-Note/goBase/funcTest_SetupAndTearDown        0.310s
//----------------------------------------------------------------------------------------------------------------------

/* 1.7.2. 子测试的Setup与Teardown */
//----------------------------------------------------------------------------------------------------------------------
//
//	父级测试集与子级测试集都需要Setup与Teardown的处理。
//
//	~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
//
//	1.在testSub_test.go文件中写入如下:
//		func testParent(t *testing.T) func(t *testing.T){
//			t.Log("parent-我是父级-基准测试-之前的setup....")
//			return func(t *testing.T){
//				t.Log("parent-我是父级-基准测试-之后的tearDown....")
//			}
//		}
//		func testChild(t *testing.T) func(t *testing.T){
//			t.Log("child-我是子级-基准测试-之前的setup....")
//			return func(t *testing.T){
//				t.Log("child-我是子级-基准测试-之后的tearDown....")
//			}
//		}
//		func TestJoinString(t *testing.T){
//			testDataList := map[string][]string{
//				"test1":{"1","2","3",},
//				"test2":{"4","5","6",},
//				//"test3":{"7","8","9",},
//			}
//
//			//父级函数的Setup与Teardown
//			tearDownTestParent := testParent(t)
//			defer tearDownTestParent(t)
//
//			for name,strArr := range testDataList {
//				t.Run(name,func(t *testing.T){
//					tearDownTestChild := testChild(t)
//					defer tearDownTestChild(t)
//
//					//调试输出
//					fmt.Println(name+"-strArr:",strArr)
//
//					//作对比后的输出
//					//if reflect.DeepEqual(funcTest.SimpleJoinString2(strArr,""),"123") {
//					//	fmt.Println(name+" is yes!")
//					//}else{
//					//	t.Log(name+" is no!")
//					//}
//				})
//			}
//		}
//
//	2.输入命令
// 		#进入到funcTest包目录下
//		cd xxx/src/Go-Note/goBase/funcTest_SetupAndTearDown
//
//		#输入命令如下:
//		xxx/src/Go-Note/goBase/funcTest_SetupAndTearDown>go test -v
//
//		#输出结果:
//		TestMain-测试之前的一些设置...
//		=== RUN   TestJoinString
//		=== RUN   TestJoinString/test1
//		test1-strArr: [1 2 3]
//		=== RUN   TestJoinString/test2
//		test2-strArr: [4 5 6]
//		--- PASS: TestJoinString (0.00s)
//		testSub_test.go:13: parent-我是父级-基准测试-之前的setup....
//		--- PASS: TestJoinString/test1 (0.00s)
//		testSub_test.go:20: child-我是子级-基准测试-之前的setup....
//		testSub_test.go:22: child-我是子级-基准测试-之后的tearDown....
//		--- PASS: TestJoinString/test2 (0.00s)
//		testSub_test.go:20: child-我是子级-基准测试-之前的setup....
//		testSub_test.go:22: child-我是子级-基准测试-之后的tearDown....
//		testSub_test.go:15: parent-我是父级-基准测试-之后的tearDown....
//		PASS
//		TestMain-测试之后的一些设置...
//		ok      Go-Note/goBase/funcTest_SetupAndTearDown        0.346s
//----------------------------------------------------------------------------------------------------------------------

/* 1.8. 示例函数 */
/* 1.8.1. 示例函数的格式 */
//----------------------------------------------------------------------------------------------------------------------
//
//	函数名以Example为前缀。示例函数既没有参数也没有返回值。标准格式如下：
//		func ExampleXxx() {
//			// ...
//		}
//
//----------------------------------------------------------------------------------------------------------------------

/* 1.8.2. 示例函数示例 */
//----------------------------------------------------------------------------------------------------------------------
//
//	1.在funcTest目录下创建simpleExample_test.go文件,写入内容如下:
//		func ExampleEcho() {
//			fmt.Println("ExampleEcho Finish!")
//		}
//
//	2.输入命令
// 		#进入到funcTest包目录下
//		cd xxx/src/Go-Note/goBase/funcTest
//
//		#输入命令如下:
//		xxx/src/Go-Note/goBase/funcTest>go test -run Example
//
//		#输出结果:
//		testing: warning: no tests to run
//		PASS
//		ok      Go-Note/goBase/funcTest 0.371s
//
//	============================================================================================
//		好吧,上面的示例并没有成功,具体有需要可自行调试测试!
//
//		参考文章：
//		http://topgoer.com/%E5%87%BD%E6%95%B0/%E5%8D%95%E5%85%83%E6%B5%8B%E8%AF%95.html#%E7%A4%BA%E4%BE%8B%E5%87%BD%E6%95%B0
//	============================================================================================
//----------------------------------------------------------------------------------------------------------------------
