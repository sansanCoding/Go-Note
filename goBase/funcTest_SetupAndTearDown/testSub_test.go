package funcTest_SetupAndTearDown

import (
	"fmt"
	"testing"
)

//func TestSub(t *testing.T){
//	fmt.Println("我是被测试的函数执行过程....")
//}

func testParent(t *testing.T) func(t *testing.T){
	t.Log("parent-我是父级-基准测试-之前的setup....")
	return func(t *testing.T){
		t.Log("parent-我是父级-基准测试-之后的tearDown....")
	}
}
func testChild(t *testing.T) func(t *testing.T){
	t.Log("child-我是子级-基准测试-之前的setup....")
	return func(t *testing.T){
		t.Log("child-我是子级-基准测试-之后的tearDown....")
	}
}
func TestJoinString(t *testing.T){
	testDataList := map[string][]string{
		"test1":{"1","2","3",},
		"test2":{"4","5","6",},
		//"test3":{"7","8","9",},
	}

	//父级函数的Setup与Teardown
	tearDownTestParent := testParent(t)
	defer tearDownTestParent(t)

	for name,strArr := range testDataList {
		t.Run(name,func(t *testing.T){
			tearDownTestChild := testChild(t)
			defer tearDownTestChild(t)

			//调试输出
			fmt.Println(name+"-strArr:",strArr)

			//作对比后的输出
			//if reflect.DeepEqual(funcTest.SimpleJoinString2(strArr,""),"123") {
			//	fmt.Println(name+" is yes!")
			//}else{
			//	t.Log(name+" is no!")
			//}
		})
	}
}