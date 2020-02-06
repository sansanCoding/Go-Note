package util

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

//interface转整型
func InterfaceToInt(data interface{}) (int,error) {
	dataStr,dataStrErr := JsonValToStr(data)
	if dataStrErr!=nil {
		return 0,dataStrErr
	}
	dataInt,dataIntErr := strconv.Atoi(dataStr)
	if dataIntErr!=nil {
		return 0,dataIntErr
	}
	return dataInt,nil
}

//interface转字符串
func InterfaceToStr(data interface{}) (string,error) {
	return JsonValToStr(data)
}

//Json值转字符串
func JsonValToStr(s interface{}) (string,error) {
	switch s.(type){
	case string:
		return s.(string),nil
	case int:
		return strconv.Itoa(s.(int)),nil
	case int8:
		s = int64(s.(int8))
		return strconv.FormatInt(s.(int64),10),nil
	case int16:
		s = int64(s.(int16))
		return strconv.FormatInt(s.(int64),10),nil
	case int32:
		s = int64(s.(int32))
		return strconv.FormatInt(s.(int64),10),nil
	case int64:
		return strconv.FormatInt(s.(int64),10),nil
	case float32:
		return strconv.FormatFloat(float64(s.(float32)), 'f', -1, 32),nil
	case float64:
		// 'b' (-ddddp±ddd，二进制指数)
		// 'e' (-d.dddde±dd，十进制指数)
		// 'E' (-d.ddddE±dd，十进制指数)
		// 'f' (-ddd.dddd，没有指数)
		// 'g' ('e':大指数，'f':其它情况)
		// 'G' ('E':大指数，'f':其它情况)
		return strconv.FormatFloat(s.(float64), 'f', -1, 64),nil
	}

	return "",errors.New( "JsonValToStr: "+fmt.Sprint(s)+" 数据类型 "+fmt.Sprint(reflect.TypeOf(s))+" 无法强转成字符串类型")
}
