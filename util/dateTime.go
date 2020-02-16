package util

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

//日期时间

type dateTime struct {

}

var DateTime *dateTime

func init(){
	DateTime = NewDateTime()
}

func NewDateTime() *dateTime {
	return &dateTime{

	}
}

//获取当前时间戳-int
func (thisObj *dateTime) CurrTimeInt() int {
	return int(time.Now().Unix())
}

//获取当前时间戳-int64
func (thisObj *dateTime) CurrTimeInt64() int64 {
	return time.Now().Unix()
}

//获取当前日期和时间(字符串值),效果如2019-12-12 22:35:54
func (thisObj *dateTime) CurrDate() string {
	now 	:= time.Now()

	year 	:= now.Year()
	month 	:= now.Month()
	day 	:= now.Day()
	hour 	:= now.Hour()
	minute 	:= now.Minute()
	second 	:= now.Second()

	YmdHis 	:= fmt.Sprintf("%02d-%02d-%02d %02d:%02d:%02d", year, month, day, hour, minute, second)
	return YmdHis
}

//获取当前年月日(字符串值),效果如2019-12-12
//@params string joinStr 拼接的字符串,如-或/
func (thisObj *dateTime) CurrYmd(joinStr string) string {
	now 	:= time.Now()

	year 	:= now.Year()
	month 	:= now.Month()
	day 	:= now.Day()

	ymd 	:= fmt.Sprintf("%02d"+joinStr+"%02d"+joinStr+"%02d", year, month, day)
	return ymd
}

//获取当前年月日十分秒(字符串值)
//@params string ymdJoinStr 年月日拼接的字符串,如-或/
//@params string hisJoinStr 时分秒拼接的字符串,如:
func (thisObj *dateTime) CurrYmdHis(ymdJoinStr string,hisJoinStr string) map[string]string {
	now 	:= time.Now()

	year 	:= now.Year()
	month 	:= now.Month()
	day 	:= now.Day()
	hour 	:= now.Hour()
	minute 	:= now.Minute()
	second 	:= now.Second()

	ymd := fmt.Sprintf("%02d"+ymdJoinStr+"%02d"+ymdJoinStr+"%02d", year, month, day)
	his := fmt.Sprintf("%02d"+hisJoinStr+"%02d"+hisJoinStr+"%02d", hour, minute, second)

	return map[string]string{
		"ymd":ymd,
		"his":his,
	}
}

//获取单项的当前时间-年月日时分秒
func (thisObj *dateTime) CurrYmdHisByItem() map[string]string {
	//获取当前时间的年月日时分秒
	now 	:= time.Now()

	year 	:= now.Year()
	month 	:= now.Month()
	day 	:= now.Day()
	hour 	:= now.Hour()
	minute 	:= now.Minute()
	second 	:= now.Second()

	//拼接效果如:2019-12-15
	ymdStr := fmt.Sprintf("%02d-%02d-%02d", year, month, day)
	//拼接效果如:10:38:37
	hisStr := fmt.Sprintf("%02d:%02d:%02d", hour, minute, second)

	//分割出年月日
	ymdArr := strings.Split(ymdStr,"-")
	//分割出时分秒
	hisArr := strings.Split(hisStr,":")

	//年月日单项值
	ymdHisItem := map[string]string{
		"year":ymdArr[0],
		"month":ymdArr[1],
		"day":ymdArr[2],
		"hour":hisArr[0],
		"minute":hisArr[1],
		"second":hisArr[2],
		"ymd":ymdStr,
		"his":hisStr,
	}

	//年月日+小时的日期时间,如2019-12-15 10:00:00
	ymdHisItem["ymdH"] 	= ymdStr+" "+ymdHisItem["hour"]+":00:00"
	//年月日+小时+分钟的日期时间,如2019-12-15 10:38:00
	ymdHisItem["ymdHi"] = ymdStr+" "+ymdHisItem["hour"]+":"+ymdHisItem["minute"]+":00"

	return ymdHisItem
}

//解析年月日时分秒为时间对象
func (thisObj *dateTime) ParseYmdHisToTimeObj(timeObj *interface{},ymdHis string,timeLayout string) error {
	//本地时间
	localTime, localTimeErr := time.LoadLocation("Asia/Shanghai")
	if localTimeErr!=nil {
		return errors.New("timeLoadLocationErr:"+localTimeErr.Error())
	}

	//日期模板,如果赋值为空字符串,则以默认的为准
	if timeLayout=="" {
		timeLayout = "2006-01-02 15:04:05"
	}

	//获取如 2019-08-19 00:00:00 的时间对象
	timeTime,timeTimeErr := time.ParseInLocation(timeLayout, ymdHis, localTime)
	if timeTimeErr!=nil {
		return errors.New("timeParseInLocationErr:"+timeTimeErr.Error())
	}

	*timeObj = timeTime

	return nil
}

//年月日时分秒转换成时间戳
//@params string ymdHis 年月日时分秒的字符串,如2019-08-26 05:00:00
func (thisObj *dateTime) YmdHisToTimeStamp(ymdHis string) (map[string]interface{}, error) {
	//本地时间
	localTime, localTimeErr := time.LoadLocation("Asia/Shanghai")
	if localTimeErr!=nil {
		return map[string]interface{}{}, localTimeErr
	}

	//日期模板
	timeLayout := "2006-01-02 15:04:05"

	//值如 2019-08-26 05:00:00
	timeStr := ymdHis
	//获取如 2019-08-19 00:00:00 的时间对象
	timeObj, err := time.ParseInLocation(timeLayout, timeStr, localTime)
	if err != nil {
		return map[string]interface{}{}, err
	}
	//获取如2019-08-19 00:00:00 的转换的时间戳,即该年月日的起始时间
	timeInt64 := timeObj.Unix()

	return map[string]interface{}{
		"ymdHisTimeStamp": timeInt64,
	}, nil
}

//年月日转换成时间戳
//@params string ymd 年月日的字符串,如2019-08-19
func (thisObj *dateTime) YmdToTimeStamp(ymd string) (map[string]interface{}, error) {
	//本地时间
	localTime, localTimeErr := time.LoadLocation("Asia/Shanghai")
	if localTimeErr!=nil {
		return map[string]interface{}{}, localTimeErr
	}

	//日期模板
	timeLayout := "2006-01-02 15:04:05"

	//拼接成如 2019-08-19 00:00:00
	timeStr := ymd + " 00:00:00"
	//获取如 2019-08-19 00:00:00 的时间对象
	timeObj, err := time.ParseInLocation(timeLayout, timeStr, localTime)
	if err != nil {
		return map[string]interface{}{}, err
	}
	//获取如2019-08-19 00:00:00 的转换的时间戳,即该年月日的起始时间
	timeInt64 := timeObj.Unix()

	return map[string]interface{}{
		"ymdTimeStamp": timeInt64,
	}, nil
}

//获取时间戳的年月日
//@params int64 timeStamp 时间戳,如1566172800
func (thisObj *dateTime) GetYmdHisByTimeStamp(timeStamp int64) map[string]string {
	//日期模板(模板的数字不可改变)
	timeLayout	 		:= "2006_01_02_15_04_05"
	//将时间戳转换为字符串格式,按日期模板转换如2019_08_20_17_57_04
	timeStampDateStr 	:= time.Unix(timeStamp, 0).Format(timeLayout)

	//分割效果如 [2019 08 20 17 57 04]
	ymdHisStrArr 		:= strings.Split(timeStampDateStr,"_")

	return map[string]string{
		"year":		ymdHisStrArr[0],
		"month":	ymdHisStrArr[1],
		"day":		ymdHisStrArr[2],
		"hour":		ymdHisStrArr[3],
		"minute":	ymdHisStrArr[4],
		"second":	ymdHisStrArr[5],
	}
}

//根据按天日期(YYYY-mm-dd)计算当天的开始和结束时间戳
//@params string dayDate 一天的日期,如2019-08-19
func (thisObj *dateTime) OneDayStartAndEndByDate(dayDate string) (map[string]interface{}, error) {
	//本地时间
	localTime, localTimeErr := time.LoadLocation("Asia/Shanghai")
	if localTimeErr!=nil {
		return map[string]interface{}{}, localTimeErr
	}

	//日期模板
	timeLayout := "2006-01-02 15:04:05"

	//拼接成如 2019-08-19 00:00:00
	timeFromStr := dayDate + " 00:00:00"

	//年月日,拼接成如 20190819
	ymdStr := strings.Join(strings.Split(dayDate, "-"), "")

	//获取如 2019-08-19 00:00:00 的时间对象
	timeFromObj, err := time.ParseInLocation(timeLayout, timeFromStr, localTime)
	if err != nil {
		return map[string]interface{}{}, err
	}

	//获取如2019-08-19 00:00:00 的转换的时间戳,即当天起始时间
	timeFromInt64 := timeFromObj.Unix()
	//获取如2019-08-19 00:00:00 的转换的时间戳+86399秒数,即当天截止时间
	timeToInt64 := timeFromInt64 + 86399

	//将时间戳转换为字符串格式,如2019-08-19 00:00:00 和 2019-08-19 23:59:59
	timeFromDateStr := time.Unix(timeFromInt64, 0).Format(timeLayout)
	timeToDateStr := time.Unix(timeToInt64, 0).Format(timeLayout)

	return map[string]interface{}{
		"ymdStr":          ymdStr,
		"timeFromInt64":   timeFromInt64,
		"timeToInt64":     timeToInt64,
		"timeFromDateStr": timeFromDateStr,
		"timeToDateStr":   timeToDateStr,
	}, nil
}
