package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
	"time"
)

//模拟php的curl操作

type curl struct {

}

var Curl *curl

func init(){
	Curl = NewCurl()
}

func NewCurl() *curl {
	return &curl{

	}
}


//请求第三方平台-一般接口
//调用示例:
//	requestApi := "https://www.baidu.com/"
//	requestParams := map[string]interface{}{
//		"apiParams":map[string]string{
//			"UserName":"Test",
//			"PassWord":"123456",
//		},
//		"apiResponseDataType":"map",
//	}
//	requestHeader := map[string]string{}
//	responseData := make(map[string]interface{})
//	requestErr,requestErrParams := util.Curl.Request("post",requestApi,requestParams,requestHeader,&responseData)
//
//	//调试输出:
//	fmt.Println("responseData:",responseData)
//	fmt.Println("requestErr:",requestErr)
//	fmt.Println("requestErrParams:",requestErrParams)
//
//@params string method 请求方法,如get 或 post
//@params string apiUri 请求第三方平台接口url地址
//@params map[string]string params.apiParams 请求接口处理的参数
//@params map[string]string header 请求接口处理的参数
//@return map[string]interface{} actionResult 存储操作结果
func (thisObj *curl) Request(method string, apiUri string, params map[string]interface{}, header map[string]string, actionResult *map[string]interface{}) (actionError error,actionErrorParams map[string]interface{}) {
	//操作结果错误信息
	actionError 		= nil
	//操作结果错误参数,供排查问题用
	actionErrorParams  	= map[string]interface{}{
		"method":method,
		"apiUri":apiUri,
		"params":params,
	}

	//必传-请求接口参数
	apiParams 			:= params["apiParams"].(map[string]string)
	//必传-请求接口后响应数据的数据类型
	apiResponseDataType := params["apiResponseDataType"].(string)

	//发起请求
	var response []byte
	var responseStatusCode int
	switch method {
	case "get":
		{
			//Get方法请求url
			httpGetUrl := Helper.GETUrlPath(apiUri, apiParams, false,false)
			//发起GET请求
			responseResult,responseError,responseErrorParams := thisObj.RequestGetHeader(httpGetUrl, 300*time.Second, header)
			//记录错误参数
			actionErrorParams["httpGetUrl"] = httpGetUrl
			actionErrorParams["requestGetResponseErrorParams"] = responseErrorParams
			if responseError!=nil {
				actionError = errors.New("requestGetResponseError:"+responseError.Error())
				return actionError,actionErrorParams
			}
			response = responseResult["responseByte"].([]byte)
			responseStatusCode = responseResult["responseStatusCode"].(int)
		}
	case "post":
		{
			//发起POST请求
			responseResult,responseError,responseErrorParams := thisObj.RequestPostHeader(apiUri, apiParams, 300*time.Second, header)
			//记录错误参数
			actionErrorParams["requestPostResponseErrorParams"] = responseErrorParams
			if responseError!=nil {
				actionError = errors.New("requestPostResponseError:"+responseError.Error())
				return actionError,actionErrorParams
			}
			response = responseResult["responseByte"].([]byte)
			responseStatusCode = responseResult["responseStatusCode"].(int)
		}
	}

	//若没有响应消息
	if len(response) <= 0 {
		actionError = errors.New("requestResponseByte_length_elt_0")
		return actionError,actionErrorParams
	}

	//记录错误参数
	actionErrorParams["requestResponseByteToString"] = string(response)

	//存储响应数据
	var responseData interface{}
	var responseDataLength int

	//根据指定的响应数据类型进行确认
	switch apiResponseDataType {
	//单个map数据
	case "map":
		{
			//响应消息
			responseData = make(map[string]interface{})
			//响应消息
			responseJsonUnmarshalError := json.Unmarshal(response, &responseData)
			//若反解析json失败,则返回错误信息
			if responseJsonUnmarshalError!=nil {
				actionError = errors.New("requestResponseByteToJsonUnmarshalError:"+responseJsonUnmarshalError.Error())
				return actionError,actionErrorParams
			}
			//响应数据长度
			responseDataLength = len(responseData.(map[string]interface{}))
		}
	//列表map数据
	case "listInterface":
		{
			//响应消息
			responseData = make([]interface{},0)
			//响应消息反解析json
			responseJsonUnmarshalError := json.Unmarshal(response, &responseData)
			//若反解析json失败,则返回错误信息
			if responseJsonUnmarshalError!=nil {
				actionError = errors.New("requestResponseByteToJsonUnmarshalError:"+responseJsonUnmarshalError.Error())
				return actionError,actionErrorParams
			}
			//响应数据长度
			responseDataLength = len(responseData.([]interface{}))
		}
	//interface数据
	case "interface":
		{
			//响应消息
			//var responseData interface{}
			//响应消息反解析json
			responseJsonUnmarshalError := json.Unmarshal(response, &responseData)
			//若反解析json失败,则返回错误信息
			if responseJsonUnmarshalError!=nil {
				actionError = errors.New("requestResponseByteToJsonUnmarshalError:"+responseJsonUnmarshalError.Error())
				return actionError,actionErrorParams
			}

			//响应数据长度
			switch responseData.(type) {
			case []interface{}:
				responseDataLength = len(responseData.([]interface{}))
			case map[string]interface{}:
				responseDataLength = len(responseData.(map[string]interface{}))
			//若最后都找不到对应的数据类型则也是错误的
			default:
				actionError = errors.New("requestResponseDataType("+fmt.Sprint(reflect.TypeOf(responseData))+")_notFound")
				return actionError,actionErrorParams
			}
		}
	}

	//操作结果存储
	(*actionResult)["responseData"] 		= responseData
	(*actionResult)["responseDataLength"] 	= responseDataLength
	(*actionResult)["responseStatusCode"] 	= responseStatusCode

	return actionError,actionErrorParams
}


//------------------------------------------------- Request start -------------------------------------------------

//大部分的Request调用示例如下:

////GET-请求示例:
////请求接口的api地址
//	apiUri := "http://192.168.0.0.1:55001/api/user/login"
////请求接口的api参数
//	apiParams := map[string]string{
//		"UserName":"test",
//		"PassWord":"123456",
//	}
////请求接口的header头设置
//	header := map[string]string{
//		"Authorization":"xxxxxxxxxxx",
//	}
//	//或
//	//header := map[string]string{
//  //}
////发起请求
//	curlRes,curlResErr,curlResErrParams := util.Curl.RequestGetHeader(apiUri, apiParams, 300*time.Second, header)
//	if curlResErr!=nil {
//		panic("curlResErr:"+curlResErr.Error())
//	}
////输出结果:
//	if len(curlRes["responseByte"])<=0 {
//		fmt.Println("curlResponseNotData!")
//	}else{
//		fmt.Println(string(curlRes["responseByte"]))
//  }

////POST-请求示例:
//	dataStr := "http://192.168.0.100:54090/api/user/login"
//	apiParams := map[string]string{
//		//类似url请求地址,相关字符进行转义处理
//		"data":url.QueryEscape(dataStr),	//url.QueryEscape(dataStr)效果如:http%3A%2F%2F192.168.0.100%3A54090%2Fapi%2Fuser%2Flogin
//	}
//	util.Curl.RequestPostHeader(apiUri, apiParams, 300*time.Second, header)

//模拟PHP-curl-GET请求,可设置header头信息
func (thisObj *curl) RequestGetHeader(urlPath string, timeout time.Duration, header map[string]string) (actionResult map[string]interface{},actionError error,actionErrorParams map[string]interface{}) {
	//操作结果
	actionResult = map[string]interface{}{
		"responseByte":[]byte(""),
		"responseStatusCode":0,
	}
	//操作错误
	actionError = nil
	//操作错误参数
	actionErrorParams = map[string]interface{}{
		"params":map[string]interface{}{
			"urlPath":urlPath,
			"timeout":timeout,
			"header":header,
		},
	}

	//请求相关配置
	client := &http.Client{
		Timeout: timeout,
	}

	//发起请求准备
	req, err := http.NewRequest("GET", urlPath, strings.NewReader(""))
	if err != nil {
		actionError = errors.New("httpNewRequestErr:"+err.Error())
		return actionResult,actionError,actionErrorParams
	}

	//设置header头信息
	for k, v := range header {
		req.Header.Set(k, v)
	}

	//执行请求
	resp, err := client.Do(req)
	if err != nil {
		actionError = errors.New("httpNewRequestClientDo:"+err.Error())
		return actionResult,actionError,actionErrorParams
	}

	//执行完毕,关闭资源
	defer resp.Body.Close()

	//操作成功存储-响应状态码
	actionResult["responseStatusCode"] = resp.StatusCode

	//读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		actionError = errors.New("httpNewRequestIoutilReadAll:"+err.Error())
		return actionResult,actionError,actionErrorParams
	}

	//操作成功存储-响应消息体
	actionResult["responseByte"] = body

	return actionResult,actionError,actionErrorParams
}

//模拟PHP-curl-POST请求,可设置header头信息
func (thisObj *curl) RequestPostHeader(urlPath string, data map[string]string, timeout time.Duration, header map[string]string) (actionResult map[string]interface{},actionError error,actionErrorParams map[string]interface{}) {
	//操作结果
	actionResult = map[string]interface{}{
		"responseByte":[]byte(""),
		"responseStatusCode":0,
	}
	//操作错误
	actionError = nil
	//操作错误参数
	actionErrorParams = map[string]interface{}{
		"params":map[string]interface{}{
			"urlPath":urlPath,
			"data":data,
			"timeout":timeout,
			"header":header,
		},
	}

	//请求相关配置
	client := &http.Client{
		Timeout: timeout,
	}

	//或者使用这种方式拼接
	//	dataUrlVal := url.Values{}
	//	for key, val := range data {
	//		dataUrlVal.Add(key, val)
	//	}
	//	httpBuildQuery := dataUrlVal.Encode()

	//准备请求参数整合
	httpBuildQuery := ""
	for k, v := range data {
		//如果传进来的是已经拼接好的,就放入map,v的值就是拼接好的,k为__IS_ONE__代表直接按v的值获取一次即可
		if len(data) == 1 && k == "__IS_ONE__" {
			httpBuildQuery = v
		} else {
			httpBuildQuery += k + "=" + v + "&"
		}
	}
	if httpBuildQuery != "" {
		httpBuildQuery = strings.TrimRight(httpBuildQuery, "&")
	}

	//记录参数
	actionErrorParams["httpBuildQuery"] = httpBuildQuery

	//发起请求准备
	req, err := http.NewRequest("POST", urlPath, strings.NewReader(httpBuildQuery))
	if err != nil {
		actionError = errors.New("httpNewRequestErr:"+err.Error())
		return actionResult,actionError,actionErrorParams
	}

	//若是按json请求,若没有指定对应的请求数据格式,则默认以post数据请求
	if _,ok:=header["__REQUEST_JSON__"]; ok {
		req.Header.Set("Content-Type", "application/json")
		delete(header,"__REQUEST_JSON__")
	}else{
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	//设置header头信息
	for k, v := range header {
		req.Header.Set(k, v)
	}

	//执行请求
	resp, err := client.Do(req)
	if err != nil {
		actionError = errors.New("httpNewRequestClientDo:"+err.Error())
		return actionResult,actionError,actionErrorParams
	}

	//执行完毕,关闭资源
	defer resp.Body.Close()

	//操作成功存储-响应状态码
	actionResult["responseStatusCode"] = resp.StatusCode

	//读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		actionError = errors.New("httpNewRequestIoutilReadAll:"+err.Error())
		return actionResult,actionError,actionErrorParams
	}

	//操作成功存储-响应消息体
	actionResult["responseByte"] = body

	return actionResult,actionError,actionErrorParams
}

//模拟PHP-curl-POST.JSON数据请求,可设置header头信息
func (thisObj *curl) RequestPostJsonHeader(urlPath string, data string, timeout time.Duration, header map[string]string) (actionResult map[string]interface{},actionError error,actionErrorParams map[string]interface{}) {
	//key有用,值暂时没用到
	header["__REQUEST_JSON__"] = ""
	return thisObj.RequestPostHeader(urlPath,map[string]string{"__IS_ONE__":data},timeout,header)
}

//------------------------------------------------- Request end -------------------------------------------------