package file

//文件下载

//方式1-借助net/http包进行下载处理:
//
//参考文章:
//https://lequ7.com/Golang-xia-zai-wen-jian.html
//
//示例1-简单代码:
//	go func(){
//		//引用包:import ( "net/http" )
//
//		//指定单个文件下载(需要打开ServeFile该方法后点击查看ServeFile方法中如何调用的示例)
//		//http.ServeFile()
//
//		//文件列表下载
//		http.ListenAndServe(":58099", http.FileServer(http.Dir("./logs")))
//		fmt.Println("文件服务监听58099端口已启动!")
//	}()
//
//示例2:
//	package main
//
//	import (
//	"fmt"
//	"io/ioutil"
//	"net/http"
//	"net/url"
//	"os"
//	)
//
//	func main() {
//		http.HandleFunc("/", downloadHandler) //   设置访问路由
//		http.ListenAndServe(":8080", nil)
//	}
//	func downloadHandler(w http.ResponseWriter, r *http.Request) {
//		r.ParseForm() //解析url传递的参数，对于POST则解析响应包的主体（request body）
//		//注意:如果没有调用ParseForm方法，下面无法获取表单的数据
//		fileName := r.Form["filename"] //filename  文件名
//		path := "/data/images/"        //文件存放目录
//		file, err := os.Open(path + fileName[0])
//		if err != nil {
//			fmt.Println(err)
//			return
//		}
//		defer file.Close()
//		content, err := ioutil.ReadAll(file)
//		fileNames := url.QueryEscape(fileName[0]) // 防止中文乱码
//		w.Header().Add("Content-Type", "application/octet-stream")
//		w.Header().Add("Content-Disposition", "attachment; filename=\""+fileNames+"\"")
//
//		if err != nil {
//			fmt.Println("Read File Err:", err.Error())
//		} else {
//			w.Write(content)
//		}
//	}
//
//示例3:
//使用IrisGo框架时的发送文件处理:
//	apiV1.Get("/test", func(ctx iris.Context) {
//		//方式1:下载文件(ctx.SendFile支持大文件下载,可以不指定header)
//		//err := ctx.SendFile("./logs/test.csv","test.csv")
//		//fmt.Println("err:",err)
//
//		//方式2:下载文件(ctx.ServerFile必须要有指定的header)
//		ctx.Header("Content-Type","application/octet-stream")
//		ctx.Header("Content-Disposition","attachment; filename=\"test1.csv\"")
//		err := ctx.ServeFile("./logs/test.csv",false)
//		fmt.Println("err:",err)
//	})