package MongoDB


//------------------------------------------ 一些MongoDB查询语句 start ------------------------------------------
//import(
//	"go.mongodb.org/mongo-driver/bson"
//	"go.mongodb.org/mongo-driver/mongo/options"
//)
//
//	ctxTimer, _ := context.WithTimeout(context.Background(), 1*time.Second)
//	collection := MongoDb.Database("test").Collection("totalTemp")
//	result := make(bson.M)
//	cur := collection.FindOne(ctxTimer,
//		bson.M{"lastRefreshTime": bson.D{{"$gt", 0}}},
//		options.FindOne().SetProjection(bson.M{"_id": 0}))
//	//如果不为空就是连接或者查询异常
//	if cur.Err() != nil && cur.Err().Error() != "mongo: no documents in result" {
//		fmt.Println("查询mongodb失败", cur.Err().Error())
//		return
//	}
//	decodeErr := cur.Decode(&result)
//	if decodeErr != nil {
//		if decodeErr.Error() == "mongo: no documents in result" {
//			result["lastRefreshTime"] = int64(0)
//		} else {
//			fmt.Println("mongodb数据解析失败", decodeErr)
//			return
//		}
//	}
//	var lastRefreshTime = int64(result["lastRefreshTime"].(int64))
//------------------------------------------ 一些MongoDB查询语句 end ------------------------------------------

//------------------------------------------ 一些MongoDB插入语句 start ------------------------------------------
//	var ctx = context.Background()
//	client := MongoDb
//	collection := client.Database("test").Collection("totalTemp")
//	insertDocument := make([]interface{}, 0)
//	for _, v := range mapRecords {
//		insertDocument = append(insertDocument, v)
//	}
//	insertDocument = append(insertDocument, bson.M{"lastRefreshTime": currentTime})
//	if err := collection.Drop(ctx); err != nil {
//		fmt.Println("表删除失败")
//		return
//	}
//	_, err2 := collection.InsertMany(ctx, insertDocument)
//	if err2 != nil {
//		fmt.Println("数据插入失败")
//		return
//	}
//------------------------------------------ 一些MongoDB插入语句 end ------------------------------------------

//@todo 未完待续