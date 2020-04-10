package method

import "fmt"

type MethodAnonymousField struct {
	Name string
	Tag int
}

type TestMethodAnonymousField struct {
	MethodAnonymousField
}

func (thisObj *MethodAnonymousField) Echo(){
	res := fmt.Sprintf("%p-%v",thisObj,thisObj)
	fmt.Println(res)
}