package MongoDB

//------------------------------------------ MongoDB注册链接函数 start ------------------------------------------
//package pool
//
//import (
//	"context"
//	"fmt"
//	"time"
//	//@todo 需要引用如下包
//	"go.mongodb.org/mongo-driver/mongo"
//	"go.mongodb.org/mongo-driver/mongo/options"
//	"go.mongodb.org/mongo-driver/mongo/readpref"
//)
//
//var MongoDb *mongo.Client
//
////外部直接调用-初始化mongodb
////@todo 如main.go文件中写入MongodbInit(),相当于提前创建好链接,后面直接用,无需再重复创建
////@todo 跟mysql链接池和redis链接池创建逻辑相似!
//func MongodbInit() {
//	ctxTime, _ := context.WithTimeout(context.Background(), 10*time.Second)
//	client, err := mongo.Connect(ctxTime,
//		options.Client().ApplyURI("mongodb://192.168.0.225:27017").SetMaxPoolSize(50),
//	)
//	err2 := client.Ping(ctxTime, readpref.Primary())
//	if err != nil || err2 != nil {
//		fmt.Println("mongodb链接创建失败,排查配置是否正确以及是否正常运行和可用!")
//		fmt.Println("pingErr:"+err.Error())
//		fmt.Println("pingErr2:"+err2.Error())
//		return
//	}
//	MongoDb = client
//	fmt.Println("mongodb链接成功......")
//}
//------------------------------------------ MongoDB注册链接函数 end ------------------------------------------
