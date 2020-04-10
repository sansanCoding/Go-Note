package goTest

import "testing"

func Test_DivisionByFloat64_1(t *testing.T) {
	localMsgPrefix := "Test_DivisionByFloat64_1-"

	res,err := DivisionByFloat64(10,1)
	if err!=nil {
		t.Error(localMsgPrefix+"err:",err)	//如果不是预期,若有错误,则提示不通过
		return
	}

	t.Log(localMsgPrefix+"res:",res)	//如果正确,记录下有用的信息或期望的信息
}

func Test_DivisionByFloat64_2(t *testing.T) {
	localMsgPrefix := "Test_DivisionByFloat64_2-"

	t.Log(localMsgPrefix+"res:","this is test!")
}
