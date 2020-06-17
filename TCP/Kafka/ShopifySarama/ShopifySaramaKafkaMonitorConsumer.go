package ShopifySarama

//import (
//	"Go-Note/config"
//	"encoding/json"
//	"fmt"
//	"strconv"
//	"sync"
//
//	//@todo 需要引用到kafka-go版本的包 Shopify/sarama
//	//"github.com/Shopify/sarama"
//)
//
//var (
//	wg sync.WaitGroup
//	msgPrefix = "ShopifySarama-"
//	commonAddrRelease 	= "cache.common.server:50001"
//	commonAddrTest 		= "localhost:50001"
//	commonTopic			= "cacheCommon"
//)
//
////调用示例:
////1.在main.go写入调用
////	//异步启动kafka消费监听
////	go ShopifySarama.KafkaConsumerMonitor()
//
////启动kafka消费监听
//func KafkaConsumerMonitor() {
//	//本地消息前缀
//	localMsgPrefix := msgPrefix+"KafkaConsumerMonitor-"
//
//	var addr string
//	var topic string
//
//	//开发环境配置
//	addr = commonAddrRelease
//	//走本地环境的kafka服务器
//	if config.IsRelease {
//		addr = commonAddrTest
//	}
//	//主题订阅,默认订阅所有应用服务器上的
//	//主要是使用在备份服务器上，备用服务器为所有应用服务器共用
//	topic = commonTopic
//
//	fmt.Println(localMsgPrefix+"kafka消费监听地址:"+addr+",订阅的主题:"+topic)
//
//	//创建kafka消费对象
//	consumer,err := sarama.NewConsumer([]string{addr}, nil)
//	if err != nil {
//		panic(localMsgPrefix+"saramaNewConsumerErr:"+err.Error())
//	}
//	//获取该topic所有分区---consumer.Partitions():获取返回给定主题的所有分区id的排序列表
//	partitionList,err := consumer.Partitions(topic)
//	if err != nil {
//		panic(localMsgPrefix+"consumerPartitionsErr:"+err.Error())
//	}
//	//循环分区列表---到这一步的循环都是同步等待处理的
//	for partition := range partitionList {
//		//获取该主题该分区的最新消息
//		pc, err := consumer.ConsumePartition(topic, int32(partition), sarama.OffsetNewest)
//		if err != nil {
//			panic(localMsgPrefix+"consumer.ConsumePartitionErr:"+strconv.FormatInt(int64(int32(partition)),10)+"-"+err.Error())
//		}
//
//		wg.Add(1)
//
//		//异步处理每个分区消息
//		go func(pc sarama.PartitionConsumer) {
//			defer wg.Done()
//			defer pc.AsyncClose()
//			//获取通道消息,阻塞直到有消息发送过来,然后再继续等待
//			for msg := range pc.Messages() {
//				//fmt.Printf("Partition:%d, Offset:%d, key:%s, value:%s\n", msg.Partition, msg.Offset, string(msg.Key),string(msg.Value))
//				kafkaMessageDispatch(string(msg.Key), string(msg.Value))
//			}
//		}(pc)
//	}
//	wg.Wait()
//	consumer.Close()
//}
//
///*
//string(msg.Value)的字符串是一个json串:
//{
//"table_operation":"update",
//"data":{
//	"id":1000,
//	"UserName":"ZhangSan",
//	"Age":20
//},
//"action":"table_Users",
//"pushTime":1553170216
//}
//*/
//
////kafka消息处理
//func kafkaMessageDispatch(key string, value string) {
//	//本地消息前缀
//	localMsgPrefix := msgPrefix+"kafkaMessageDispatch-"
//
//	data := make(map[string]interface{})
//	valueByte := []byte(value)
//	err := json.Unmarshal(valueByte, &data)
//	if err != nil {
//		fmt.Println("kafkaMessageDispatch-JsonUnmarshalErr:"+err.Error())
//		return
//	}
//	action := data["action"].(string)
//	switch action {
//	case "table_Users":
//		//相关用户信息处理
//		//actionUsers(data)
//	}
//}