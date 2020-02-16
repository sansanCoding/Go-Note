package util

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"runtime"
	"sync"
	"time"
)

//定义的日志目录,该目录存在于当前项目目录之中,如xxx/src/项目目录/logs/
var LogDir = "logs/"

//日志文件对象map
var logrusObjMap sync.Map

//运行日志数据缓存
var runLogDataCache sync.Map

//日志文件-消息前缀
var logFileMsgPrefix = "util_logFile-"

//公用-详情日志的日志文件参数
type DetailLogNewFileParams struct{
	//本地消息前缀
	LocalMsgPrefix 	string
	//要创建的目录,如:dir1/dir2
	DirPath 		string
	//要创建的文件,如:xxxx_data_detail.log
	FilePath		string
	//diyLog的info信息
	DiyLogInfo		string
}

//私有结构体
type logFile struct {

}

//公用变量
var LogFile *logFile

//初始化执行
func init() {
	LogFile = NewLogFile()
}

func NewLogFile() *logFile {
	return &logFile{

	}
}

//核心:创建日志文件
func (thisObj *logFile) createLogFile(filePath string,data map[string]interface{},info string){
	//缓存对象key:将文件路径设置成md5值
	md5sign 		:= md5.Sum([]byte(filePath))
	filePathKey 	:= fmt.Sprintf("%x", md5sign)

	//获取之前是否创建过该文件对象
	filePathVal,filePathKeyIsExi := logrusObjMap.Load(filePathKey)

	//存在直接写入文件内容,不存在则创建
	if filePathKeyIsExi {

		filePathVal.(*logrus.Logger).WithFields(data).Info(info)

	}else{

		//1.文件操作
		//私有logrus对象
		logrusLogFile := logrus.New()
		//默认输出到os.Stdout中
		logrusLogFile.Out = os.Stdout
		//输出格式为json
		logrusLogFile.Formatter = new(logrus.JSONFormatter)
		logrusLogFile.Level = logrus.DebugLevel
		//打开文件
		// You could set this to any `io.Writer` such as a file
		file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			logrusLogFile.Info(logFileMsgPrefix+"createLogFile:["+filePath+"]Failed to log to file, using default stderr")
			return
		}
		//输出文件
		logrusLogFile.Out = file
		//写入文件内容
		logrusLogFile.WithFields(data).Info(info)

		//2.存储该文件操作的文件对象
		logrusObjMap.Store(filePathKey,logrusLogFile)

	}
}

//------------------------------ 文件操作工具集-(基础)- start ------------------------------
//清空文件内容
func (thisObj *logFile) TruncateFile(filePath string) error {
	return os.Truncate(filePath,0)
}

//递归创建文件夹(调用os.MkdirAll递归创建文件夹)
//注意:该函数递归创建的都是文件夹,哪怕是文件名也会当做文件夹名创建~~~
func (thisObj *logFile) CreateDir(dirPath string) error {
	if !thisObj.IsExistsByFile(dirPath) {
		err := os.MkdirAll(dirPath,os.ModePerm)
		return err
	}
	return nil
}

//判断所给路径文件/文件夹是否存在(返回true是存在)
func (thisObj *logFile) IsExistsByFile(path string) bool {
	_, err := os.Stat(path)    //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

//判断文件夹是否存在(同this.IsExistsByFile()函数处理相同)
func (thisObj *logFile) IsExistsByDir(path string) bool {
	return thisObj.IsExistsByFile(path)
}
//------------------------------ 基础-文件操作工具集-(基础)- end ------------------------------

//------------------------------ 自定义日志文件与目录操作 start ------------------------------
//注意:凡是自定义日志文件或文件夹走的目录都是基于LogDir这个变量值的

//自定义日志文件和写入内容
//注意:***该函数没有根据文件路径相对应的递归创建文件夹***
func (thisObj *logFile) DiyLog(filePath string,data map[string]interface{},info string){
	thisObj.createLogFile(LogDir+filePath,data,info)
}

//自定义文件夹
func (thisObj *logFile) DiyDir(dirPath string) (string,error) {
	dirFullPath := LogDir+dirPath
	return dirFullPath,thisObj.CreateDir(dirFullPath)
}

//判断自定义文件是否存在
func (thisObj *logFile) IsExistsDiyFile(filePath string) bool {
	return thisObj.IsExistsByFile(LogDir+filePath)
}

//清空自定义文件内容
func (thisObj *logFile) TruncateDiyFile(filePath string) error {
	return thisObj.TruncateFile(LogDir+filePath)
}
//------------------------------ 自定义日志文件与目录操作 end ------------------------------

//测试日志文件
func (thisObj *logFile) TestLog(data map[string]interface{},info string){
	thisObj.DiyLog("test.log",data,info)
}

//错误日志
func (thisObj *logFile) ErrorLog(data map[string]interface{},info string){
	thisObj.DiyLog("error.log",data,info)
}

//异常日志
func (thisObj *logFile) ExceptionLog(errParams map[string]interface{},errInfo string){
	//获取调用该方法的所在文件和行数
	//	fileName:绝对路径的所在文件,如xxx/src/xxx/main.go
	//	line:行数,如20
	pc, fileName, line, _ := runtime.Caller(1)

	//获取调用该方法的方法名,如main.main
	//	main.xxx main是包名
	//	xxx.main main是方法名
	funcName := runtime.FuncForPC(pc).Name()

	//写入到日志中
	thisObj.ErrorLog(map[string]interface{}{
		"1.errParams":	errParams,
		"2.errFile": 	fmt.Sprintf("【fileName:%v------line:%v------funcName:%v】", fileName,line,funcName),
	},errInfo)
}

//异常栈记录
//@params map[string]interface{} errParams 写入日志的错误参数
//@params int kb 异常栈的字节长度
//@params interface{} err 类似errors.New()的错误值(即error值)之类
//调用示例:
//	util.LogFile.PanicTrace(map[string]interface{}{
//		"1.tag":"mainTestPanic!",
//		"2.msg":"test",
//	},4,err)
func (thisObj *logFile) PanicTrace(errParams map[string]interface{},kb int, err interface{}) {
	s := []byte("/src/runtime/panic.go")
	e := []byte("\ngoroutine ")
	line := []byte("\n")
	stack := make([]byte, kb<<10) //4KB
	length := runtime.Stack(stack, true)
	start := bytes.Index(stack, s)
	stack = stack[start:length]
	start = bytes.Index(stack, line) + 1
	stack = stack[start:]
	end := bytes.LastIndex(stack, line)
	if end != -1 {
		stack = stack[:end]
	}
	end = bytes.Index(stack, e)
	if end != -1 {
		stack = stack[:end]
	}
	stack = bytes.TrimRight(stack, "\n")

	//存储异常信息
	errInfo := ""
	if stack != nil && err != nil {
		errInfo = fmt.Sprintf("%v \r\n======= stack: =======\r\n %s", err, string(stack))
	} else {
		errInfo = fmt.Sprintf("%v \r\n======= stack: =======\r\n %s", "err is nil", "stack is nil")
	}

	//记录到异常日志中
	thisObj.ExceptionLog(errParams,errInfo)
}

//详情-日志记录(可作为其他方法公用处理)
//@params bool 						isClean 			是否清除日志文件内容,true 是,false 不是
//@params map[string]interface{}	newLogFileParams 	创建日志文件的参数
//@params map[string]interface{}	logParams 			日志参数
func (thisObj *logFile) DetailLog(isClean bool,newLogFileParams DetailLogNewFileParams,logParams map[string]interface{}) {
	//本地消息前缀
	localMsgPrefix 	:= newLogFileParams.LocalMsgPrefix
	//要创建的目录,如:dir1/dir2
	dirPath 		:= newLogFileParams.DirPath
	//要创建的文件,如:xxx_data_detail.log
	filePath		:= newLogFileParams.FilePath
	//diyLog的info信息
	diyLogInfo		:= newLogFileParams.DiyLogInfo

	//捕获异常
	defer func() {
		if err := recover(); err != nil {
			//记录异常栈
			thisObj.PanicTrace(map[string]interface{}{
				"1.func":localMsgPrefix+"-panic!",
				"2.isClean":isClean,
				"3.newLogFileParams":fmt.Sprintf("%+v",newLogFileParams),
				"4.logParams":logParams,
			},4,err)
		}
	}()

	//要写入的日志文件
	logPath	:= filePath

	//1.创建对应的目录
	_,diyDirErr := thisObj.DiyDir(dirPath)
	if diyDirErr!=nil {
		panic(map[string]interface{}{
			"1.msg":localMsgPrefix+"createDirErr!",
			"2.dirPath":dirPath,
			"3.diyDirErr":diyDirErr,
		})
	}

	//组装带有目录路径的日志文件,拼接效果如:dir1/dir2/xxx_data_detail.log
	//注意:上面创建日志时就已保存在logs目录下,所以这里不能使用LogFile.DiyDir()返回的dirFullPath组成全路径
	logFullPath := dirPath+"/"+logPath

	//2.判断 是否清除日志文件内容 且 要写入的日志文件存在不存在;若存在,则清空再写入日志内容.
	if isClean && thisObj.IsExistsDiyFile(logFullPath) {
		truncateFileErr := thisObj.TruncateDiyFile(logFullPath)
		if truncateFileErr!=nil {
			panic(map[string]interface{}{
				"1.msg":localMsgPrefix+"truncateFileErr!",
				"2.logFullPath":logFullPath,
				"3.truncateFileErr":truncateFileErr,
			})
		}
	}

	//3.写入日志内容
	thisObj.DiyLog(logFullPath,map[string]interface{}{
		"1.logParams":logParams,
	},diyLogInfo)
}

//运行日志记录
//@params string 					tag 						标记参数(即哪个功能业务调用的,以作区分之用)
//@params map[string]interface{} 	params 						参数,包含日志参数等其他参数传入
//		  map[string]interface{}	params["logParams"]			日志参数(必传)
//		  int						params["howLongToClean"]	多长时间清除日志一次,默认以 小时 为单位(必传)
//@params struct 					newLogFileParams 			新建日志文件的相关参数
func (thisObj *logFile) RunLog(tag string,params map[string]interface{},newLogFileParams DetailLogNewFileParams) {
	//捕获异常
	defer func() {
		if err := recover(); err != nil {
			//记录异常栈
			thisObj.PanicTrace(map[string]interface{}{
				"1.func":logFileMsgPrefix+"RunLog-panic!",
				"2.tag":tag,
				"3.params":params,
				"4.newLogFileParams":fmt.Sprintf("%+v",newLogFileParams),
			},4,err)
		}
	}()

	//必传-日志参数
	logParams := params["logParams"].(map[string]interface{})
	//必传-多长时间清除日志一次,默认以 小时 为单位
	howLongToClean := params["howLongToClean"].(int)

	//默认不清除日志内容
	isClean := false

	//获取当前时间的年月日时分秒-单项值
	currYmdHisItem 		:= DateTime.CurrYmdHisByItem()
	//获取当前时间的年月日+小时的日期时间,值如2019-12-15 10:00:00
	currYmdHisDateTime 	:= currYmdHisItem["ymdH"]

	//获取日志数据缓存key
	//	拼接效果如:xxx_RunLog
	logDataCacheKey := tag+"_RunLog"
	//获取日志数据缓存结果
	logDataCacheRes,logDataCacheResOk := runLogDataCache.Load(logDataCacheKey)
	//若数据存在,则先判断是否需要清除日志内容
	if logDataCacheResOk {
		//获取日志写入日期时间
		logWriteDate := logDataCacheRes.(map[string]interface{})["logWriteDateTime"].(string)

		//日志写入日期时间-时间对象
		var logWriteDateTimeObj interface{}
		DateTime.ParseYmdHisToTimeObj(&logWriteDateTimeObj,logWriteDate,"")

		//当前年月日时分秒的日期时间-时间对象
		var currYmdHisDateTimeObj interface{}
		DateTime.ParseYmdHisToTimeObj(&currYmdHisDateTimeObj,currYmdHisDateTime,"")

		//判断日志写入日期时间,是否已达清除日志时间
		//	当前日期时间 减去 日志写入日期时间 得出的小时数,是否大于等于 指定的清除日志时长(即小时数)。
		//	若大于等于,即可以清除日志内容;反之则不做处理。
		if int( currYmdHisDateTimeObj.(time.Time).Sub(logWriteDateTimeObj.(time.Time)).Hours() ) >= howLongToClean {
			//设置为清除日志内容
			isClean = true
			//同时将当前的日期时间写入到日志数据缓存中
			runLogDataCache.Store(logDataCacheKey,map[string]interface{}{
				"logWriteDateTime":currYmdHisDateTime,
			})
		}
	}else{ //若数据不存在,则将当天日期写入到日志数据缓存中
		runLogDataCache.Store(logDataCacheKey,map[string]interface{}{
			"logWriteDateTime":currYmdHisDateTime,
		})
	}

	//记录详情日志
	thisObj.DetailLog(isClean,newLogFileParams,logParams)
}
