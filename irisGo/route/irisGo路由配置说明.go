package route

//引用包
//import (
// "github.com/kataras/iris"
//)

//第一步声明:
//	app := iris.New()
//
//第二步注册相关路由:
//	//注册模板文件路径
//	app.RegisterView(iris.HTML("./views", ".html"))
//
//	//静态文件路径
//	app.StaticWeb("/assets", "./assets")
//	//直接访问输出ok
//	app.Get("/", func(ctx iris.Context) { ctx.WriteString("ok") })
//	//设置中间件
//	apiRoute := app.Party(
//		"/{appTag:string}",//访问的路由第一个/的值是appTag,访问路由示例:http://localhost:8098/IOS/api/v1/login
//		middlewares.CorsMiddleware(),   //跨域中间件
//		middlewares.BeforeMiddleware(), //前置中间件
//	).AllowMethods(iris.MethodOptions)
//
//	//单例
//	{
//		//生成二维码,访问路由示例:http://localhost:8098/IOS/qrcode/123456
//		apiRoute.Get("/qrcode/{text:string}", func(ctx iris.Context) { controllers.HtmlController.QrcodeText() })
//	}
//
//	//非登录路由
//	apiV1 := apiRoute.Party("/api/v1")
//	{
//		//登录,访问路由示例:http://localhost:8098/IOS/api/v1/login
//		apiV1.Post("/login", func(ctx iris.Context) { controllers.UserController(ctx).Login() })
//	}
//
//	//需要登录之后才能访问的路由,走jwt中间件验证token处理,访问路由示例:http://localhost:8098/IOS/api/v1/auth/userInfo
//	apiV1Auth := apiRoute.Party("/api/v1/auth", middlewares.JwtAuthenticate().Serve, middlewares.JwtHandler())
//	{
//		//获取用户信息
//		apiV1Auth.Get("/userInfo", func(ctx iris.Context) { controllers.UserController(ctx).UserInfo() })
//	}
//
//	//请求来自后台的处理
//	adminHandle := apiRoute.Party("/admin_handle", middlewares.AdminHandle())
//	{
//		//清除后台配置缓存
//		adminHandle.Post("/clear_admin_config_cache", func(ctx iris.Context) { controllers.AdminController(ctx).ClearAdminConfigCache() })
//	}