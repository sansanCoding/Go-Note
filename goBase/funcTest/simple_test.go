package funcTest

////测试函数:简单-拼接字符串
//func TestSimpleJoinString(t *testing.T){
//	joinRes := SimpleJoinString([]string{"1","2","3"},",")
//	fmt.Println("joinRes:",joinRes)
//}

////测试函数:简单-拼接字符串对比
//func TestSimpleJoinStringCompare(t *testing.T){
//	joinRes1 := SimpleJoinString([]string{"1","2","3"},",")
//	joinRes2 := SimpleJoinString([]string{"1","2","3","4"},",")
//	if reflect.DeepEqual(joinRes1,joinRes2) {
//		fmt.Printf("TestSimpleJoinStringCompareRes: joinRes1:%v eq joinRes2:%v \r\n",joinRes1,joinRes2)
//	}else{
//		t.Errorf("TestSimpleJoinStringCompareRes: joinRes1:%v neq joinRes2:%v",joinRes1,joinRes2)
//	}
//}

////测试函数:简单-拼接字符串对比2(修正错误)
//func TestSimpleJoinStringCompare2(t *testing.T){
//	joinRes1 := SimpleJoinString([]string{"1","2","3"},",")
//	joinRes2 := SimpleJoinString([]string{"1","2","3"},",")
//	if reflect.DeepEqual(joinRes1,joinRes2) {
//		fmt.Printf("TestSimpleJoinStringCompareRes: joinRes1:%v eq joinRes2:%v \r\n",joinRes1,joinRes2)
//	}else{
//		t.Errorf("TestSimpleJoinStringCompareRes: joinRes1:%v neq joinRes2:%v",joinRes1,joinRes2)
//	}
//}

////测试函数-子测试
//func TestSimpleJoinStringCompareRun(t *testing.T){
//	testList := map[string]map[string][]string{
//		"test1":{"compare1":[]string{"1","2","3"},"compare2":[]string{"1","2","3"}},
//		"test2":{"compare1":[]string{"1","2","3"},"compare2":[]string{"11","2","3"}},
//		"test3":{"compare1":[]string{"1","2","3"},"compare2":[]string{"1","2","3"}},
//	}
//	for name,testMap := range testList {
//		t.Run(name,func(t *testing.T){
//			compare1Str := SimpleJoinString(testMap["compare1"],",")
//			compare2Str := SimpleJoinString(testMap["compare2"],",")
//			if reflect.DeepEqual(compare1Str,compare2Str) {
//				fmt.Printf("TestSimpleJoinStringCompareRun: %#v eq %#v \r\n",testMap["compare1"],testMap["compare2"])
//			}else{
//				t.Errorf("TestSimpleJoinStringCompareRun: %#v neq %#v",testMap["compare1"],testMap["compare2"])
//			}
//		})
//	}
//}