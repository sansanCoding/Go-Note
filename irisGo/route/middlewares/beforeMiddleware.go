package middlewares

//路由-前置中间件

//import (
//	"Go-Note/util"
//	"github.com/kataras/iris"
//	"time"
//)
//
////前置中间件
//func BeforeMiddleware() (handler iris.Handler) {
//	handler = func(ctx iris.Context) {
//		//获取app标识
//		appTag := ctx.Params().Get("appTag")
//		//请求时间开始计时
//		ctx.Values().Set("requestCurrentTime", time.Now().UnixNano()/1e3)
//		//若不是指定的标识，则提示错误
//		if appTag!="xxx" {
//			util.Helper.HelperResponseFailedJSON(ctx, "错误的APP标识", 0)
//			return
//		}
//		ctx.Next()
//	}
//	return handler
//}
