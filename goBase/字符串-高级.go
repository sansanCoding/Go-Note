package goBase

import (
	"Go-Note/util"
	"fmt"
	"strings"
	"sync"
	"time"
)

//字符串-高级

type stringAdvanced struct {

}

var StringAdvanced *stringAdvanced

func init(){
	StringAdvanced = NewStringAdvanced()
}

func NewStringAdvanced()*stringAdvanced{
	return &stringAdvanced{

	}
}

//执行入口
func (thisObj *stringAdvanced) Do(params map[string]interface{}){
	//传参必须-方法名
	methodName := params["methodName"].(string)

	//CallMethodReflect调试:
	res,resOk := util.Helper.CallMethodReflect(thisObj,methodName,[]interface{}{})

	//输出结果:
	fmt.Println(res,resOk)
	for k,v := range res {
		fmt.Println("CallMethodReflectRes:",k,v)
	}
}

//高效性能的字符串拼接:
//主要借助 strings.Builder 和 预分配内存
//命令行-输入:{"optTag":"StringAdvanced","optParams":{"methodName":"StringsJoin"}}
func (thisObj *stringAdvanced) StringsJoin(){
	{
		var strBuilder strings.Builder
		//这里预先分配的内存是字节数,如abc就是3个字节
		strBuilder.Grow(50000)
		for i:=0;i<50000;i++ {
			strBuilder.WriteString("1")
		}
		fmt.Println( "strBuilderRes:",len(strBuilder.String()),strBuilder.String() )

		//输出结果:
		//strBuilderRes: 50000 11111111111111111111.........
	}

	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")

	//但是strings.Builder是非线程安全的,即多个协程处理时,若没有锁机制,会造成数据错乱的
	{
		var strBuilder strings.Builder
		//这里预先分配的内存是字节数,如abc就是3个字节
		strBuilder.Grow(100)
		for i:=0;i<100;i++ {
			go strBuilder.WriteString("1")
		}
		fmt.Println( "strBuilderRes:",len(strBuilder.String()),strBuilder.String() )

		//输出结果:
		//strBuilderRes: 97 1111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111
	}

	{
		//线程安全-依靠锁机制实现
		var lock sync.Mutex
		dataChan := make(chan string)
		var strBuilder strings.Builder
		for i:=0;i<10;i++ {
			go func(){
				dataChan<-"1"
			}()
		}
		//与通道无关,使用排斥锁实现同步效果!
		//	按理说通道也是一个一个来的,以为通过通道可以实现同步效果,估计strings.Builder来不及写入造成的丢失吧!
loop:
		for {
			select {
				case res := <- dataChan :
					{
						//fmt.Println("res:",res)
						lock.Lock()
						strBuilder.WriteString(res)
						lock.Unlock()
					}
				default:
					{
						fmt.Println("forChannlerFinish!")
						break loop
					}
			}
		}

		fmt.Println( "strBuilderStr:",len(strBuilder.String()),strBuilder.String() )

		time.Sleep(5*time.Second)

		fmt.Println("finish!!!")
	}
}

//字符串分散
//命令行-输入:{"optTag":"StringAdvanced","optParams":{"methodName":"StringsDispersed"}}
func (thisObj *stringAdvanced) StringsDispersed(){
	str := "123"
	//只有将字符串打散追加到[]byte组里时,字符串...才有效果
	var bytes []byte
	bytes = append(bytes,str...)
	fmt.Println(string(bytes))
}
