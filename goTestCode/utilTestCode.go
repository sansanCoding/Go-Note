package goTestCode

import (
	"Go-Note/util"
	"fmt"
	"net/url"
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
