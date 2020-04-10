package funcTest_SetupAndTearDown

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("TestMain-测试之前的一些设置...") 	// 测试之前的做一些设置
	//如果 TestMain 使用了 flags，这里应该加上flag.Parse()
	retCode := m.Run()                         	// 执行测试
	fmt.Println("TestMain-测试之后的一些设置...") 	// 测试之后做一些拆卸工作
	os.Exit(retCode)                           	// 退出测试
}
