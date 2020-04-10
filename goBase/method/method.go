package method

import "fmt"

type User struct {
	Name string
	Age int
}

//方法接受者是值类型-外部调用方法
//	当方法接受者是值类型时，该方法操作对应 接受者的值 是 接受者值的【副本】操作；
//	在方法里操作 接受者值的【副本】，不会影响到外部。
func (u User) Echo(){
	fmt.Println("User.Name-before:",u.Name)
	u.Name = "方法接受者是值类型，所以这个字符串不会在外部被使用到"
	fmt.Println("User.Name:",u.Name)
	fmt.Println("User.Age:",u.Age)
}

//方法接受者是指针类型-外部调用方法
//	当方法接受者是指针类型时，该方法操作对应 接受者的值 是 接受者值的【指针】操作；
//	在方法里操作 接受者值的【指针】，会影响到外部。
func (u *User) Echo2(){
	fmt.Println("User.Name-before:",u.Name)
	u.Name = "方法接受者是指针类型，所以这个字符串会在外部显示"
	fmt.Println("User.Name:",u.Name)
	fmt.Println("User.Age:",u.Age)
}