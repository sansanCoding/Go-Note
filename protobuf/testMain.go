package main

import (
	"Go-Note/protobuf/Test"
	"fmt"
	"github.com/golang/protobuf/proto"
)

/*
	@todo 操作说明:
	@todo 1.必须要有protoc存在,即必须安装protobuf,参考protobuf安装说明.txt
	@todo 2.安装完protobuf后,还需要go get -u github.com/golang/protobuf/protoc-gen-go,针对go编译插件!
	@todo 3.写一个.proto后缀的文件存在
	@todo 4.使用protoc --go_out=将proto解析为go代码文件
	@todo 示例如 protoc --go_out=. --proto_path=. *.proto (--go_out=.意思输出go代码文件到当前目录下)
	@todo ------------------------------------------------------------------------------
	@todo protoc --proto_path=IMPORT_PATH --<lang>_out=DST_DIR path/to/file.proto
	@todo --proto_path=IMPORT_PATH：可以在 .proto 文件中 import 其他的 .proto 文件，proto_path 即用来指定其他 .proto 文件的查找目录。如果没有引入其他的 .proto 文件，该参数可以省略。
	@todo --<lang>_out=DST_DIR：指定生成代码的目标文件夹，例如 –go_out=. 即生成 GO 代码在当前文件夹，另外支持 cpp/java/python/ruby/objc/csharp/php 等语言

	@todo 总结:
	@todo protoc+protoc-gen-go,就可以让.proto文件解析为go代码文件!
*/

//@todo 直接CD进入到Go-Note/protobuf目录下,使用go run testMain.go介入测试!

func main(){
	test := &Test.Test{
		Name:"zhangsan",
		Age:21,
	}

	data,err := proto.Marshal(test)
	fmt.Println("data:",data,string(data))
	fmt.Println("err:",err)

	newTest := &Test.Test{}
	err = proto.Unmarshal(data, newTest)

	fmt.Println("err:",err)

	fmt.Println("test.GetName():",test.GetName())

	fmt.Println("newTest.GetName():",newTest.GetName())

}


