package WebSocket
//
//@todo 仅供参考,不作为实际开发代码使用!
//
//import (
//	"Go-Note/util"
//	"database/sql"
//	"fmt"
//	"strconv"
//	"strings"
//	"time"
//	"sync"
//	"encoding/json"
//
//	//@todo 1.运行前的准备工作:
//	//@todo 在main.go中声明一个全局变量 KafkaWebSocketDataCh 并分配内存
//	//声明:kafka消费信息时进行websocket服务端广播客户端
//	//var KafkaWebSocketDataCh chan string
//	//分配内存:存储kafka触发WebSocket推送的数据(该channel是阻塞型的)
//	//KafkaWebSocketDataCh = make(chan string)
//	//kafka消费者代码后面写入该通道数据(这里就以TestData为示例,可以写入任何数据类型)
//	//KafkaWebSocketDataCh <- "TestData"
//
//	//@todo 2.外部开启WebSocket服务端监听的调用示例:
//	//@todo go WebSocket.WebSocket(":39113")	//39113是监听端口
//
//	//@todo 3.需要存在以下2个包代码才可运行此文件代码
//	//"github.com/kataras/iris"
//	//"github.com/kataras/iris/websocket"
//)
//
////协程监听kafka推送数据只能被启动一次
//var wsHandleConnectionChannelDoing sync.Map
////协程监听kafka推送数据只能被启动一次-互斥锁(阻塞型)
//// 不管是1个协程还是多个协程中使用该锁,未解锁前都是阻塞型的;只要其中一个协程抢到该锁,该协程在不解锁前,其他协程也被阻塞着!
//var wsHandleConnectionChannelDoingLockMutex sync.Mutex
//
////WebSocket服务端广播给客户端-锁
//var wsServerBroadcastActionLock sync.Map
//
////启动websocket
//func WebSocket(addr string) {
//
//	// 配置websocket
//	ws := websocket.New(websocket.Config{
//		ReadBufferSize:  40960,
//		WriteBufferSize: 40960,
//	})
//
//	// 监听客户端连接
//	ws.OnConnection(handleConnection)
//
//
//	//@todo 注意:这是kafka推送消息给消费者后,才触发的websocket服务端广播websocket客户端的处理
//	//--------------------------------------- kafka推送触发到websocket服务端进行广播处理 start ---------------------------------------
//	//使用互斥锁,为的是保证只让go for启动一次
//	func(){
//		wsHandleConnectionChannelDoingLockMutex.Lock()
//		defer wsHandleConnectionChannelDoingLockMutex.Unlock()
//
//		//@todo wsHandleConnectionChannelDoing的sync.Map可以换成sync.Once处理，保证只执行一次!
//		//若当前协程已在监听kafka推送数据,则不允许再次启动监听
//		_, existsDoing := wsHandleConnectionChannelDoing.Load("handleConnectionChannel")
//		if !existsDoing {
//			//go for 只允许启动一次
//			wsHandleConnectionChannelDoing.Store("handleConnectionChannel", true)
//
//			//开启协程监听kafka推送数据
//			go func(){  //这个协程是为了让config.KafkaWebSocketDataCh的channel不阻塞继续往下的代码执行
//				for{    //这个for是为了读取channel中多个数据进行的循环处理(只要往这个channel塞多少次数据,这里的for就会循环多少次)
//					// [其实config.KafkaWebSocketDataCh的channel是读取和写入都是阻塞型的,所以for循环看起来是死循环,实际是channel阻塞,没有继续循环,直到有写入channel的数据,这里进行读取了,才会循环下去]
//					//借助channel接收数据
//					kafkaSendString := <- KafkaWebSocketDataCh
//
//					//再开启协程处理多个平台的websocket广播处理
//					go wsServerBroadcastAction(ws, kafkaSendString)
//				}
//			}()
//		}
//	}()
//	//--------------------------------------- kafka推送触发到websocket服务端进行广播处理 end ---------------------------------------
//
//
//	//在端点上注册一个服务
//	app := iris.New()
//
//	//监听多个不同地址的websocket
//	pathList := map[string]interface{}{
//		//这里根据需求可以定义多个,也可以从外部文件获取到
//		"aa":"",
//		"bb":"",
//		"cc":"",
//	}
//	var listenUrl = ""
//	for key, _ := range pathList {
//		listenUrl = "/websocket_statistic_order/" + key
//		app.Get(listenUrl, ws.Handler())
//	}
//
//	app.Run(iris.Addr(addr))
//}
//
//func handleConnection(c websocket.Connection) {
//	//--------------------------------------- 原websocket客户端触发到websocket服务端进行广播处理 start ---------------------------------------
//	{
//		//获取地址标识
//		path := c.Context().Path()
//		lastPosition := strings.LastIndex(path, "/")   //没有返回-1
//
//		//地址标识不存在处理
//		if lastPosition == -1 {
//			ress :=  "地址标识不存在！"
//			mess, _ := json.Marshal(ress)
//			c.To(c.ID()).EmitMessage(mess)
//			return
//		}
//
//		//多个地址写入不同房间
//		//这里的path效果如:websocket_statistic_order/aa
//		//这里的everyPath效果如:aa
//		everyPath := path[lastPosition+1:]
//		//如果该地址没有加入过
//		isJoin := c.IsJoined(everyPath)
//		if !isJoin {
//			c.Join(everyPath)
//		}
//
//		//@todo 注意:这个是浏览器js操作触发的websocket请求,与kafka并不冲突,一个是浏览器触发的websocket请求,一个是kafka消息队列触发的websocket请求
//		//@todo 为什么还保留浏览器触发的weboskcet请求处理，是因为浏览器添加了心跳检测这些机制，还需要进行处理
//		//若有浏览器html页面上触发的客户端websocket请求
//		c.OnMessage(func(msg []byte) {
//			//返回数据
//			res := make(map[string]interface{})
//
//			//获取地址标识(防止乱串数据,所以在此又重新获取一遍,保持最正确的地址标识号)
//			path := c.Context().Path()
//			lastPosition := strings.LastIndex(path, "/")
//			path  := path[lastPosition+1:]
//
//			//请求数据
//			var data map[string]interface{}
//			//Unmarshal只支持[]byte类型
//			json.Unmarshal(msg, &data)
//
//			//1.接收消息后,对应操作类型的数据处理
//			if _, ok := data["type"]; ok {
//				switch data["type"] {
//				//获取提醒状态
//				case "get_order_alert_status":
//					{
//						//涉及数据库查询操作等
//					}
//				//心跳，用户检测websocket是否正常
//				case "heart":
//					{
//						dataMap := make(map[string]interface{})
//						dataMap["type"] = "heart"
//						res["type"] = 1
//						res["data"] = dataMap
//					}
//				//默认输出
//				default:
//					{
//						dataMap := make(map[string]interface{})
//						dataMap["type"] = "StatisticOrder-say hello"
//						res["type"] = 1
//						res["data"] = dataMap
//					}
//				}
//
//				//2.进行webSocket广播推送消息
//				msg, _ := json.Marshal(res["data"])
//				if res["type"] == 2 {
//					//推送给同一地址的所有用户
//					c.To(everyPath).EmitMessage(msg)
//				} else {
//					//推送给当前客户端
//					c.To(c.ID()).EmitMessage(msg)
//				}
//			} else {
//				res["data"] = "请求方法不存在！"
//				mess, _ := json.Marshal(res["data"])
//				c.To(c.ID()).EmitMessage(mess)
//			}
//		})
//
//		//添加关闭的方法，从room删掉用户
//		c.OnDisconnect(func() {
//			//获取地址标识
//			path := c.Context().Path()
//			lastPosition := strings.LastIndex(path, "/")
//			closePath  := path[lastPosition+1:]
//			closeJoin := c.IsJoined(closePath)
//			if closeJoin {
//				c.Leave(closePath)
//			}
//		})
//	}
//	//--------------------------------------- 原websocket客户端触发到服务端进行广播处理 end ---------------------------------------
//
//}
//
////WebSocket服务端广播给客户端
//func wsServerBroadcastAction(ws *websocket.Server, kafkaData string) {
//	//本地消息前綴
//	localMsgPrefix := "wsServerBroadcastAction-"
//
//	//默认地址标记号(一旦产生错误就以默认地址标记号记录)
//	path := "__DEFAULT__"
//	//默认操作类型
//	statType := ""
//	//日志参数
//	logParams := map[string]interface{}{
//		"msg":"default",
//		"kafkaData":kafkaData,
//	}
//
//	//kafka推送数据：xx_xx
//	kafkaDataStrArr := strings.Split(kafkaData, "_")
//	//如果不是按照指定格式拼接的字符串,如a_b_c则是错误的，只能是a_b类似的字符串才可以通过
//	if len(kafkaDataStrArr) != 2 {
//		logParams["msg"] = "kafkaDataStrArr_len_neq_2"
//		wsServerBroadcastActionRunLog(path,logParams)
//		return
//	}
//
//	//推送地址
//	path = kafkaDataStrArr[0]
//	//操作类型
//	statType = kafkaDataStrArr[1]
//
//	//加锁是为了防止同一类型重复操作查询数据库,即同一个地址同一个类型不可重复处理!
//	//锁key,值如:a_b
//	lockKeys := kafkaData
//
//	//捕获异常
//	defer func() {
//		if err := recover(); err != nil {
//			// 解锁--不管成功还是异常等情况,都要解锁
//			wsServerBroadcastActionLock.Store(lockKeys, false)
//			//记录异常栈
//			util.LogFile.PanicTrace(map[string]interface{}{
//				"0.func":localMsgPrefix+"panic!",
//				"1.kafkaData":path,
//				"2.panicErr":err,
//			},4,err)
//		}
//	}()
//
//	//获取锁
//	lock, existLock := wsServerBroadcastActionLock.Load(lockKeys)
//	//若是当前正在锁定中,则不进行处理
//	if existLock && lock.(bool) {
//		logParams["msg"] = "IsLockIng"
//		wsServerBroadcastActionRunLog(path,logParams)
//		return
//	}
//	//设置锁
//	wsServerBroadcastActionLock.Store(lockKeys, true)
//
//	//根据房间号获取链接组,这里是根据指定的房间号获取,如aa则只获取aa的;
//	//  前提必须是先让ws.OnConnection(handleConnection)监听 客户端链接产生(如果客户端链接不产生,就无法使用c.Join方法将平台号作为房间号加入,这里也就无法获取到该房间号的链接组);
//	//  且在handleConnection方法中,将平台号作为房间号加入到链接中,如c.Join(path);
//	//  不然这里获取该房间号链接组都是空的!
//	wsConnections := ws.GetConnectionsByRoom(path)
//	//若websocket客户端链接数小于等于0,则不广播处理
//	if len(wsConnections)<=0 {
//		// 解锁--针对性解锁
//		wsServerBroadcastActionLock.Store(lockKeys, false)
//
//		logParams["msg"] = "wsConnections_len_elt_0"
//		wsServerBroadcastActionRunLog(path,logParams)
//		return
//	}
//
//	//统计结果
//	recordRes := make(map[string]interface{})
//	//根据操作类型进行处理
//	switch statType {
//	//数据库统计逻辑
//	case "xxxx":
//		recordRes = map[string]interface{}{
//			"data":map[string]interface{}{
//				"xx":"",
//				"xxx":"",
//			},
//		}
//	}
//	//推送数据
//	wsSendData := recordRes["data"].(map[string]interface{})
//
//	//链接总数
//	wsConnectionsTotalLen := 0
//	//链接错误总数
//	wsConnectionsErrTotalLen := 0
//	//链接组错误集
//	wsConnectionsEmitErrs := make([]string,0)
//	//推送消息
//	if len(wsSendData) > 0 {
//		//数据转为jsonByte
//		jsonMsgByte, _ := json.Marshal(wsSendData)
//
//		//循环处理每个链接进行广播处理
//		for _,c := range wsConnections {
//			//链接总数自增
//			wsConnectionsTotalLen++
//
//			//发起消息
//			cEmitErr := c.EmitMessage(jsonMsgByte)
//
//			//发起消息若有错误
//			if cEmitErr!=nil {
//				//链接错误总数
//				wsConnectionsErrTotalLen++
//				//1.记录到日志参数中
//				wsConnectionsEmitErrs = append(wsConnectionsEmitErrs,cEmitErr.Error())
//				//2.关闭该链接(这里不是说非得要关闭链接,而是产生了错误,就直接关闭该链接了,是否需要关闭该链接还得要针对错误区别处理)
//				c.Disconnect()
//			}
//		}
//	}
//
//	//操作完成写入日志
//	logParams["msg"] = "WebSocketSendFinish" //WebSocket推送完毕,不管是不是有数据
//	logParams["webSocketSendInfo"] = map[string]interface{}{ //WebSocket推送信息
//		"0.wsSendData":wsSendData,
//		"1.wsConnectionsTotalLen":wsConnectionsTotalLen,
//		"2.wsConnectionsErrTotalLen":wsConnectionsErrTotalLen,
//		"3.wsConnectionsEmitErrs":wsConnectionsEmitErrs,
//	}
//	wsServerBroadcastActionRunLog(path,logParams)
//
//	//解锁--不管成功还是异常等情况,都要解锁
//	wsServerBroadcastActionLock.Store(lockKeys, false)
//
//}
//
////主动推送WebSocket客户端-运行日志
//func wsServerBroadcastActionRunLog(path string,logParams map[string]interface{}){
//	var newLogFileParams util.DetailLogNewFileParams
//	newLogFileParams.DiyLogInfo = "主动推送WebSocket客户端-运行日志"
//	newLogFileParams.FilePath = "WebSocketWsServerBroadcastActionRun.log"
//	newLogFileParams.DirPath = "WebSocketWsServerBroadcastAction/"+path
//	newLogFileParams.LocalMsgPrefix = "websocket_wsServerBroadcastActionRunLog_runLog"
//	util.LogFile.RunLog("Test",map[string]interface{}{
//		"logParams":logParams,
//		"howLongToClean":24,//24小時清除一次日志内容
//	},newLogFileParams)
//}