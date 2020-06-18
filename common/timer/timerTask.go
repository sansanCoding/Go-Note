package timer

//
//import (
//	"fmt"
//	//@todo 需要包含如下包:
//	"github.com/robfig/cron"
//)
//
////定时任务文件
//func InitTimerTask() {
//	//创建定时任务对象
//	c := cron.New()
//
//	fmt.Println("测试定时器-开启测试!")
//	//根据定时时间执行计划任务
//	testCron(c)
//
//	//启动计划任务
//	c.Start()
//	//关闭这计划任务, 但是不能关闭已经在执行中的任务.
//	defer c.Stop()
//	select {}
//
//}
//
////参考分时日月周的定时任务参数的效果:
////https://crontab.guru/#0_0_*_*_*
//
////测试-定时任务
//func testCron(c *cron.Cron) {
//
//	//定时格式说明:秒 分 时 日 月 周
//
//	//每天凌晨5点5分触发
//	c.AddFunc("0 5 5 * * *", func() {
//		fmt.Println("0 5 5 * * *	---test-cron!!!")
//	})
//
//	//每5秒触发一次
//	c.AddFunc("*/5 * * * * *", func() {
//		fmt.Println("*/5 * * * * *	---test-cron!!!")
//	})
//
//	//每5分钟执行一次
//	c.AddFunc("0 */5 * * * *", func() {
//		fmt.Println("0 */5 * * * *	---test-cron!!!")
//	})
//
//	//每天凌晨3点执行一次
//	c.AddFunc("0 0 3 * * *", func() {
//		//fmt.Println("0 0 3 * * *	---test-cron!!!")
//	})
//
//	//每1分钟的第26秒执行一次和第52秒执行一次
//	c.AddFunc("26,52 * * * * *", func() {
//		//fmt.Println("26,52 * * * * *	---test-cron!!!")
//	})
//}