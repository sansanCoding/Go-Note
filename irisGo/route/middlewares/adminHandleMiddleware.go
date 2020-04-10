package middlewares

//路由-后台中间件

//import (
//	"Go-Note/util"
//	"encoding/json"
//	"github.com/kataras/iris"
//)
//
////后台中间件
//func AdminHandle() (handler iris.Handler) {
//	return func(ctx iris.Context) {
//		//后台请求过来的一律走post请求,所以这里的ctx.FormValue()获取的是post表单数据
//		token := ctx.FormValue("__Token__")
//		requestData := ctx.FormValues()
//		//如果token是空的,则输出错误信息
//		if token == "" {
//			util.Helper.HelperResponseFailedJSON(ctx, "token_empty", 0)
//			return
//		}
//		//存储反解析的参数
//		decryptParams := make(map[string]interface{}, 0)
//		//这里直接是按指定token字符串进行验证,也可以作类似如AES加解密的验证
//		if token == "xxxxxxxxxxx" {
//			//若指定的请求参数字段不存在
//			if requestData["__Token_Data__"] == nil {
//				util.Helper.HelperResponseFailedJSON(ctx, "tokenData_empty", 0)
//				return
//			}
//			tokenData := requestData["__Token_Data__"].(string)
//
//			//反解析json数据
//			json.Unmarshal([]byte(tokenData), &decryptParams)
//		} else {
//			util.Helper.HelperResponseFailedJSON(ctx, "token_error", 0)
//			return
//		}
//
//		//设置值
//		//将请求过来的参数重新设置到控制器所需的获取请求参数函数中
//		for field, val := range decryptParams {
//			ctx.ViewData(field, val)
//		}
//
//		//获取值(即获取ctx.ViewData(field, val)设置的全部值)
//		//获取后台请求过来时的参数(具体参考adminHandleMiddleware.go里的逻辑)
//		//postData:= ctx.GetViewData()
//
//		//执行下一步
//		ctx.Next()
//	}
//}
