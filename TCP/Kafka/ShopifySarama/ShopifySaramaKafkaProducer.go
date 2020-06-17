package ShopifySarama

//import (
//	"Go-Note/config"
//	"encoding/json"
//	"fmt"
//	"time"
//
//	//@todo 需要引用到kafka-go版本的包 Shopify/sarama
//	//"github.com/Shopify/sarama"
//)
//
////@add 2019.10.12 临时调试改为由调用时创建处理,调用完毕后关闭通道处理
////var Producer sarama.AsyncProducer
//
////创建kafka生产者
//func CreateProducer() sarama.AsyncProducer {
//	//本地消息前缀
//	localMsgPrefix := msgPrefix+"CreateProducer-"
//
//	var addr string
//
//	//开发环境配置
//	addr = commonAddrRelease
//	//走本地环境的kafka服务器
//	if config.IsRelease {
//		addr = commonAddrTest
//	}
//
//	//设置配置
//	configProducer := sarama.NewConfig()
//	//等待服务器所有副本都保存成功后的响应
//	configProducer.Producer.RequiredAcks = sarama.NoResponse
//	//随机的分区类型
//	configProducer.Producer.Partitioner = sarama.NewRandomPartitioner
//	//是否等待成功和失败后的响应,只有上面的RequireAcks设置不是NoReponse这里才有用.
//	//config.Producer.Return.Successes = true
//	//config.Producer.Return.Errors = true
//	//使用配置,新建一个异步生产者
//	producer,err := sarama.NewAsyncProducer([]string{addr}, configProducer)
//	if err != nil {
//		panic(localMsgPrefix+"saramaNewAsyncProducerErr:"+err.Error())
//	}
//
//	//Producer = producer
//	return producer
//}
//
////发送生产者消息
//func SendProducerMessage(topic string, key string, message []byte) {
//	//@add 2019.10.12 临时调试改为由调用时创建处理
//	producer := CreateProducer()
//
//	//发送的消息,主题,key
//	msg := &sarama.ProducerMessage{
//		Topic: topic,
//		Key:   sarama.StringEncoder(key),
//		Value: sarama.ByteEncoder(message),
//	}
//	producer.Input() <- msg
//
//	//@add 2019.10.12 临时调试改为由调用时创建处理,调用完毕并销毁
//	producer.AsyncClose()
//}
//
////生产数据
//func SendProducerData(params map[string]interface{}) {
//	//本地消息前缀
//	localMsgPrefix := msgPrefix+"SendProducerData-"
//
//	//捕获异常
//	defer func() {
//		if err := recover(); err != nil {
//			panic(fmt.Sprintf(localMsgPrefix+"panicErr:%+v",map[string]interface{}{
//				"0.err":err,
//				"1.params":params,
//			}))
//		}
//	}()
//
//	//发送消息的主题
//	topic := commonTopic
//
//	//要发送的消息key与val
//	//sendKeys与sendVals的关联是:sendKeys的每一个值就是sendVals的下标
//	sendKeys := params["sendKeys"].([]string)
//	sendVals := params["sendVals"].(map[string]interface{})
//
//	//循环发送消息组
//	for _, key := range sendKeys {
//		val := sendVals[key]
//		dataPush := map[string]interface{}{
//			"action":   key,               	//接口名称
//			"data":     val,        		//接口数据
//			"pushTime": time.Now().Unix(), 	//推送时间
//		}
//		message, _ := json.Marshal(dataPush)
//		SendProducerMessage(topic, key, message)
//	}
//}