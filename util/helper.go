package util

import (
	"bytes"
	"crypto/md5"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"
)

//helper工具集

type helper struct {

}

var Helper *helper

func init() {
	Helper = NewHelper()
}

func NewHelper() *helper {
	return &helper{

	}
}

//获取当前时间戳-int类型数据
func (thisObj *helper) CurrTimeInt() int {
	return int(time.Now().Unix())
}

//json解析
//@params interface{} t 基本上传参都是map数据类型,如map[string]interface{}或是map[string]string
func (thisObj *helper) JSONMarshal(t interface{}) ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)	//HTML标签是否转义处理 => true:{"test":"\u003cbr /\u003e"} false:{"test":"<br />"}
	err := encoder.Encode(t)
	return buffer.Bytes(), err
}

//创建随机字符串
//@params int keyLength 要创建字符串的长度
//@params int typeN 1 按数字随机,非1 按0~9,a~z,A~Z随机
func (thisObj *helper) CreateRandStr(randStrLength int, typeN int) string {
	key := ""
	var dic = make([]string, 0)
	if typeN == 1 {
		dic = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	} else {
		dicStr := "a,b,c,d,e,f,g,h,i,j,k,l,m,n,o,p,q,r,s,t,u,v,w,x,y,z,1,2,3,4,5,6,7,8,9,0,A,B,C,D,E,F,G,H,I,J,K,L,M,N,O,P,Q,R,S,T,U,V,W,X,Y,Z"
		dic = strings.Split(dicStr, ",")
	}
	length := len(dic)
	for i := 0; i < randStrLength; i++ {
		rand.Seed(time.Now().UnixNano())
		randNum := rand.Intn(length - 1)
		key += dic[randNum]
	}
	return key
}

//http响应json-成功数据
//@params string message 响应消息
//@params interface{} data 响应数据,一般也是map数据类型
func (thisObj *helper) HttpResponseJsonSuccess(message string, data interface{}) {
	//http响应结果
	result := map[string]interface{}{"code": "success", "msg": message, "data": data}
	//设置头信息
	//xxx.Http.SetHeader("Content-Type", "application/json;charset=utf-8")
	//解析成json输出
	res, _ := json.Marshal(result)
	if len(res) > 4*1024 {
		//xxx.Http.WriteGzip(res)
	} else {
		//xxx.Http.Write(res)
	}
}

//失败响应处理
//需要引用包: import ( "github.com/kataras/iris" )
//func (thisObj *helper) HelperResponseFailedJSON(c iris.Context, message string, code interface{}) {
//	//获取前置中间件的设置的初始开始时间
//	diffTime := c.Values().Get("requestCurrentTime")
//	currentTime := time.Now().UnixNano() / 1e3 //计算得出 微秒
//	timeConsumed := currentTime - diffTime.(int64)
//	result := iris.Map{"code": code, "msg": message, "status": 0, "time_consumed": timeConsumed}
//	c.JSON(result)
//}

//http响应seqMap的Json-成功数据
//@params string message 响应消息
//@params string data 响应数据,解析成为json的字符串数据,如`{"test":"123"}`
func (thisObj *helper) HttpResponseSeqJsonSuccess(message string, data string) {
	//seqMap存储数据
	result := NewSeqMap()
	result.Put("code", "success")
	result.Put("msg", message)
	result.Put("data", data)
	//设置头信息
	//xxx.Http.SetHeader("Content-Type", "application/json;charset=utf-8")
	//解析成json输出
	res := []byte(result.JsonSeq())
	if len(res) > 4*1024 {
		//xxx.Http.WriteGzip(res)
	} else {
		//xxx.Http.Write(res)
	}
}

//http响应seqMap的Json（转换成Unicode字符）-成功数据
//@params string message 响应消息
//@params string data 响应数据,解析成为json的字符串数据,如`{"test":"123"}`
func (thisObj *helper) HttpResponseSeqUnicodeJsonSuccess(message string, data string) {
	//seqMap存储数据
	result := NewSeqMap()
	result.Put("code", "success")
	result.Put("msg", message)
	result.Put("data", data)
	//设置头信息
	//xxx.Http.SetHeader("Content-Type", "application/json;charset=utf-8")
	//解析成json输出
	res := []byte(result.JsonSeqUnicode())
	if len(res) > 4*1024 {
		//xxx.Http.WriteGzip(res)
	} else {
		//xxx.Http.Write(res)
	}
}

//http响应json
//@params map[string]interface{} data 响应数据
func (thisObj *helper) HttpResponseJSON(data map[string]interface{}) {
	//xxx.Http.SetHeader("Content-Type", "application/json;charset=utf-8")
	res, _ := json.Marshal(data)
	if len(res) > 4*1024 {
		//xxx.Http.WriteGzip(res)
	} else {
		//xxx.Http.Write(res)
	}
}

//根据字符串调用对应对象的方法
//调用示例:
//	type Test struct {
//
//	}
//
//	func NewTest()*Test{
//		return &Test{
//
//		}
//	}
//
//	func (thisObj *Test) Echo() int {
//		return 123456
//	}
//	res,resOk := util.Helper.CallMethodReflect(NewTest(),"Echo")
//	fmt.Println(res,resOk)
//	fmt.Println(res[0].Int())
//输出结果:
//	[<int Value>] true
//	123456
//@params interface{} any 指定对象(指针)
//@params string methodName 指定对象里被调用的方法
//@params []interface{} args 指定对象里被调用的方法传参
//@return []reflect.Value reflectValue 被调用对象返回的结果集
//@return bool runOk 调用是否成功,true 成功,false 失败(失败情况就得根据情况排查问题出在哪里)
func (thisObj *helper) CallMethodReflect(any interface{}, methodName string, args []interface{}) (reflectValue []reflect.Value, runOk bool) {
	//准备调用方法的参数
	inputs := make([]reflect.Value, len(args))
	for i := range args {
		inputs[i] = reflect.ValueOf(args[i])
	}

	//执行调用指定对象的方法
	v := reflect.ValueOf(any).MethodByName(methodName)
	//若指定对象被调用的方法不存在或无效
	if v.String() == "<invalid Value>" {
		runOk = false
		return []reflect.Value{}, runOk
	}

	//调用成功
	runOk = true
	reflectValue = v.Call(inputs)
	return reflectValue, runOk
}

//获取用户IP
//@desc X-Forwarded-For:若是代理过来的，会获取到多个ip，最后一个ip就是真实的
func (thisObj *helper) GetIP(req *http.Request) string {
	remoteAddr := req.RemoteAddr
	if ip := req.Header.Get("X-Real-IP"); ip != "" {
		remoteAddr = ip
	} else if ip = req.Header.Get("X-Forwarded-For"); ip != "" {
		//原逻辑:容易获取多个ip
		//remoteAddr = ip
		//现逻辑:获取最后一个真实ip
		//使用X-Forwarded-For获取到ip会有多个,如117.136.39.96, 183.232.151.36;修改成只获取最后一个.
		//若多个ip存在时,按英文逗号分隔
		ipArr := strings.Split(ip,",")
		//若存在多个ip值时,需要去除空格；获取最后一位IP
		ip := Trim(ipArr[len(ipArr)-1])
		//先判断获取最后一位IP是否是公网IP，如果不是就获取第一个IP处理
		if !thisObj.CheckIsPublicIP(net.ParseIP(ip)) {
			ip = Trim(ipArr[0])
		}
		remoteAddr = ip
	} else {
		remoteAddr, _, _ = net.SplitHostPort(remoteAddr)
	}

	if remoteAddr == "::1" {
		remoteAddr = "127.0.0.1"
	}
	return remoteAddr
}
//检测IP是否是公网IP
func (thisObj *helper) CheckIsPublicIP(IP net.IP) bool {
	if IP.IsLoopback() || IP.IsLinkLocalMulticast() || IP.IsLinkLocalUnicast() {
		return false
	}
	if ip4 := IP.To4(); ip4 != nil {
		switch true {
		case ip4[0] == 10:
			return false
		case ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 31:
			return false
		case ip4[0] == 192 && ip4[1] == 168:
			return false
		default:
			return true
		}
	}
	return false
}
//把IP127.0.0.1格式转化为long类型
func (thisObj *helper) IP2long(ipStr string) uint32 {
	str := net.ParseIP(ipStr)
	if str == nil {
		return 0
	}
	//IPv4
	ip := str.To4()
	if ip == nil {
		//IPv6
		ip = str.To16()
	}
	if ip == nil {
		return 0
	}
	return binary.BigEndian.Uint32(ip)
}
//把long类型IP地址转化为127.0.0.1格式
func (thisObj *helper) Long2IP(ipLong uint32) string {
	ipByte := make([]byte, 4)
	binary.BigEndian.PutUint32(ipByte, ipLong)
	ipStr := net.IP(ipByte)
	return ipStr.String()
}
//检测是否是IPv4
func (thisObj *helper) IsIPv4(ip string) bool {
	//返回的值不等于nil,就是如果ip不是IPv4地址,To4返回nil,当且仅当地址是IPv4地址时,此条件才应返回true.
	if net.ParseIP(ip).To4()!=nil {
		return true
	}
	return false
}
//检测是否是IPv6
func (thisObj *helper) IsIPv6(ip string) bool {
	if thisObj.IsIPv4(ip) {
		return false
	}
	return true
}
////多个IP地址集合解析(包含IPv6地址解析),如58.177.183.37,134.159.119.48
//func (thisObj *helper) IP2Region(params map[string]interface{}) (actionResult []map[string]interface{}) {
//	//必传-多个ip地址集合
//	ips := params["ips"].(string)
//	//可选-是否需要排重
//	isUnique := false
//	if params["isUnique"]!=nil {
//		isUnique = params["isUnique"].(bool)
//	}
//	//去除ip左右多出的逗号
//	ips = strings.Trim(ips,",")
//
//	//多个ip地址数组
//	ipsArr := strings.Split(ips,",")
//	//ip排重
//	ipsMap := make(map[string]map[string]interface{})
//
//	//IPV4-解析
//	ip2regionLibrary := NewIp2regionLibrary()
//	defer ip2regionLibrary.Close()
//	//IPV6-解析
//	ipv6regionLibrary := NewIpv6regionLibrary(map[string]interface{}{})
//	defer ipv6regionLibrary.Close()
//
//	//操作结果存储
//	actionResult = make([]map[string]interface{},0)
//	for _,ip := range ipsArr {
//		//操作结果每项数据
//		actionResultItem := map[string]interface{}{
//			"ipStr":ip,
//			"ipAddr":"---",
//		}
//
//		//------------ ip地址值转换处理 ------------
//		//模拟php的代码:
//		//	有可能传递过来的是ip2long转换的整型ip值 转换成ip字符串
//		//	$ipStr = is_numeric($ip) ? long2ip($ip) : $ip;
//		ipStr := ""
//		//若是整型值
//		if thisObj.IsNumeric(ip) {
//			//转成字符串
//			ipInt,ipIntErr := strconv.Atoi(ip)
//			//若转换出错
//			if ipIntErr!=nil {
//				actionResultItem["ipAddr"] = "IP_TO_INT_ERROR"
//				actionResult = append(actionResult,actionResultItem)
//				continue
//			}
//			//将整型值解析成ip地址
//			ipStr = thisObj.Long2IP(uint32(ipInt))
//		}else{
//			//存储ip地址
//			ipStr = ip
//		}
//
//		//确认都是字符串ip返回
//		actionResultItem["ipStr"] = ipStr
//
//		//是否排重
//		if isUnique {
//			//------------ 排重处理 ------------
//			//若该IP已处理过,则不进行处理
//			if ipsMap[ipStr]!=nil {
//				continue
//			}
//		}else{
//			//------------ 不排重处理 ------------
//			//若是不排重,且之前找到过了,则不需要再重复找了
//			if ipsMap[ipStr]!=nil {
//				actionResult = append(actionResult,ipsMap[ipStr])
//				continue
//			}
//		}
//
//		//------------ ip地址解析处理 ------------
//		//如果是IPv4的处理
//		if thisObj.IsIPv4(ipStr) {
//			//IPV4地址解析处理
//			ipv4Res,ipv4ResErr := ip2regionLibrary.DoIpAddressAnalyze(ipStr)
//			if ipv4ResErr!=nil {
//				actionResultItem["ipAddr"] = "IPV4_ADDRESS_ANALYZE_ERROR"
//			}else{
//				//调试输出:
//				//fmt.Println("ipStr:",ipStr)
//				//fmt.Println("ipv4Res:",ipv4Res)
//				//输出结果:
//				//ipStr: 134.159.119.47
//				//ipv4Res: map[City:香港 CityId:0 Country:中国 ISP:澳大利亚电信 Province:香港 Region:0]
//
//				if len(ipv4Res)>0 {
//					actionResultItem["ipAddr"] = ipv4Res["Country"].(string)+"-"+ipv4Res["Province"].(string)+"-"+ipv4Res["City"].(string)
//				}else{
//					actionResultItem["ipAddr"] = "IPV4_ADDRESS_NOT_FOUND"
//				}
//			}
//		//如果是IPv6的处理
//		}else if thisObj.IsIPv6(ipStr) {
//			ipv6Res,ipv6ResErr := ipv6regionLibrary.GetIPv6Address(map[string]interface{}{
//				"ipv6":ipStr,
//			})
//			if ipv6ResErr!=nil {
//				actionResultItem["ipAddr"] = "IPV6_ADDRESS_ANALYZE_ERROR"
//			}else{
//				//调试输出:
//				//fmt.Println("ipStr:",ipStr)
//				//fmt.Println("ipv6Res:",ipv6Res)
//				//输出结果:
//				//ipStr: 2408:84e7:4ad:1776:9c97:e79d:1e94:e59e
//				//ipv6Res: map[address:中国山西省 中国联通3GNET网络(全省通用) id:50440 ipv6:2408:84e7:4ad:1776:9c97:e79d:1e94:e59e]
//
//				if len(ipv6Res)>0 {
//					actionResultItem["ipAddr"] = ipv6Res["address"].(string)
//				}else{
//					actionResultItem["ipAddr"] = "IPV6_ADDRESS_NOT_FOUND"
//				}
//			}
//		//既不是IPv4,也不是IPv6,则是错误的
//		}else{
//			actionResultItem["ipAddr"] = "IP_IS_NOT_V4_OR_V6"
//		}
//
//		//存储操作结果
//		actionResult = append(actionResult,actionResultItem)
//
//		//存储查找过的结果
//		ipsMap[ipStr] = actionResultItem
//	}
//
//	return actionResult
//}

// IsNumeric is_numeric()
// Numeric strings consist of optional sign, any number of digits, optional decimal part and optional exponential part.
// Thus +0123.45e6 is a valid numeric value.
// In PHP hexadecimal (e.g. 0xf4c3b00c) is not supported, but IsNumeric is supported.
func (cthis *helper) IsNumeric(val interface{}) bool {
	switch val.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return true
	case float32, float64, complex64, complex128:
		return true
	case string:
		str := val.(string)
		//现逻辑
		//@update 2020.03.21 新增若是传入空格的修正处理
		//	若这里不进行空格去除处理,到下面894行执行的时候会报错 runtime error: index out of range [0] with length 0,
		// 	因为等strings.TrimSpace(str)执行去除空格后,str[0]是不可能有值的!
		str = strings.TrimSpace(str)
		if str == "" {
			return false
		}
		//原逻辑
		// Trim any whitespace
		//str = strings.TrimSpace(str)
		if str[0] == '-' || str[0] == '+' {
			if len(str) == 1 {
				return false
			}
			str = str[1:]
		}
		// hex
		if len(str) > 2 && str[0] == '0' && (str[1] == 'x' || str[1] == 'X') {
			for _, h := range str[2:] {
				if !((h >= '0' && h <= '9') || (h >= 'a' && h <= 'f') || (h >= 'A' && h <= 'F')) {
					return false
				}
			}
			return true
		}
		// 0-9, Point, Scientific
		p, s, l := 0, 0, len(str)
		for i, v := range str {
			if v == '.' { // Point
				if p > 0 || s > 0 || i+1 == l {
					return false
				}
				p = i
			} else if v == 'e' || v == 'E' { // Scientific
				if i == 0 || s > 0 || i+1 == l {
					return false
				}
				s = i
			} else if v < '0' || v > '9' {
				return false
			}
		}
		return true
	}

	return false
}

//GET请求url-路径拼接
//@params string urlPath 请求url,如http://192.168.0.1:58909
//@params map[string]string data 请求参数
//@params bool isEncode 特殊字符是否转义处理
//@params bool isHttp 是否将https转换成http处理
func (thisObj *helper) GETUrlPath(urlPath string, data map[string]string, isEncode bool, isHttp bool) string {
	//是否是http请求,如果是,则将https替换成http
	if isHttp {
		urlPath = strings.Replace(urlPath, "https://", "http://", 1)
	}

	//拼接uri
	urlPath += "?"
	for k, v := range data {
		if isEncode {
			v = url.QueryEscape(v)
		}
		urlPath += k + "=" + v + "&"
	}
	urlPath = strings.TrimRight(urlPath, "&")

	return urlPath
}

//用户名称尾部隐藏处理
//@params string userName 用户名
func (thisObj *helper) UserNameTailHidden(userName string) string {
	usernameShow := ""

	//先将字符串转成rune
	//	rune 等同于int32,常用来处理unicode或utf-8字符
	userNameRune := []rune(userName)
	//获取用户名称字符串长度(不是底层字节长度,如你好,就是2个字符串长度)
	userNameLen := len(userNameRune)
	//大于3个长度,则从尾部截取3个替换成***
	if userNameLen > 3 {
		usernameShow = string(userNameRune[:userNameLen-3])
	} else {
		switch userNameLen {
		case 3:
			usernameShow = string(userNameRune[:userNameLen-2])
		case 2:
			usernameShow = string(userNameRune[:userNameLen-1])
		//若是1个长度,则以该值为起始,效果如:张***
		case 1:
			usernameShow = userName
		}
	}

	return usernameShow + "***"
}

//总页数相关处理
//@params int page 当前页
//@params int pageRows 每页数量
//@params int dataTotal 数据总数
func (thisObj *helper) PageTotal(page int, pageRows int, dataTotal int) map[string]int {
	//计算总页数
	totalPages := int(math.Ceil(float64(dataTotal) / float64(pageRows)))
	//总页数最小值限制
	if totalPages < 1 {
		totalPages = 1
	}
	//当前页最小值限制
	if page < 1 {
		page = 1
	}
	//当前页最大值限制
	if page > totalPages {
		page = totalPages
	}
	//计算每页提取数
	limitStart := (page - 1) * pageRows
	return map[string]int{
		"totalPages":  totalPages,
		"currentPage": page,
		"limitStart":  limitStart,
	}
}

//根据用户id获取订单号
func (thisObj *helper) GetOrderIdByUserId(userId int,orderIdLen int) string {
	//当前时间的时分秒+微妙
	timeString 	:= time.Now().Format("150405.000000")
	//拼接当前用户id
	str 		:= timeString + strconv.Itoa(userId)
	//转成md5值
	md5str 		:= fmt.Sprintf("%x", md5.Sum([]byte(str)))
	//生成指定位数强唯一的订单号
	orderId 	:= string([]byte(md5str)[:orderIdLen])

	return orderId
}

//根据findKey查找RequestParams值
//@params interface{} 	requestParams 	get或post的请求参数
//@params string		findKey			要查找的key
//@params string		findValDataType 查找值的数据类型
//@params interface{}	defaultVal 		默认值
func (thisObj *helper) FindRequestParamsVal(requestParams interface{},findKey string,findValDataType string,defaultVal interface{}) (interface{},error) {
	switch requestParams.(type) {
	//如get请求参数
	case map[string]string:
		{
			//根据findKey找值
			findVal,findValExi := requestParams.(map[string]string)[findKey]
			//若值不存在,则已默认值返回
			if !findValExi {
				return defaultVal,errors.New("val_not_found")
			}
			//找到值的数据类型
			switch findValDataType {
			case "int":
				return InterfaceToInt(findVal)
			case "string":
				return InterfaceToStr(findVal)
			}
		}
	//如post请求参数
	case map[string]interface{}:
		{
			//根据findKey找值
			findVal,findValExi := requestParams.(map[string]interface{})[findKey]
			//若值不存在,则已默认值返回
			if !findValExi {
				return defaultVal,errors.New("val_not_found")
			}
			//找到值的数据类型
			switch findValDataType {
			case "int":
				return InterfaceToInt(findVal)
			case "string":
				return InterfaceToStr(findVal)
			}
		}
	}

	panic("requestParams_dataType_notFound")
}

//模拟刪除map[string]interface{}里的元素值
//适用于delete()函数执行删除后,造成原map值也跟着被删除,解决该问题的处理
func (thisObj *helper) DeleteElementByMapStrInterface(mapData map[string]interface{},deleteKey string) map[string]interface{} {
	mapDataTemp := make(map[string]interface{})
	for k,v := range mapData {
		//若是与删除key相等,则不进行存储
		if k==deleteKey {
			continue
		}
		mapDataTemp[k] = v
	}
	return mapDataTemp
}

//error对象转字符串输出
func (thisObj *helper) ErrorToString(err error) string {
	if err!=nil {
		return err.Error()
	}
	return "nil"
}