package util

//@todo 只作获取POST数据代码参考,不作实际生产环境使用!

//import (
//	"bytes"
//	"errors"
//	"fmt"
//	"github.com/buger/jsonparser"
//	"github.com/kataras/iris"
//	"io"
//	"io/ioutil"
//	"net/url"
//	"strings"
//)


//--------------------------------------------------------------------------------------------------------
//	//获取请求参数示例:
//	requestData := ctx.(iris.Context).GetViewData()
//	requestDataOk := false
//	//若不是后台请求过来的
//	if len(requestData) == 0 {
//		//获取POST参数(移动端参数获取)
//		requestData, requestDataOk = util.PostPutDeleteRequestParamCompatibility(ctx)
//		if !requestDataOk {
//			//获取GET参数(WEB端参数获取)
//			requestData, requestDataOk = util.GetRequestParam(ctx)
//			if !requestDataOk {
//				util.Helper.HelperResponseFailedJSON(ctx, "获取参数失败!", 0)
//				return
//			}
//		}
//	}
//--------------------------------------------------------------------------------------------------------


//获取POST,PUT,DELETE请求参数
//func PostPutDeleteRequestParamCompatibility(ctx iris.Context) (map[string]interface{}, error) {
//	contentType := ctx.GetHeader("Content-type")
//	var paramMap = make(map[string]interface{}, 0)
//	bodyByte, _ := ioutil.ReadAll(ctx.Request().Body) //把body内容读出
//	//如果没有参数
//	if len(bodyByte) <= 0 {
//		return nil, errors.New("bodyByteLen_elt_0")
//	}
//	//ios或者安卓进来的
//	if contentType == "application/x-www-form-urlencoded" {
//		stringBody := bytes.Split(bodyByte, []byte("&"))
//		for _, v := range stringBody {
//			keyValue := bytes.Split(v, []byte("="))
//			if len(keyValue) <= 1 {
//				return nil, errors.New("stringBody.keyValueLen_elt_1")
//			}
//			//特殊字符解码
//			newS, _ := url.QueryUnescape(string(keyValue[1]))
//			paramMap[string(keyValue[0])] = newS
//		}
//	} else {
//		err := jsonparser.ObjectEach(bodyByte, func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
//			paramMap[string(key)] = string(value)
//			return nil
//		})
//		if err != nil {
//			return nil, errors.New(fmt.Sprintf("jsonparser.ObjectEach-格式解析错误:%v", err))
//		}
//	}
//	return paramMap, nil
//}

////检查用户的请求头信息
//func CheckUserAgent(ctx iris.Context) int {
//	// 1android,2ios,3pc,4windowsphone
//	uag := strings.ToLower(ctx.GetHeader("User-Agent"))
//	var osType int = 3
//	if strings.Contains(uag, "android") || strings.Contains(uag, "okhttp/3.11.0") { // 安卓框架问题，接口请求的ua:okhttp/3....
//		osType = 1
//	}
//	if strings.Contains(uag, "iphone;") || strings.Contains(uag, "ipad;") || strings.Contains(uag, "ipod;") || strings.Contains(uag, "ios") {
//		osType = 2
//	}
//	if strings.Contains(uag, "windows phone") {
//		osType = 4
//	}
//	return osType
//}

//获取GET参数
//func GetRequestParam(ctx iris.Context) (map[string]interface{}, bool) {
//	var paramMap = make(map[string]interface{})
//	ctx.Request().ParseForm()
//	if len(ctx.Request().Form) > 0 {
//		for k, v := range ctx.Request().Form {
//			paramMap[k] = v[0]
//		}
//	}
//	if len(paramMap) > 0 {
//		return paramMap, true
//	} else {
//		return paramMap, false
//	}
//}

///**
//* @param ctx iris.Contex
//* @return map[string]interface{}, bool
//* when POST commit and HEAD set multipart/form-data, shall use this func compatible(android,IOS,H5,PC)
//*/
//func PostRequestMultiForm(ctx iris.Context)  (map[string]interface{}, error)  {
//	contentType := ctx.GetHeader("Content-type")
//	var paramMap = make(map[string]interface{}, 0)
//	postContentType := strings.Split(contentType, ";")
//	if postContentType[0] == "multipart/form-data" {
//		multiPartReader, readErr := ctx.Request().MultipartReader()
//		if readErr != nil {
//			return nil, errors.New(fmt.Sprintf("MultipartReaderErr:%v", readErr))
//		}
//		for {
//			part, nextErr := multiPartReader.NextPart()
//			if nextErr == io.EOF {
//				break
//			}
//			if nextErr != nil {
//				return nil, errors.New(fmt.Sprintf("NextPart %v", nextErr))
//				//break
//			}
//			formName := part.FormName()
//			fileName := part.FileName()
//			if formName != "" && fileName == "" {
//				formValue,_ := ioutil.ReadAll(part)
//				paramMap[formName] = formValue
//			}
//			if fileName != "" {
//				fileData,_ := ioutil.ReadAll(part)
//				paramMap[formName] = fileData
//			}
//		}
//		// @todo here is bug, if it is file commit, shall except
//		for key, _ := range paramMap  {
//			paramMap[key] = string(paramMap[key].([]byte))
//		}
//		return paramMap, nil
//	} else {
//		return PostPutDeleteRequestParamCompatibility(ctx)
//	}
//}
