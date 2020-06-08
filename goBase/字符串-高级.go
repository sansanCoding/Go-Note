package goBase

import (
	"Go-Note/util"
	"fmt"
	"reflect"
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

//字符串修改
//	注意:要修改字符串，需要先将其转换成[]rune或[]byte，完成后再转换为string。无论哪种转换，都会重新分配内存，并复制字节数组。
//	总结:
//		1.str:="abc",str[0]获取值可以,但是不能直接赋值!
//			如str[0] = "d",写的时候语法报错.
//			如str[0] = str[0:1],写的时候语法报错.
//		2.使用直接下标str[0],不管是英文还是中文获取到的是ASCII码值!
//		3.使用范围下标str[0:1],获取到才是下标的值!
//		4.若想获取指定下标的值,要么使用[]byte()或[]rune()转成数组,或者将字符串切分成数组,通过数组的指定下标获取其值!
//调试-命令行输入:
//	{"optTag":"StringAdvanced","optParams":{"methodName":"StringsChange"}}
func (thisObj *stringAdvanced) StringsChange(){
	//[]byte:针对ASCII码值修改,如abc123
	//[]rune:针对unicode值修改,如中文abc123 #如果字符串中包含英文中文数字等这些,尽量使用[]rune()处理
	//[]byte和[]rune都是由byte数组组成

	//1.字符串修改示例
	//声明一个字符串
	str := "abc"
	//将字符串转成[]byte()
	strByte := []byte(str)
	//将字符串转成[]rune()
	strRune := []rune(str)
	//若想要strByte[0]的值赋值给strRune[1]上,则需要通过string转回再使用[]rune()处理一次;
	//因为数据类型不一致,所以无法直接将[]byte()的数据赋值给[]rune()上.
	strRune[1] = []rune(string(strByte[0]))[0]
	//输出示例:
	fmt.Println("str:",str)
	fmt.Println("strByte:",string(strByte))
	fmt.Println("strRune:",string(strRune))
	//输出结果:
	//str: abc
	//strByte: abc
	//strRune: aac

	//2.直接使用下标的赋值示例
	//这样赋值,语法报错:
	//示例1:
	//str[0] = "11"
	//示例2:
	//str[0] = str[0:1]

	//3.关于ASCII码和unicode字符串的获取示例
	//ASCII码值得字符串-获取,不能直接用于unicode
	fmt.Println("------ASCII码值的字符串-下标获取以及范围下标获取的特殊性:------")
	fmt.Println("指定下标获取,获取到的是ASCII码:",str[0],reflect.TypeOf(str[0]).String())
	fmt.Println("范围下标获取,获取到的是字符串范围下标的值:",str[0:1],reflect.TypeOf(str[0:1]).String())
	fmt.Println("------------------------------------------------")
	//输出结果:
	//------ASCII码值的字符串-下标获取以及范围下标获取的特殊性:------
	//指定下标获取,获取到的是ASCII码: 97 uint8
	//范围下标获取,获取到的是字符串范围下标的值: a string
	//------------------------------------------------

	//unicode值-获取
	str = "你好"
	fmt.Println("------unicode值的字符串-下标获取以及范围获取的特殊性:------")
	fmt.Println("指定下标获取,获取到的是ASCII码:",[]rune(str)[0],reflect.TypeOf([]rune(str)[0]).String())
	fmt.Println("范围下标获取,获取到的是字符串范围下标的值:",[]rune(str)[0:1],string([]rune(str)[0:1]),reflect.TypeOf([]rune(str)[0:1]).String())
	fmt.Println("------------------------------------------------")
	//输出结果:
	//------unicode值的字符串-下标获取以及范围获取的特殊性:------
	//指定下标获取,获取到的是ASCII码: 20320 int32
	//范围下标获取,获取到的是字符串范围下标的值: [20320] 你 []int32
	//------------------------------------------------

	fmt.Println("TestEnd!")
}
