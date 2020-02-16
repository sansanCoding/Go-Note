package util

import (
	"strings"
	"sync"
)

//池操作,如对象池,字符串池

type pool struct{

}

//单例对象
var Pool *pool

func init() {
	Pool = NewPool()
}

func NewPool() *pool {
	return &pool{

	}
}

//----------------------------------- 池声明集合区 start -----------------------------------
//1.bytes.Buffer对象池声明
//var bufPool = &sync.Pool{
//	New: func() interface{} {
//		return new(bytes.Buffer)
//	},
//}
//2.bytes.Buffer对象池操作
//拼接2个字符串,效果如s1("a","b") 输出如a:b
//func s1(a, b1 string) string {
//	b := bufPool.Get().(*bytes.Buffer)
//	b.Reset()
//	b.WriteString(a)
//	b.WriteString(":")
//	b.WriteString(b1)
//	s := b.String()
//	bufPool.Put(b)
//	return s
//}

//strings.Builder字符串拼接处理-对象池
var stringBuilderPool = &sync.Pool{
	New: func() interface{} {
		return new(strings.Builder)
	},
}
//----------------------------------- 池声明集合区 end -----------------------------------

//----------------------------------- 池操作集合区 start -----------------------------------

//拼接字符串,效果如StringJoin(",","a","b") 输出如 a,b
func (thisObj *pool) StringJoin(sep string,strMore ...string) string {
	strMoreLen := len(strMore)

	//获取该对象
	b := stringBuilderPool.Get().(*strings.Builder)
	//每次都将该对象的数据重置清空
	b.Reset()

	//循环将每个字符串压入到对象中
	for i:=0;i<strMoreLen;i++ {
		b.WriteString(strMore[i])
		//若是最后一个则不需要拼接符号
		if i<strMoreLen-1 {
			b.WriteString(sep)
		}
	}
	s := b.String()

	//存储到对象池中
	stringBuilderPool.Put(b)
	return s
}

//----------------------------------- 池操作集合区 end -----------------------------------