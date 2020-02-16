package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

//自定义map,json解析时,按map录入的key作为json字符串解析

type SeqMap struct {
	keys   []string
	seqMap map[string]interface{}
	err    string
}

func NewSeqMap() SeqMap {
	var seq SeqMap
	seq.keys = make([]string, 0)
	seq.seqMap = make(map[string]interface{})
	return seq
}

// 通过map主键唯一的特性过滤重复元素
func (seq *SeqMap) removeRepByMap() {
	result := []string{}                    // 存放结果
	tempMap := make(map[string]interface{}) // 存放不重复主键
	for i := len(seq.keys) - 1; i >= 0; i-- {
		e := seq.keys[i]
		l := len(tempMap)
		tempMap[e] = 0
		if len(tempMap) != l { // 加入map后，map长度变化，则元素不重复
			result = append(result, e)
		}
	}
	seq.keys = make([]string, 0) // 存放结果
	for i := len(result) - 1; i >= 0; i-- {
		seq.keys = append(seq.keys, result[i])
	}
}

//添加
func (seq *SeqMap) Put(key string, value interface{}) {
	seq.keys = append(seq.keys, key)
	seq.seqMap[key] = value
}

//获取
func (seq *SeqMap) Get(key string) interface{} {
	return seq.seqMap[key]
}

//是否存在
func (seq *SeqMap) Has(key string) bool {
	_, ok := seq.seqMap[key]
	return ok
}

//获取map值
func (seq *SeqMap) GetMap(key string) (SeqMap, error) {
	reflect.TypeOf(seq.seqMap[key])
	if reflect.TypeOf(seq.seqMap[key]).Name() == "SeqMap" {
		return seq.seqMap[key].(SeqMap), nil
	} else {
		return NewSeqMap(), errors.New("the value is not a SeqMap!")
	}
}

//移除某个元素
func (seq *SeqMap) Remove(key string) {
	for i, k := range seq.keys {
		if k == key {
			seq.keys = append(seq.keys[:i], seq.keys[i+1:]...)
			break
		}
	}
	delete(seq.seqMap, key)
}

//获取所有key
func (seq *SeqMap) Keys() []string {
	seq.removeRepByMap()
	return seq.keys
}

//获取所有值
func (seq *SeqMap) Values() map[string]interface{} {
	return seq.seqMap
}

//解析为json字符串输出
//注:无法处理富文本内容
func (seq *SeqMap) JsonSeq() string {
	seq.removeRepByMap()
	resStr := "{"
	for _, key := range seq.keys {
		var sumMap SeqMap
		if reflect.TypeOf(seq.seqMap[key]).Name() == "SeqMap" {
			sumMap = seq.seqMap[key].(SeqMap)
			str := (&sumMap).JsonSeq()
			resStr += fmt.Sprintf("\"%v\":%v,", key, str)
		} else {
			jsonStr, _ := json.Marshal(seq.seqMap[key])
			resStr += fmt.Sprintf("\"%v\":%v,", key, string(jsonStr))
		}
	}
	resStr = resStr[0 : len(resStr)-1]
	resStr += "}"
	return strings.Replace(strings.Replace(strings.Replace(resStr, "\"{", "{", -1), "}\"", "}", -1), "\\", "", -1)
}

//生成JSON串，返回Unicode编码内容
////主要是解决富文本内容,示例如下:
//	seqMap := util.NewSeqMap()
//	seqMap.Put("test5","<br />https://www.baidu.com/<br />")
//	seqMap.Put("test2","1")
//	res := seqMap.JsonSeqUnicode()
//	fmt.Println(res)
////输出结果:
//	{"test5":"<br />https://www.baidu.com/<br />","test2":"1"}
func (seq *SeqMap) JsonSeqUnicode() string {
	seq.removeRepByMap()
	resStr := "{"
	for _, key := range seq.keys {
		var sumMap SeqMap
		if reflect.TypeOf(seq.seqMap[key]).Name() == "SeqMap" {
			sumMap = seq.seqMap[key].(SeqMap)
			str := (&sumMap).JsonSeq()
			resStr += fmt.Sprintf("\"%v\":%v,", key, str)
		} else {
			jsonStr, _ := json.Marshal(seq.seqMap[key])
			resStr += fmt.Sprintf("\"%v\":%v,", key, string(jsonStr))
		}
	}
	resStr = resStr[0 : len(resStr)-1]
	resStr += "}"
	resStr = strings.ReplaceAll(resStr, "\\\\", "\\")
	resStr = strings.ReplaceAll(resStr, "\\u003c", "<")
	resStr = strings.ReplaceAll(resStr, "\\u003e", ">")
	resStr = strings.ReplaceAll(resStr, "\\u0026", "&")
	resStr = strconv.QuoteToASCII(resStr)
	resStr = strings.ReplaceAll(resStr, "\\\\", "\\")
	resStr = strings.ReplaceAll(resStr, "\\\\\"", "\"")
	resStr = strings.ReplaceAll(resStr, "\\\"", "\"")
	resStr = strings.Replace(resStr, "\"{", "{", -1)
	resStr = strings.Replace(resStr, "}\"", "}", -1)
	return resStr
}
