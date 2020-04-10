package funcTest

import "strings"

//单元测试-简单示例

//测试函数-简单-拼接字符串
func SimpleJoinString(strArr []string,joinTag string) string {
	str := ""
	for _,v := range strArr {
		str += v+joinTag
	}
	return strings.Trim(str,joinTag)
}

//测试函数-简单-拼接字符串 优化性能版
func SimpleJoinString2(strArr []string,joinTag string) string {
	var strBuilder strings.Builder
	//strBuilder.Grow(10000)
	for _,v := range strArr {
		strBuilder.WriteString(v)
		strBuilder.WriteString(joinTag)
	}
	return strings.Trim(strBuilder.String(),joinTag)
}

//测试函数-一个计算第n个斐波那契数的函数
func Fib(n int) int {
	if n<2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}