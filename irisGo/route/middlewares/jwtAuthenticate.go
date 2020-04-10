package middlewares

//路由-jwt验证中间件

//import (
//	"Go-Note/util"
//	"fmt"
//	"github.com/dgrijalva/jwt-go"
//	jwtmiddleware "github.com/iris-contrib/middleware/jwt"
//	"github.com/kataras/iris"
//	"strings"
//)
//
////jwt验证中间件
//func JwtAuthenticate() *jwtmiddleware.Middleware {
//	//jwt中间件
//	jwtHandler := jwtmiddleware.New(jwtmiddleware.Config{
//		//这个方法将验证jwt的token
//		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
//			return config.TokenKey, nil
//		},
//		//设置后，中间件会验证令牌是否使用特定的签名算法进行签名
//		//如果签名方法不是常量，则可以使用ValidationKeyGetter回调来实现其他检查
//		//加密的方式
//		SigningMethod: jwt.SigningMethodHS256,
//		//验证未通过错误处理方式
//		ErrorHandler: func(ctx iris.Context, s string) {
//			//若请求头中没有包含token相关字段
//			if s == "Required authorization token not found" {
//				util.Helper.HelperResponseFailedJSON(ctx, "需要令牌!", -3)
//				return
//			}
//			//token过期
//			if s == "Token is expired" {
//				util.Helper.HelperResponseFailedJSON(ctx, "令牌已过期", -1)
//				return
//			}
//			//签名方式不正确
//			if strings.Contains(s, "signing method but token specified") {
//				fmt.Println("token签名方式不正确,需要hs256")
//				util.Helper.HelperResponseFailedJSON(ctx, "需要令牌!!", -3)
//				return
//			}
//			//产生其他错误输出
//			fmt.Println(s)
//			//响应到浏览器的body消息
//			util.Helper.HelperResponseFailedJSON(ctx, "需要令牌!!!", -3)
//			return
//		},
//		ContextKey: "jwt",
//		//debug 模式
//		Debug:      false,
//		Expiration: true,
//	})
//	return jwtHandler
//}
//
////token检测处理
//func JwtHandler() iris.Handler {
//	return func(ctx iris.Context) {
//		//获取token对象
//		token := ctx.Values().Get("jwt").(*jwt.Token)
//		//获取app标记
//		appTag := ctx.Params().Get("AppTag")
//		//获取token字符串
//		tokenRaw := token.Raw
//		//获取用户相关信息
//		claim, _ := token.Claims.(jwt.MapClaims)
//		//获取用户id
//		userId := claim["sub"]
//		_, ret := common.LoginCommon.CacheUserToken(ctx, appTag, "check", map[string]interface{}{
//			"userId": userId,
//			"token":  tokenRaw,
//		})
//		//若token验证不通过
//		if !ret {
//			util.Helper.HelperResponseFailedJSON(ctx, "需要令牌!!!!", -3)
//			return
//		}
//		//设置用户id
//		ctx.Values().Set("userId", userId)
//		//执行下一步
//		ctx.Next()
//	}
//}
