package goTest

import "errors"

//除法-浮点数
func DivisionByFloat64(num1,num2 float64) (float64,error) {
	if num2==0 {
		return 0,errors.New("num2不能为0")
	}

	return num1 / num2,nil
}