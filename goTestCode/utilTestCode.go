package goTestCode

import (
	"Go-Note/goBase/method"
	"Go-Note/util"
	"errors"
	"fmt"
	"net/url"
	"reflect"
	"strings"
)

//util包的测试代码

type utilTestCode struct {

}

var UtilTestCode *utilTestCode

func init(){
	UtilTestCode = NewUtilTestCode()
}

func NewUtilTestCode() *utilTestCode {
	return &utilTestCode{

	}
}

//执行入口
func (thisObj *utilTestCode) Do(params map[string]interface{}){
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

//curl调试
//命令行-输入:{"optTag":"Util","optParams":{"methodName":"Curl"}}
func (thisObj *utilTestCode) Curl(){
	requestApi := "https://www.baidu.com/"

	////POST+Header
	////按json数据请求
	////res,resErr,resErrParams := util.Curl.RequestPostJsonHeader(requestApi,"{\"test\":\"test\"}",10*time.Second,map[string]string{})
	////按post数据请求
	//res,resErr,resErrParams := util.Curl.RequestPostHeader(requestApi,map[string]string{
	//	"test":"test",
	//},10*time.Second,map[string]string{})
	//fmt.Println("res:",res)
	//fmt.Println("resErr:",resErr)
	//fmt.Println("resErrParams:",resErrParams)

	//POST+Header
	requestParams := map[string]interface{}{
		"apiParams":map[string]string{
			"UserName":"Test",
			"PassWord":"123456",
		},
		"apiResponseDataType":"map",
	}
	requestHeader := map[string]string{}
	responseData := make(map[string]interface{})
	requestErr,requestErrParams := util.Curl.Request("post",requestApi,requestParams,requestHeader,&responseData)

	//调试输出:
	fmt.Println("responseData:",responseData)
	fmt.Println("requestErr:",requestErr)
	fmt.Println("requestErrParams:",requestErrParams)

}

//seqMap调试
//命令行-输入:{"optTag":"Util","optParams":{"methodName":"SeqMap"}}
func (thisObj *utilTestCode) SeqMap(){
	seqMap := util.NewSeqMap()
	seqMap.Put("test5","<br />https://www.baidu.com/<br />")
	seqMap.Put("test2","1")
	resKeys := seqMap.Keys()
	fmt.Println("resKeys:",resKeys)
	resStr 	:= seqMap.JsonSeqUnicode()
	fmt.Println("resStr:",resStr)
}

//Base64UrlSafeEncode调试
//命令行-输入:{"optTag":"Util","optParams":{"methodName":"Base64UrlSafeEncode"}}
func (thisObj *utilTestCode) Base64UrlSafeEncode(){
	qrcodeContent := "http://www.baidu.com/test=1+1"
	urlPath := "http://192.168.0.1:58909?qrcode="+url.QueryEscape(util.Base64Encoding.RawURLEncoding([]byte(qrcodeContent)))
	fmt.Println("urlPath:",urlPath)
}

//Regex调试
//命令行-输入:{"optTag":"Util","optParams":{"methodName":"Regex"}}
func (thisObj *utilTestCode) Regex(){
	//"Testsansan":返回true
	//"1Testsansan":返回false
	regexpCheckRes := util.Regex.CheckUserName("1Testsansan")
	fmt.Println("regexpCheckRes:",regexpCheckRes)
}

//用户名称尾部隐藏-调试
//命令行-输入:{"optTag":"Util","optParams":{"methodName":"UserNameTailHidden"}}
func (thisObj *utilTestCode) UserNameTailHidden(){
	res := util.Helper.UserNameTailHidden("张")
	fmt.Println("res:",res)
}

//IP与整型转换-调试
//命令行-输入:{"optTag":"Util","optParams":{"methodName":"IPAndIntegerConversion"}}
func (thisObj *utilTestCode) IPAndIntegerConversion(){
	ipStr 		:= "127.0.0.1"
	ip2Long 	:= util.Helper.IP2long(ipStr)
	long2Str 	:= util.Helper.Long2IP(ip2Long)
	fmt.Println("ipStr:",ipStr,"ip2Long:",ip2Long,"long2Str:",long2Str)

	ipStr 		= "10.10.10.101"
	ip2Long 	= util.Helper.IP2long(ipStr)
	long2Str 	= util.Helper.Long2IP(ip2Long)
	fmt.Println("ipStr:",ipStr,"ip2Long:",ip2Long,"long2Str:",long2Str)
}

//运行日志
//命令行-输入:{"optTag":"Util","optParams":{"methodName":"RunLog"}}
func (thisObj *utilTestCode) RunLog(){
	var newLogFileParams util.DetailLogNewFileParams
	newLogFileParams.DiyLogInfo = "TestDiyLogInfo"
	newLogFileParams.FilePath = "Test.log"
	newLogFileParams.DirPath = "Test1/Test2"
	newLogFileParams.LocalMsgPrefix = "TestMsgPrefix"
	util.LogFile.RunLog("Test",map[string]interface{}{
		"logParams":map[string]interface{}{
			"1.data":"testtesttest",
		},
		"howLongToClean":1,//1小時清除一次日志内容
	},newLogFileParams)
}

//AES加密
//命令行-输入:{"optTag":"Util","optParams":{"methodName":"AESEncrypt"}}
func (thisObj *utilTestCode) AESEncrypt(){
	//标记是从 (APP) 还是 (PC或其他终端) 过来的
	requestTypeVal := "pv10"

	//AES加密key
	aesEncryptKey := ""
	//截取最后四个字符获取请求类型,【注：APP和PC端解密key不一样】
	//requestType  := string(encryptStr[encryptLen-4:encryptLen-3])
	requestType  := requestTypeVal[:1]
	switch requestType {
	//APP请求
	case "p":
		aesEncryptKey = "~!@#$%^&*(VCG789"
		break
	//PC请求或其他终端请求
	case "h":
		aesEncryptKey = "QPO(%#&@1_!-@=+~"
		break
	}

	//矢量字符串,对秘钥进行倒转并转为小写
	aesEncryptIv := strings.ToLower(util.Reverse(aesEncryptKey))

	//原始数据
	originData := "UserName=Test&Age=30"

	//AES加密处理
	encryptByte,err := util.AesCrypt.AES128Pkcs7PaddingIvCBCEncrypt([]byte(originData),[]byte(aesEncryptKey), []byte(aesEncryptIv))

	//输出调试:
	if err!=nil {
		fmt.Println("encryptByteErr:",err)
	}else{
		//程序调试输出的加密数据,输出:Q52GSHZIhbBKXn0Duhrp7EuYVu+HDPHetxMFE8Vf3eo=
		fmt.Println("encryptData:",string(encryptByte))
		//非调试,按流程走的时候,最后要加上requestTypeVal值,以确认是从哪个客户端来的
		//fmt.Println("encryptData:",string(encryptByte)+requestTypeVal)
	}

}

//AES解密
//命令行-输入:{"optTag":"Util","optParams":{"methodName":"AESDecrypt"}}
func (thisObj *utilTestCode) AESDecrypt(){
	//标记是从 (APP) 还是 (PC或其他终端) 过来的
	requestTypeVal := "pv10"

	//请求过来的加密数据
	encryptData := "Q52GSHZIhbBKXn0Duhrp7EuYVu+HDPHetxMFE8Vf3eo="+requestTypeVal

	//加密数据 【注：客户端可能携带空格、换行符\r\n过来】
	encryptStr := util.Trim(encryptData, " ")
	encryptStr  = strings.Replace(encryptStr, "\n", "", -1)
	encryptStr  = strings.Replace(encryptStr, "\r", "", -1)
	encryptLen := len(encryptStr)
	if encryptLen <= 4 {
		fmt.Println("参数错误!")
		return
	}

	//AES加密key
	aesEncryptKey := ""
	//截取最后四个字符获取请求类型,【注：APP和PC端解密key不一样】
	requestType  := string(encryptStr[encryptLen-4:encryptLen-3])
	switch requestType {
	//APP请求
	case "p":
		aesEncryptKey = "~!@#$%^&*(VCG789"
		break
	//PC请求或其他终端请求
	case "h":
		aesEncryptKey = "QPO(%#&@1_!-@=+~"
		break
	}

	//矢量字符串,对秘钥进行倒转并转为小写
	aesEncryptIv := strings.ToLower(util.Reverse(aesEncryptKey))

	//截取到最后四个字符之前获取请求数据
	aesEncryptVal := string(encryptStr[0:encryptLen-4])

	//AES解密处理
	input, err := util.AesCrypt.AES128Pkcs7PaddingIvCBCDecrypt(aesEncryptVal, []byte(aesEncryptKey), []byte(aesEncryptIv))

	//输出调试:
	fmt.Println("input:",string(input))	//输出结果:UserName=Test&Age=30
	fmt.Println("err:",err)
}

//函数defer+闭包
//命令行-输入:{"optTag":"Util","optParams":{"methodName":"FuncDeferClosure"}}
func (thisObj *utilTestCode) FuncDeferClosure(){
	fmt.Println("~~~~~~~~~~~ testFunc(defer+闭包调用的参数,被调用的参数根据外部环境而改变,而非当时传参的值) ~~~~~~~~~~~")

	testFunc := func(){
		for i:=0;i<3;i++ {
			defer func(){
				fmt.Println("i:",i)
			}()
		}
	}
	testFunc()
	//输出结果:
	//i: 3
	//i: 3
	//i: 3

	fmt.Println("~~~~~~~~~~~ testFunc1(defer+闭包调用的参数,被调用的参数存储当时传参的值-1) ~~~~~~~~~~~")

	testFunc1 := func(){
		for i:=0;i<3;i++ {
			//一种是临时声明一个变量,刷新匿名函数里的外部变量
			i:=i
			defer func(){
				fmt.Println("i:",i)
			}()
		}
	}
	testFunc1()
	//输出结果:
	//i: 2
	//i: 1
	//i: 0

	fmt.Println("~~~~~~~~~~~ testFunc2(defer+闭包调用的参数,被调用的参数存储当时传参的值-2) ~~~~~~~~~~~")

	testFunc2 := func(){
		for i:=0;i<3;i++ {
			//一种是相当于函数传参,推荐这种,逻辑清晰,容易理解,方便维护
			defer fmt.Println("i:",i)
			defer func(data int){
				fmt.Println("data:",data)
			}(i)
		}
	}
	testFunc2()
	//输出结果:
	//data: 2
	//i: 2
	//data: 1
	//i: 1
	//data: 0
	//i: 0
}
//函数defer陷阱
//命令行-输入:{"optTag":"Util","optParams":{"methodName":"FuncDeferTrap"}}
func (thisObj *utilTestCode) FuncDeferTrap(){
	fmt.Println("~~~~~~~~~~~~~~~~~~~ defer陷阱-1: defer 与 closure 的传参-复制和传参-闭包引用,造成的结果会不同 ~~~~~~~~~~~~~~~~~~~")
	{
		testFunc := func() (err error) {
			//这个defer执行时,fmt.Println函数将当时的err复制保存一份,相当于当时传参值是什么,到时候会输出什么
			defer fmt.Println("err:",err)	//err被复制
			//这个defer执行时,也是和上面的一样,将当时的err复制保存一份,相当于当时传参值是什么,到时候会输出什么
			defer func(tempErr error){
				fmt.Println("funcErr:",tempErr)
			}(err)	//err被复制
			//这个defer执行时,由于err是闭包引用的值,等被执行时,err就是最新修改的值
			defer func(){
				fmt.Println("deferErr:",err)	//err 闭包引用
			}()

			err = errors.New("test_error")
			return
		}
		testFunc()
		//输出结果:
		//deferErr: test_error
		//funcErr: <nil>
		//err: <nil>
	}

	fmt.Println("~~~~~~~~~~~~~~~~~~~ defer陷阱-2: defer 与 return 的具名返回值处理流程 以及 defer在return之后的处理流程 ~~~~~~~~~~~~~~~~~~~")
	{
		testFunc := func() (actionResult int) {
			//1.先将具名返回值actionResult赋值为1
			actionResult = 1

			//2.actionResult是闭包引用的值,所以在return之前如果还有最新赋值操作,则defer的actionResult输出最新赋值结果
			defer func(){
				fmt.Println("actionResult:",actionResult)
			}()

			//3.这里的return返回的表达式,如果没有具名返回值,则是直接返回;
			//  一旦有具名返回值,会在return之前和defer之前,先将该表达式处理的结果赋值给具名返回值;
			//	等defer再用到具名返回值是,具名返回值则是最新赋值结果.
			//return 22	//输出结果 actionResult: 22
			//return 10*22	//输出结果 actionResult: 220
			//return actionResult+22 //输出结果 actionResult: 23

			//return actionResult+=22)					//这一种还有下面一种都是语法编译不过去的
			//return actionResult = actionResult +22
			//return func() int { return actionResult+22 }()	//当然这种匿名函数写法也是actionResult闭包引用+22返回的结果 //输出结果 actionResult: 23

			//由此可以看出,defer是被程序运行时添加到一种类似排队效果的处理,而不是在语法定义层面就会被记录,如func(),可在函数声明前调用!!!
			//	简单来说,defer是没遇到return之前,会被记录下来,到时候统一在return之前挨个执行,执行顺序就是先进后出的方式!
			defer func(){
				fmt.Println("~~~我这个defer是在return之后,所以是不会被执行的~~~")
			}()
			return	//这个return是因为上面加了defer,不写return会形成语法错误,编译不过去
		}
		testFunc()
	}
}
//函数defer nil函数
//命令行-输入:{"optTag":"Util","optParams":{"methodName":"FuncDeferNil"}}
func (thisObj *utilTestCode) FuncDeferNil(){
	//改进版
	{
		func(){
			var testFunc func() = nil
			if testFunc!=nil {
				defer testFunc()
			}
			fmt.Println(testFunc,reflect.TypeOf(testFunc).String())	//输出结果:<nil> func()
			fmt.Println("this is test")	//输出结果:this is test
		}()
	}

	fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")

	//异常版
	{
		func(){
			//不过从开发角度讲,应该不会这么干,即使真要真没干,估计要提前判断下
			var testFunc func() = nil	//testFunc声明数据类型为func(),值为nil
			defer testFunc()
			fmt.Println("this is test")	//先输出在抛出异常

			//输出结果:
			//this is test
			//紧接着抛出异常:
			//runtime error: invalid memory address or nil pointer dereference
		}()
	}
}
//函数defer 不检查错误
//命令行-输入:{"optTag":"Util","optParams":{"methodName":"FuncDeferNotCheckError"}}
func (thisObj *utilTestCode) FuncDeferNotCheckError(){
	//defer忽略调用函数所返回的错误
	func(){
		testFunc := func() error {
			return errors.New("testError!")
		}
		defer testFunc()
		fmt.Println("this is test finish!")
		//输出结果:
		//this is test finish!
	}()

	fmt.Println("~~~~~~~~~~~~~~~~~ 处理defer调用函数所返回的错误 ~~~~~~~~~~~~~~~~~")

	//处理defer调用函数所返回的错误
	func(){
		testFunc := func() error {
			return errors.New("testError!")
		}
		//通过匿名函数获取defer调用函数所返回的错误进行处理
		defer func(){
			err := testFunc()
			if err!=nil {
				fmt.Println("deferFuncErr:",err)
			}
		}()
		fmt.Println("this is test finish!")
		//输出结果:
		//this is test finish!
		//deferFuncErr: testError!
	}()

	fmt.Println("~~~~~~~~~~~~~~~~~ 通过具名返回值处理defer调用函数所返回的错误 ~~~~~~~~~~~~~~~~~")

	//通过具名返回值处理defer调用函数所返回的错误
	testFuncError := func() (actionError error){
		testFunc := func() error {
			return errors.New("testError!")
		}
		//通过匿名函数获取defer调用函数所返回的错误进行处理
		defer func(){
			err := testFunc()
			if err!=nil {
				fmt.Println("deferFuncErr:",err)
				actionError = err
			}
		}()
		fmt.Println("this is test finish!")
		return
	}()
	fmt.Println("testFuncError:",testFuncError)
	//输出结果:
	//this is test finish!
	//deferFuncErr: testError!
	//testFuncError: testError!
}

//函数异常处理
//命令行-输入:{"optTag":"Util","optParams":{"methodName":"FuncPanicRecover"}}
func (thisObj *utilTestCode) FuncPanicRecover(){
	func(){
		defer func(){
			fmt.Println("recoverError:",recover())	//捕获有效!
		}()

		defer recover()	//捕获无效!!!

		defer fmt.Println("fmtRecoverError:",recover())	//捕获无效!!!

		defer func(){
			func(){
				fmt.Println("funcFuncRecoverError:",recover()) //捕获无效!!!
			}()
		}()

		panic("this is test panic!")
		//输出结果:
		//funcFuncRecoverError: <nil>
		//fmtRecoverError: <nil>
		//recoverError: this is test panic!
	}()
}

//函数错误处理
//命令行-输入:{"optTag":"Util","optParams":{"methodName":"FuncError"}}
func (thisObj *utilTestCode) FuncError(){
	func(){
		err := fmt.Errorf("errorTag => %d",10)
		fmt.Println("err:",reflect.TypeOf(err).String(),"======",err)

		err1 := errors.New("this is errors!")
		fmt.Println("err1:",reflect.TypeOf(err1).String(),"======",err1)

		//输出结果:
		//err: *errors.errorString ====== errorTag => 10
		//err1: *errors.errorString ====== this is errors!
	}()
}

//GO实现try catch异常处理
//命令行-输入:{"optTag":"Util","optParams":{"methodName":"FuncTryCatch"}}
func (thisObj *utilTestCode) FuncTryCatch(){
	try := func(tryFunc func(),catchFunc func(interface{})){
		defer func(){
			if err:=recover(); err!=nil {
				catchFunc(err)
			}
		}()
		tryFunc()
	}

	try(func(){
		fmt.Println("12321")
		panic("this is try catch panic!")
	},func(err interface{}){
		fmt.Println("err:",err)
	})
}

//方法
//命令行-输入:{"optTag":"Util","optParams":{"methodName":"Method"}}
func (thisObj *utilTestCode) Method(){
	//值类型调用方法
	u1 := method.User{"test1",21}
	u1.Name = "test111"
	u1.Echo2()

	fmt.Println("")
	fmt.Println( "u1.Name:",u1.Name )
	fmt.Println("")

	//指针类型调用方法
	u2 := &method.User{"test2",22}
	u2.Name = "test222"
	u2.Echo2()

	fmt.Println("")
	fmt.Println( "u2.Name:",u2.Name )
	fmt.Println("")
}
//方法-匿名字段
//命令行-输入:{"optTag":"Util","optParams":{"methodName":"MethodAnonymousField"}}
func (thisObj *utilTestCode) MethodAnonymousField(){
	obj := method.TestMethodAnonymousField{method.MethodAnonymousField{"test",123}}
	res := fmt.Sprintf("%p-%v",&obj,&obj)
	fmt.Println(res)
	fmt.Println("~~~~~~~~~~~~~~~")
	obj.Echo()
}
