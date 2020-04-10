package funcTest

import (
	"testing"
	"time"
)

//基准函数-简单-拼接字符串
func BenchmarkSimpleJoinString(b *testing.B) {
	for i:=0;i<b.N;i++ {
		SimpleJoinString([]string{"1","2","3"},",")
	}
}

//基准函数-简单-拼接字符串 优化性能版
func BenchmarkSimpleJoinString2(b *testing.B) {
	for i:=0;i<b.N;i++ {
		SimpleJoinString2([]string{"1","2","3"},",")
	}
}

//基准函数-一个计算第n个斐波那契数的函数
func baseBenchmarkFib(b *testing.B,n int){
	for i:=0;i<b.N;i++ {
		Fib(n)
	}
}
//多个函数性能对比检测
func BenchmarkFib1(b *testing.B) { baseBenchmarkFib(b,1) }
func BenchmarkFib2(b *testing.B) { baseBenchmarkFib(b,2) }
func BenchmarkFib4(b *testing.B) { baseBenchmarkFib(b,4) }
//func BenchmarkFib6(b *testing.B) { baseBenchmarkFib(b,6) }
func BenchmarkFib10(b *testing.B) { baseBenchmarkFib(b,10) }
func BenchmarkFib20(b *testing.B) { baseBenchmarkFib(b,20) }
func BenchmarkFib40(b *testing.B) { baseBenchmarkFib(b,40) }
//func BenchmarkFib60(b *testing.B) { baseBenchmarkFib(b,60) }

//基准函数-没有重置时间测试
func BenchmarkNotResetTimerTest(b *testing.B){
	//time.Sleep(5*time.Second)
	//b.ResetTimer() //重置计时器
	for i:=0;i<b.N;i++ {
		SimpleJoinString2([]string{"1","2","3"},",")
	}
}
//基准函数-有重置时间测试
func BenchmarkHasResetTimerTest(b *testing.B){
	time.Sleep(5*time.Second)
	b.ResetTimer() //重置计时器
	for i:=0;i<b.N;i++ {
		SimpleJoinString2([]string{"1","2","3"},",")
	}
}

//基准函数-非并行测试
func BenchmarkNotParallel(b *testing.B){
	for i:=0;i<b.N;i++ {
		SimpleJoinString2([]string{"1","2","3"},",")
	}
}
//基准函数-并行测试
func BenchmarkTestParallel(b *testing.B){
	// b.SetParallelism(1) // 设置使用的CPU数(通常不需要设置该值;虽然有基于windows系统的测试,效果貌似区别不大,或许测试用例不对吧!)
	b.RunParallel(func(pb *testing.PB){
		for pb.Next() {
			SimpleJoinString2([]string{"1","2","3"},",")
		}
	})
}