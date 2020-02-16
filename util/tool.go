package util

import (
	//"github.com/mojocn/base64Captcha"
	//"github.com/shopspring/decimal"
	"bytes"
	"errors"
	"fmt"
	"html"
	"math"
	"net/url"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

//工具集(可直接使用util.方法名调用)

//有部分类似php的方法参考来自于: "github.com/syyongx/php2go"

//interface转整型
func InterfaceToInt(data interface{}) (int,error) {
	dataStr,dataStrErr := JsonValToStr(data)
	if dataStrErr!=nil {
		return 0,dataStrErr
	}
	dataInt,dataIntErr := strconv.Atoi(dataStr)
	if dataIntErr!=nil {
		return 0,dataIntErr
	}
	return dataInt,nil
}

//interface转字符串
func InterfaceToStr(data interface{}) (string,error) {
	return JsonValToStr(data)
}

//Json值转字符串
func JsonValToStr(s interface{}) (string,error) {
	switch s.(type){
	case string:
		return s.(string),nil
	case int:
		return strconv.Itoa(s.(int)),nil
	case int8:
		s = int64(s.(int8))
		return strconv.FormatInt(s.(int64),10),nil
	case int16:
		s = int64(s.(int16))
		return strconv.FormatInt(s.(int64),10),nil
	case int32:
		s = int64(s.(int32))
		return strconv.FormatInt(s.(int64),10),nil
	case int64:
		return strconv.FormatInt(s.(int64),10),nil
	case float32:
		return strconv.FormatFloat(float64(s.(float32)), 'f', -1, 32),nil
	case float64:
		// 'b' (-ddddp±ddd，二进制指数)
		// 'e' (-d.dddde±dd，十进制指数)
		// 'E' (-d.ddddE±dd，十进制指数)
		// 'f' (-ddd.dddd，没有指数)
		// 'g' ('e':大指数，'f':其它情况)
		// 'G' ('E':大指数，'f':其它情况)
		return strconv.FormatFloat(s.(float64), 'f', -1, 64),nil
	}

	return "",errors.New( "JsonValToStr: "+fmt.Sprint(s)+" 数据类型 "+fmt.Sprint(reflect.TypeOf(s))+" 无法强转成字符串类型")
}


//@todo 未实践,请勿使用!
//// 创建验证码(base64code) - 可以有几种方式: 数字/声音/公式
//func CreateCaptcha(width int ,height int) (string, string,string) {
//
//	//声音验证码配置
//	//var configA = base64Captcha.ConfigAudio{
//	//	CaptchaLen: 6,
//	//	Language:   "zh",
//	//}
//	//创建声音验证码
//	//GenerateCaptcha 第一个参数为空字符串,包会自动在服务器一个随机种子给你产生随机uiid.
//	//idKeyA, capA := base64Captcha.GenerateCaptcha("", configA)
//	//以base64编码
//	//base64stringA := base64Captcha.CaptchaWriteToBase64Encoding(capA)
//
//	codeCharacter := func() (string, string,string) { //字符,公式,验证码配置
//		//current := time.Now().Second() % 4
//		current:=3
//		codeMode := base64Captcha.CaptchaModeNumber
//		//const CaptchaModeNumber:数字,CaptchaModeAlphabet:字母,CaptchaModeArithmetic:算术,CaptchaModeNumberAlphabet:数字字母混合.
//		switch current {
//		case 0:
//			codeMode = base64Captcha.CaptchaModeAlphabet
//		case 1:
//			codeMode = base64Captcha.CaptchaModeNumberAlphabet
//		case 2:
//			codeMode = base64Captcha.CaptchaModeArithmetic
//		}
//		var config = base64Captcha.ConfigCharacter{
//			Height:             height,
//			Width:              width,
//			Mode:               codeMode,
//			ComplexOfNoiseText: base64Captcha.CaptchaComplexLower,
//			ComplexOfNoiseDot:  base64Captcha.CaptchaComplexLower,
//			IsShowHollowLine:   false,
//			IsShowNoiseDot:     false,
//			IsShowNoiseText:    false,
//			IsShowSlimeLine:    false,
//			IsShowSineLine:     false,
//			CaptchaLen:         4,
//			IsUseSimpleFont:    true,
//		}
//		//GenerateCaptcha 第一个参数为空字符串,包会自动在服务器一个随机种子给你产生随机uiid.
//		idKey, cap := base64Captcha.GenerateCaptcha("", config)
//		ci := cap.(*base64Captcha.CaptchaImageChar)
//		//以base64编码
//		base64String := base64Captcha.CaptchaWriteToBase64Encoding(cap)
//		return idKey, base64String,ci.VerifyValue
//	}
//
//	codeNum := func() (string, string,string) { //数字验证码
//		var config = base64Captcha.ConfigDigit{
//			Height:     height,
//			Width:      width,
//			MaxSkew:    0.7,
//			DotCount:   80,
//			CaptchaLen: 4,
//		}
//		//创建数字验证码.
//		//GenerateCaptcha 第一个参数为空字符串,包会自动在服务器一个随机种子给你产生随机uiid.
//		idKey, cap := base64Captcha.GenerateCaptcha("", config)
//		ci := cap.(*base64Captcha.CaptchaImageDigit)
//
//		//以base64编码
//		base64String := base64Captcha.CaptchaWriteToBase64Encoding(cap)
//		return idKey, base64String,ci.VerifyValue
//	}
//	startCodeNum:=false
//	if startCodeNum{
//		return codeNum()
//	}
//	//if time.Now().Second()%2 == 0 { //随机生成验证码
//	//	return codeNum()
//	//}
//	return codeCharacter()
//}

//@todo 未实践,请勿使用!
////insert字段多行插入组装
//func UtilMysqlFiledInsertMultiRow(data []map[string]interface{}) string {
//	fields := " ("
//	vals := ""
//	//这里的字段顺序会变，因为map是无序的,所以放入一个数组，用于跟values一一对应
//	mapSort := make([]string, 0)
//	for k, _ := range data[0] {
//		fields += k + ","
//		mapSort = append(mapSort, k)
//	}
//	fields = strings.TrimRight(fields, ",") + ") values "
//	for _, vData := range data {
//		val := "("
//		for _, field := range mapSort {
//			fieldValue := vData[field]
//			vType := reflect.TypeOf(fieldValue).String()
//			if vType == "decimal.Decimal" {
//				val += fieldValue.(decimal.Decimal).String()
//			} else if vType == "string" {
//				val += "'" + fieldValue.(string) + "'"
//			} else {
//				byteV, _ := json.Marshal(fieldValue)
//				val += string(byteV)
//			}
//			val += ","
//		}
//		vals = strings.TrimRight(val, ",") + "),"
//	}
//	vals = strings.TrimRight(vals, ",")
//	return fields + vals
//}


//--------------------------------------- 数组指针处理 start ---------------------------------------
//调用示例:
//	//模拟php的end()函数
//	arrPointer 	:= util.NewArrayPointer([]int{1,2,3,})
//	endVal 		:= arrPointer.End().(int)

type arrayPointer struct {
	//要处理的数组
	arr []interface{}
	//移动的索引值
	index int
}

//模拟php的end(),next(),current(),key()等函数
//current() - 返回数组中的当前单元
//each() 	- 返回数组中当前的键／值对并将数组指针向前移动一步
//prev() 	- 将数组的内部指针倒回一位
//reset() 	- 将数组的内部指针指向第一个单元
//next() 	- 将数组中的内部指针向前移动一位
func NewArrayPointer(data interface{}) *arrayPointer {
	arrayPointer := new(arrayPointer)
	//数据转换
	arr := make([]interface{}, 0)
	switch data.(type) {
	//转换不同的数据类型为[]interface{}
	case []string:
		for _, v := range data.([]string) {
			arr = append(arr, v)
		}
	case []int:
		for _, v := range data.([]int) {
			arr = append(arr, v)
		}
	default:
		panic("NewArrayPointer=>data.(type)_not_found=>data.(type) is " + fmt.Sprint(reflect.TypeOf(data)))
	}

	arrayPointer.arr = arr
	return arrayPointer
}

//将数组的内部指针指向最后一个单元
func (this *arrayPointer) End() interface{} {
	//获取最后一位索引值
	endKey := len(this.arr) - 1
	//这里写循环处理是因为要考虑到,索引值越界
	for k, v := range this.arr {
		if k == endKey {
			//存储最后一位索引值-即指针指向最后一位
			this.index = k
			return v
		}
	}
	return nil
}

//返回数组中内部指针指向的当前单元的键名
func (this *arrayPointer) Key() interface{} {
	//这里写循环处理是因为要考虑到,索引值越界
	for k := range this.arr {
		//根据索引值返回key值
		if this.index == k {
			return k
		}
	}
	return nil
}
//--------------------------------------- 数组指针处理 end ---------------------------------------

//获取int数组的所有值
//	会新创建[]int返回,这样做是为了对形参不是引用效果
func ArrayIntValues(arr []int) (newArrayInt []int) {
	for _, v := range arr {
		newArrayInt = append(newArrayInt, v)
	}
	return
}
//获取string数组的所有值
func ArrayStringValues(arr []string) (newArrayString []string) {
	for _, v := range arr {
		newArrayString = append(newArrayString, v)
	}
	return
}
//获取interface{}数组的所有值
func ArrayInterfaceValues(arr []interface{}) (newArrayString []interface{}) {
	for _, v := range arr {
		newArrayString = append(newArrayString, v)
	}
	return
}
//获取map[int]string的所有值
//	返回切片[]string
func MapIntStringValues(data map[int]string) (newArrayString []string) {
	for _, v := range data {
		newArrayString = append(newArrayString, v)
	}
	return
}

//获取string数组的和值
func ArrayStringSum(arr []string) int {
	arrSum := 0
	for _, v := range arr {
		vInt, _ := strconv.Atoi(v)
		arrSum += vInt
	}
	return arrSum
}
//获取string数组的和值
func ArrayIntSum(arr []int) int {
	arrSum := 0
	for _, v := range arr {
		arrSum += v
	}
	return arrSum
}

//从float64数组中获取最大值
func ArrayFloat64Max(nums []float64) float64 {
	//数组可以有1个值和多个值
	numsLen := len(nums)
	if numsLen < 1 {
		panic("ArrayFloat64Max-nums: the nums length is less than 1")
	}
	max := nums[0]
	for i := 1; i < numsLen; i++ {
		max = math.Max(max, nums[i])
	}
	return max
}

//从float64数组中获取最小值
func ArrayFloat64Min(nums []float64) float64 {
	//数组可以有1个值和多个值
	numsLen := len(nums)

	if numsLen < 1 {
		panic("ArrayFloat64Min-nums: the nums length is less than 1")
	}
	min := nums[0]
	for i := 1; i < numsLen; i++ {
		min = math.Min(min, nums[i])
	}
	return min
}

//string数组统计值出现的次数
//类似php.array_count_values
//newArrayMap[]的值是根据arr的顺序存储的
//newArrayMap[]的值map.key等于value就是arr的value值
//newArrayMap[]的值map.key等于total就是arr的value值出现的次数
//如arr := []string{7,4,5,2}
//输出效果如:[map[total:1 value:7] map[total:1 value:4] map[total:1 value:5] map[total:1 value:2]]
//如arr := []string{7,4,5,7}
//输出效果如:[map[total:2 value:7] map[total:1 value:4] map[total:1 value:5]]
func ArrayStringCountValues(arr []string) (newArrayMap []map[string]interface{}) {
	//1.先遍历所有值的出现次数
	totalMap := make(map[string]int)
	for _, v := range arr {
		tempVal, tempValExi := totalMap[v]
		if tempValExi {
			tempVal++
		} else {
			tempVal = 1
		}
		totalMap[v] = tempVal
	}

	//2.根据之前遍历所有值的出现次数,再重新按照arr的排序存储数据
	findMap := make(map[string]string)
	for _, v := range arr {
		//若不存在,则代表可以存储
		_, findExi := findMap[v]
		if !findExi {
			//这里没有判断是否存在,是因为该值在上面处理后肯定会存在的
			totalVal, _ := totalMap[v]
			//存储到最后所需要的结果中
			newArrayMap = append(newArrayMap, map[string]interface{}{
				"value": v,
				"total": totalVal,
			})
			//将已存储过的值记录,防止newArrayMap重复记录值
			findMap[v] = v
		}
	}
	return
}

//string数组比较相等
func ArrayStringCompareEq(arr1 []string, arr2 []string) bool {
	arr1Len := len(arr1)
	arr2Len := len(arr2)
	//若2个数组的数据长度不相等,则返回false
	if arr1Len != arr2Len {
		return false
	}

	//若2个数组中的某个值不相等,则返回false
	for i := 0; i < arr2Len; i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}
//int数组比较相等
func ArrayIntCompareEq(arr1 []int, arr2 []int) bool {
	arr1Len := len(arr1)
	arr2Len := len(arr2)
	//若2个数组的数据长度不相等,则返回false
	if arr1Len != arr2Len {
		return false
	}

	//若2个数组中的某个值不相等,则返回false
	for i := 0; i < arr2Len; i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}

//int数组计算交集
func ArrayIntIntersect(nums1 []int, nums2 []int) []int {

	res := make([]int, 0, len(nums1))
	nc := make(map[int]int)

	for _, n := range nums1 {
		nc[n]++
	}

	for _, n := range nums2 {
		if nc[n] > 0 {
			res = append(res, n)
			nc[n]--
		}
	}

	return res
}

//int数组计算差集(翻译对比php.array_diff)
//对比 arr1 和其他一个或者多个数组，返回在 arr1 中但是不在其他 arrs 里的值。
func ArrayIntDiff(arr1 []int, arrs ...[]int) (newArrayInt []int) {
	if len(arrs) == 0 {
		return arr1
	}
	i := 0
loop:
	for {
		if i == len(arr1) {
			break
		}
		v := arr1[i]
		for _, arr := range arrs {
			for _, val := range arr {
				if v == val {
					i++
					continue loop
				}
			}
		}
		newArrayInt = append(newArrayInt, v)
		i++
	}
	return
}
//string数组计算差集(翻译对比php.array_diff)
//对比 arr1 和其他一个或者多个数组，返回在 arr1 中但是不在其他 arrs 里的值。
func ArrayStringDiff(arr1 []string, arrs ...[]string) (newArrayString []string) {
	if len(arrs) == 0 {
		return arr1
	}
	i := 0
loop:
	for {
		if i == len(arr1) {
			break
		}
		v := arr1[i]
		for _, arr := range arrs {
			for _, val := range arr {
				if v == val {
					i++
					continue loop
				}
			}
		}
		newArrayString = append(newArrayString, v)
		i++
	}
	return
}

//string数组去重(翻译对比php.array_unique)
//	1.相同的数值只保留1个
//	2.这里要注意,php的索引数组,使用array_unique后,索引是保留的不会重置
//		示例如下:
//		$arr=[1,2,3,4,5,6,1,2,7,8,9];
//		$res=array_unique($arr);
//		print_r($res);
// 		输出:
//			Array
//			(
//				[0] => 1 [1] => 2 [2] => 3 [3] => 4 [4] => 5
// 				[5] => 6	//从这开始屏蔽掉了[6]=>1,[7]=>2
// 				[8] => 7
// 				[9] => 8 [10] => 9
//	3.注意go-ArrayStringUnique()复刻的逻辑是索引不保留,会重置
//		所以在使用ArrayStringUnique()后的数据不要再用同一个变量去重前的索引用到去重后的索引中
func ArrayStringUnique(arr []string) (newArrayString []string) {
	tempArrayMap := make(map[string]string)
	//循环遍历其值,已存在的不再存储
	for _, v := range arr {
		_, findValExi := tempArrayMap[v]
		if !findValExi {
			tempArrayMap[v] = v
			newArrayString = append(newArrayString, v)
		}
	}
	return
}
//int数组去重(翻译对比php.array_unique)
//	1.相同的数值只保留1个
//	2.其余说明同ArrayStringUnique()
func ArrayIntUnique(arr []int) (newArrayInt []int) {
	tempArrayMap := make(map[int]int)
	//循环遍历其值,已存在的不再存储
	for _, v := range arr {
		_, findValExi := tempArrayMap[v]
		if !findValExi {
			tempArrayMap[v] = v
			newArrayInt = append(newArrayInt, v)
		}
	}
	return
}

//搜索string数组中是否存在该索引
func SearchIndexByArrayString(index int, arr []string) bool {
	for k := range arr {
		if k == index {
			return true
		}
	}
	return false
}
//搜索int数组中是否存在该索引
func SearchIndexByArrayInt(index int, arr []int) bool {
	for k := range arr {
		if k == index {
			return true
		}
	}
	return false
}
//搜索float64数组中是否存在该索引
func SearchIndexByArrayFloat64(index int, arr []float64) bool {
	for k := range arr {
		if k == index {
			return true
		}
	}
	return false
}
//搜索interface数组是否存在该索引
func SearchIndexByArrayInterface(index int, arr []interface{}) bool {
	for k := range arr {
		if k == index {
			return true
		}
	}
	return false
}

//检测int数组中是否存在某个元素
func InArrayInt(ele int, arr []int) bool {
	for _, v := range arr {
		if v == ele {
			return true
		}
	}
	return false
}
//检测字符串数组中是否存在某个元素
func InArrayString(ele string, arr []string) bool {
	for _, v := range arr {
		if v == ele {
			return true
		}
	}
	return false
}
//检测字符串map中是否存在某个元素
//	该map的key是整型,value是字符串
func InMapIntString(ele string, data map[int]string) bool {
	for _, v := range data {
		if v == ele {
			return true
		}
	}
	return false
}
//检测字符串map中是否存在某个元素
//	该map的key是整型,value是int
func InMapInt(ele int, data map[int]int) bool {
	for _, v := range data {
		if v == ele {
			return true
		}
	}
	return false
}
//检测字符串map中是否存在某个元素
//	该map的key是字符串,value是字符串
func InMapString(ele string, data map[string]string) bool {
	for _, v := range data {
		if v == ele {
			return true
		}
	}
	return false
}

//int数组合并成字符串
func IntJoin(a []int, sep string) string {
	aLen := len(a)
	if aLen == 0 {
		return ""
	}
	if aLen == 1 {
		return fmt.Sprintf("%v", a[0])
	}

	buffer := &bytes.Buffer{}
	buffer.WriteString(fmt.Sprintf("%v", a[0]))
	for i := 1; i < aLen; i++ {
		buffer.WriteString(sep)
		buffer.WriteString(fmt.Sprintf("%v", a[i]))
	}
	return buffer.String()
}

// NumberFormat number_format()
// decimals: Sets the number of decimal points.
// decPoint: Sets the separator for the decimal point.
// thousandsSep: Sets the thousands separator.
func NumberFormat(number float64, decimals uint, decPoint, thousandsSep string) string {
	neg := false
	if number < 0 {
		number = -number
		neg = true
	}
	dec := int(decimals)
	// Will round off
	str := fmt.Sprintf("%."+strconv.Itoa(dec)+"F", number)
	prefix, suffix := "", ""
	if dec > 0 {
		prefix = str[:len(str)-(dec+1)]
		suffix = str[len(str)-dec:]
	} else {
		prefix = str
	}
	sep := []byte(thousandsSep)
	n, l1, l2 := 0, len(prefix), len(sep)
	// thousands sep num
	c := (l1 - 1) / 3
	tmp := make([]byte, l2*c+l1)
	pos := len(tmp) - 1
	for i := l1 - 1; i >= 0; i, n, pos = i-1, n+1, pos-1 {
		if l2 > 0 && n > 0 && n%3 == 0 {
			for j := range sep {
				tmp[pos] = sep[l2-j-1]
				pos--
			}
		}
		tmp[pos] = prefix[i]
	}
	s := string(tmp)
	if dec > 0 {
		s += decPoint + suffix
	}
	if neg {
		s = "-" + s
	}

	return s
}

//数组字符串 转 数组 整型
func ArrayStringToInt(s []string) (newArrayInt []int) {
	for _, v := range s {
		val,_:= strconv.Atoi(v)
		newArrayInt = append(newArrayInt, val)
	}
	return
}
//数组整型 转 数组字符串
func ArrayIntToString(i []int) (newArrayString []string) {
	for _, v := range i {
		val := strconv.Itoa(v)
		newArrayString = append(newArrayString, val)
	}
	return
}
//数组interface{} 转 数组字符串
func ArrayInterfaceToString(i []interface{}) (newArrayString []string) {
	for _, v := range i {
		val,_ := InterfaceToStr(v)
		newArrayString = append(newArrayString, val)
	}
	return
}
//数组浮点型64位 转 数组字符串
func ArrayFloat64ToString(f []float64) (newArrayString []string) {
	for _, v := range f {
		val,_ := JsonValToStr(v)
		newArrayString = append(newArrayString,val)
	}
	return
}

//获取字符串数组所有的key
func ArrayStringKeys(s []string) (newArrayInt []int) {
	for k := range s {
		newArrayInt = append(newArrayInt, k)
	}
	return
}
//获取字符串数组所有的key,返回切片字符串的key
func StringKeysByArrayString(s []string) (newArrayString []string) {
	for k := range s {
		newArrayString = append(newArrayString, strconv.Itoa(k))
	}
	return
}

// ArrayMerge array_merge()
func ArrayMergeString(ss ...[]string) (newArrayString []string) {
	n := 0
	for _, v := range ss {
		n += len(v)
	}
	newArrayString = make([]string, 0, n)
	for _, v := range ss {
		newArrayString = append(newArrayString, v...)
	}
	return
}
//ArrayMerge array_merge()
func ArrayMergeInt(ii ...[]int) (newArrayInt []int) {
	n := 0
	for _, v := range ii {
		n += len(v)
	}
	newArrayInt = make([]int, 0, n)
	for _, v := range ii {
		newArrayInt = append(newArrayInt, v...)
	}
	return
}
// ArrayMerge array_merge()
func ArrayMerge(ss ...[]interface{}) []interface{} {
	n := 0
	for _, v := range ss {
		n += len(v)
	}
	s := make([]interface{}, 0, n)
	for _, v := range ss {
		s = append(s, v...)
	}
	return s
}

// ArrayShift array_shift()
// Shift an element off the beginning of slice
func ArrayShift(s *[]interface{}) interface{} {
	if len(*s) == 0 {
		return nil
	}
	f := (*s)[0]
	*s = (*s)[1:]
	return f
}
//ArrayShift array_shift()
func ArrayStringShift(s *[]string) interface{} {
	if len(*s) == 0 {
		return nil
	}
	f := (*s)[0]
	*s = (*s)[1:]
	return f
}

// ArrayPop array_pop()
// Pop the element off the end of slice
func ArrayPop(s *[]interface{}) interface{} {
	if len(*s) == 0 {
		return nil
	}
	ep := len(*s) - 1
	e := (*s)[ep]
	*s = (*s)[:ep]
	return e
}
//ArrayPop array_pop()
func ArrayStringPop(s *[]string) interface{} {
	if len(*s) == 0 {
		return nil
	}
	ep := len(*s) - 1
	e := (*s)[ep]
	*s = (*s)[:ep]
	return e
}
//ArrayPop array_pop()
func ArrayIntPop(s *[]int) interface{} {
	if len(*s) == 0 {
		return nil
	}
	ep := len(*s) - 1
	e := (*s)[ep]
	*s = (*s)[:ep]
	return e
}

// ArrayPad array_pad()
func ArrayPad(s []interface{}, size int, val interface{}) []interface{} {
	if size == 0 || (size > 0 && size < len(s)) || (size < 0 && size > -len(s)) {
		return s
	}
	n := size
	if size < 0 {
		n = -size
	}
	n -= len(s)
	tmp := make([]interface{}, n)
	for i := 0; i < n; i++ {
		tmp[i] = val
	}
	if size > 0 {
		return append(s, tmp...)
	}
	return append(tmp, s...)
}
// ArrayStringPad 模拟php.array_pad()
func ArrayStringPad(s []string, size int, val string) []string {
	if size == 0 || (size > 0 && size < len(s)) || (size < 0 && size > -len(s)) {
		return s
	}
	n := size
	if size < 0 {
		n = -size
	}
	n -= len(s)
	tmp := make([]string, n)
	for i := 0; i < n; i++ {
		tmp[i] = val
	}
	if size > 0 {
		return append(s, tmp...)
	}
	return append(tmp, s...)
}

//字符串填充
// StringPad 模拟php.str_pad()
// size正数,朝右边填充;size负数,朝左边填充
func StringPad(s string, size int, val string) string {
	if size == 0 || (size > 0 && size < len(s)) || (size < 0 && size > -len(s)) {
		return s
	}
	n := size
	if size < 0 {
		n = -size
	}
	n -= len(s)
	tmp := ""
	for i := 0; i < n; i++ {
		tmp += val
	}
	if size > 0 {
		return s + tmp
	}
	return tmp + s
}

//字符串截取
//类似php.substr
func StringSubstr(str string, start int, length int) string {
	//如果指定长度为0,则返回空字符串
	if length == 0 {
		return ""
	}

	//如果起始位大于等于0 且 指定长度大于0
	if start >= 0 && length > 0 {
		strLen := len(str)
		//若起始值大于等于字符串长度
		if start >= strLen {
			return ""
		}
		end := start + length
		//若截止值大于等于字符串长度
		if end >= strLen {
			end = strLen
		}
		return str[start:end]
	}
	//如果起始位大于等于0 且 指定长度小于0
	if start >= 0 && length < 0 {
		strLen := len(str)
		//若strLen+length的值小于等于0
		if strLen+length <= 0 {
			return ""
		}
		//若start大于0
		if start > 0 {
			//若strLen+length的值小于等于strLen-start的值
			//if (strLen + length) <= (strLen - start) {
			//	return ""
			//}
			//若最终截止值小于最终起始值,则是错误的
			if strLen+length <= start {
				return ""
			}
		}
		return str[start : strLen+length]
	}

	//如果起始位小于0 且 指定长度大于0
	if start < 0 && length > 0 {
		//获取字符串长度
		strLen := len(str)
		//若start是负数,strLen+start还小于等于0
		if strLen+start <= 0 {
			if 0+length >= strLen {
				return str[0:]
			} else {
				return str[0 : 0+length]
			}
		}
		//若指定的长度大于字符串长度
		if length >= strLen {
			return str[strLen+start:]
		}
		//若最终起始值小于等于截止值,则返回最终起始值到最后的值
		if strLen+start <= strLen+start+length {
			return str[strLen+start:]
		}
		return str[strLen+start : strLen+start+length]
	}
	//如果起始位小于0 且 指定长度小于0
	if start < 0 && length < 0 {
		//获取字符串长度
		strLen := len(str)
		//若start是负数,strLen+start还小于等于0
		if strLen+start <= 0 {
			if strLen+length <= 0 {
				return ""
			} else {
				return str[0 : strLen+length]
			}
		}
		//若strLen+length小于等0
		if strLen+length <= 0 {
			return ""
		}
		//若最终截止值小于最终起始值,则是错误的
		if strLen+length <= strLen+start {
			return ""
		}
		return str[strLen+start : strLen+length]
	}

	return ""
}
//中文字符串截取(也适用于英文字符串)
//类似php.mb_substr
func MbStringSubstr(str string, start int, length int) string {
	strRune := []rune(str)
	strLen 	:= len(strRune)

	//如果指定长度为0,则返回空字符串
	if length == 0 {
		return ""
	}

	//如果起始位大于等于0 且 指定长度大于0
	if start >= 0 && length > 0 {
		//若起始值大于等于字符串长度
		if start >= strLen {
			return ""
		}
		end := start + length
		//若截止值大于等于字符串长度
		if end >= strLen {
			end = strLen
		}
		return string(strRune[start : end])
	}
	//如果起始位大于等于0 且 指定长度小于0
	if start >= 0 && length < 0 {
		//若strLen+length的值小于等于0
		if strLen+length <= 0 {
			return ""
		}
		//若start大于0
		if start > 0 {
			//若strLen+length的值小于等于strLen-start的值
			//if (strLen + length) <= (strLen - start) {
			//	return ""
			//}
			//若最终截止值小于最终起始值,则是错误的
			if strLen+length <= start {
				return ""
			}
		}
		return string(strRune[start : strLen+length])
	}

	//如果起始位小于0 且 指定长度大于0
	if start < 0 && length > 0 {
		//若start是负数,strLen+start还小于等于0
		if strLen+start <= 0 {
			if 0+length >= strLen {
				return string(strRune[0 : ])
			} else {
				return string(strRune[0 : 0+length])
			}
		}
		//若指定的长度大于字符串长度
		if length >= strLen {
			return string(strRune[strLen+start : ])
		}
		//若最终起始值小于等于截止值,则返回最终起始值到最后的值
		if strLen+start <= strLen+start+length {
			return string(strRune[strLen+start : ])
		}
		return string(strRune[strLen+start : strLen+start+length])
	}
	//如果起始位小于0 且 指定长度小于0
	if start < 0 && length < 0 {
		//若start是负数,strLen+start还小于等于0
		if strLen+start <= 0 {
			if strLen+length <= 0 {
				return ""
			} else {
				return string(strRune[0 : strLen+length])
			}
		}
		//若strLen+length小于等0
		if strLen+length <= 0 {
			return ""
		}
		//若最终截止值小于最终起始值,则是错误的
		if strLen+length <= strLen+start {
			return ""
		}
		return string(strRune[strLen+start : strLen+length])
	}

	return ""
}

//字符串首次出现的位置
//类似php.stripos()
func Stripos(haystack, needle string, offset int) int {
	length := len(haystack)
	if length == 0 || offset > length || -offset > length {
		return -1
	}

	haystack = haystack[offset:]
	if offset < 0 {
		offset += length
	}
	pos := strings.Index(strings.ToLower(haystack), strings.ToLower(needle))
	if pos == -1 {
		return -1
	}
	return pos + offset
}

//查找字符串的首次出现到结尾的字符串
//类似php.stristr()
func Stristr(haystack string, needle string) string {
	if needle == "" {
		return ""
	}
	idx := strings.Index(strings.ToLower(haystack), strings.ToLower(needle))
	if idx == -1 {
		return ""
	}
	return haystack[idx+len([]byte(needle))-1:]
}

//字符串数组截取并替换
//类似php.array_splice
//@params []string 	input		输入的数组
//@params int 		offset		数组起始位,可0、可大于0、可负数
//@params int 		length		指定的数据长度(数组起始位后拿取多少个数据)
//@params []string 	replaceData 要替换的数据(数组起始位开始进行替换的数据)
func ArrayStringSplice(input *[]string, offset int, length int, replaceData []string) []string {
	//如果有传递替换的数据,但是指定长度为0,则是异常的,要是有替换则必须指定长度
	if len(replaceData) > 0 && length == 0 {
		panic("replaceDataLen_gt_0_But_length_eq_0")
	}

	//存储截取数据
	subData := make([]string, 0)

	//数组长度
	inputLen := len(*input)

	//a.若起始位大于数组长度,则对应返回全部数组数
	//	正数起始位:subData无值,input原值	-- 模拟php.array_splice的效果
	//	负数起始位:subData原值,input无值	-- 模拟php.array_splice的效果
	if offset > 0 && offset > inputLen {
		return subData
	} else if offset < 0 && (offset*-1) > inputLen {
		subData = (*input)
		(*input) = make([]string, 0)
		return subData
	}

	//b.若没有指定长度,代表是offset起到结尾的所有部分
	if length == 0 {
		//如果 offset 为正，则从 input 数组中该值指定的偏移量开始移除
		if offset >= 0 {
			subData = (*input)[offset:]  //存储从指定位置截取到结尾所有部分
			(*input) = (*input)[:offset] //存储剩余部分
			return subData
		} else { //如果 offset 为负，则从 input 末尾倒数该值指定的偏移量开始移除。
			subData = (*input)[inputLen+offset:]
			(*input) = (*input)[:inputLen+offset]
			return subData
		}
	}

	//c.若有指定长度的处理
	patchData := make([]string, 0) //存储补位数据

	//若起始位大于等于0,则以起始位到剩余结尾位的数据处理
	//若起始位小于0,则以input末尾倒数该起始位到剩余结尾位的数据处理
	if offset >= 0 {
		subData = (*input)[offset:]
	} else {
		subData = (*input)[inputLen+offset:]
	}

	//若是大于0,则获取之前的补位数据
	if offset > 0 {
		patchData = (*input)[0:offset]
		//若是小于0,则获取0到负数之间的补位数据
	} else if offset < 0 {
		patchData = (*input)[0 : inputLen+offset]
	}

	tempSubData := make([]string, 0)
	tempInputData := make([]string, 0)
	tag := 0
	//循环处理要提取的指定长度的数据
	for _, v := range subData {
		tag++
		//在未达到指定长度数据之前,是存储到截取数据里的,达到指定长度后将剩余的都存到原数组中
		if tag <= length {
			tempSubData = append(tempSubData, v)
		} else {
			tempInputData = append(tempInputData, v)
		}
	}
	//存储指定位置和指定长度的数据
	subData = tempSubData
	//将之前补位的数据以及替换的数据添加到原数组中
	(*input) = append(patchData, replaceData...)
	//将剩余的数据添加到原数组中
	(*input) = append((*input), tempInputData...)
	return subData
}

// ArraySlice array_slice()
func ArraySlice(s []interface{}, offset, length uint) []interface{} {
	if offset > uint(len(s)) {
		panic("offset: the offset is less than the length of s")
	}
	end := offset + length
	if end < uint(len(s)) {
		return s[offset:end]
	}
	return s[offset:]
}

//模拟php的array_flip
func ArrayStringFlip(s []string) map[string]string{
	flipMapString := make(map[string]string)
	//将原字符串数组的value作为key，key作为value
	for k,v := range s {
		//转换成string数据类型存储
		flipMapString[v] = strconv.Itoa(k)
	}
	return flipMapString
}

//=====map自定义排序模板,可参照该模板逻辑,产生[],map等数据类型的自定义排序函数=====
//--------------------定义:
////map排序，按key排序
//type MapSorter []MapItem
//
//func NewMapSorter(m map[string]string) MapSorter {
//	ms := make(MapSorter, 0, len(m))
//	for k, v := range m {
//		ms = append(ms, MapItem{Key: k, Val: v})
//	}
//
//	return ms
//}
//
//type MapItem struct {
//	Key string
//	Val string
//}
//
//func (ms MapSorter) Len() int {
//	return len(ms)
//}
//
//func (ms MapSorter) Swap(i, j int) {
//	ms[i], ms[j] = ms[j], ms[i]
//}
//
////按键排序
//func (ms MapSorter) Less(i, j int) bool {
//	return ms[i].Key < ms[j].Key
//}
//--------------------调用:
//res := util.NewMapSorter(map[string]string{"z": "2","c":"3"})
//
//sort.Sort(res)
//
//fmt.Println("res:",res)
//==========================================================================

//------------------- []map[string]interface 自定义排序 start-------------------
//排序值结构体
type arrMapStringInterfaceItem struct {
	Val 			map[string]interface{}
	ValSortFunc 	func(iMap map[string]interface{},jMap map[string]interface{})bool
}
//排序数组
type arrMapStringInterfaceSorter []arrMapStringInterfaceItem
//实现sort函数接口的方法
func (ms arrMapStringInterfaceSorter) Len() int {
	return len(ms)
}
//实现sort函数接口的方法
func (ms arrMapStringInterfaceSorter) Swap(i, j int) {
	ms[i], ms[j] = ms[j], ms[i]
}
//实现sort函数接口的方法
// 按自定义函数排序
func (ms arrMapStringInterfaceSorter) Less(i, j int) bool {
	return ms[i].ValSortFunc(ms[i].Val,ms[j].Val)
}
//调用入口
//	调用示例：
//	arrMap := []map[string]interface{}{
//		{"priority":2,"id":8},
//		{"priority":2,"id":10},
//		{"priority":1,"id":1},
//	}
//
//	arrRes := util.NewArrMapStringInterfaceSorter(arrMap,func(iMap map[string]interface{},jMap map[string]interface{})bool{
//		//升序ASC 	1 < 2 返回true 	如1,2,3,4,5
//		//降序DESC 	1 > 2 返回true	如5,4,3,2,1
//		if iMap["priority"].(int)>jMap["priority"].(int) {
//			return true
//		}
//		return false
//	})
//
//	fmt.Println("arrRes:",arrRes)
func NewArrMapStringInterfaceSorter(arrMap []map[string]interface{},sortFunc func(iMap map[string]interface{},jMap map[string]interface{})bool ) []map[string]interface{} {
	//1.先将要排序的值放入到排序数组中
	arrMapSorterList := make(arrMapStringInterfaceSorter,0,len(arrMap))
	for _,v := range arrMap {
		arrMapSorterList = append(arrMapSorterList,arrMapStringInterfaceItem{Val:v,ValSortFunc:sortFunc})
	}
	//2.根据排序数组实现sort函数接口的方法进行排序处理
	sort.Sort(arrMapSorterList)

	//3.最后再将排序后的值提取出来(因为有自定义函数,不可能将该自定义函数也一并返回的)
	sortValList := make([]map[string]interface{},0,len(arrMap))
	for _,v := range arrMapSorterList {
		sortValList = append(sortValList,v.Val)
	}
	return sortValList
}
//------------------- []map[string]interface 自定义排序 end-------------------

//php的range函数
func MakeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

//生成固定范围内的整型数组
func GetRandIntArr(start int, end int) []int  {
	var intArr = []int{}
	if start > end {
		return intArr
	}
	for i:=start;i<=end;i++  {
		intArr = append(intArr, i)
	}
	return intArr
}

//@todo 此处需要引用扩展:"github.com/Pallinder/go-randomdata"
////获取从最小值~最大值之间的随机整数,如传入1,100,则最大随机数是99,最小是1
//func (thisObj *helper) RandInt64(min int,max int) (res int64) {
//	res = int64(min)
//
//	//线上报过这样的一个错误:{"filename":"/usr/local/go/src/math/rand/rng.go","funcname":"Uint64","level":"info","line":249,"msg":"异常原因: runtime error: index out of range -------math/rand.(*rngSource).Uint64(...)\n\t/usr/local/go/src/math/rand/rng.go:249...
//	//传参是10和25,目前还不知道具体原因,先在这么里这么加一个异常捕获同时以最小值返回!
//	defer func(){
//		if err := recover(); err != nil {
//			//这里应该记录异常日志,不再抛出
//			panic(err)
//		}
//	}()
//
//	//返回随机数
//	//需要引用扩展:"github.com/Pallinder/go-randomdata"
//	//res = int64(randomdata.Number(min, max+1))
//
//	return
//}
//
////获取从0-最大值的随机整数,如传入100,则最大随机数是99,最小是0
////@desc 调用示例
////	randRes := util.Helper.RandIntn(2)
////	fmt.Print(randRes)
////	输出效果如:0 1 1 2 2 0 0 0 2 1 0 1 等
//func (thisObj *helper) RandIntn(n int) (res int) {
//	res = 0
//
//	//注释同newRand.RandInt64(),若有异常则以最小值返回
//	defer func(){
//		if err := recover(); err != nil {
//			//这里应该记录异常日志,不再抛出
//			panic(err)
//		}
//	}()
//
//	//传参效果是永远比传入值小1,如传入100,最大值就是99
//	n += 1
//
//	//返回随机数
//	//需要引用扩展:"github.com/Pallinder/go-randomdata"
//	//res = randomdata.Number(0, n)
//
//	return
//}
//
////随机打乱切片数值
////@desc 调用示例
////	numRangeRes := []int{0,1,2,3,4,5,6,7,8,9}
////	numRangeRes,numRangeResInterface = util.Helper.ShuffleArrayFloat64(numRangeRes)
////	fmt.Println(numRangeRes)
////	fmt.Println(numRangeResInterface)
////	输出:
//// 		原先值：[0 1 2 3 4 5 6 7 8 9]
////		随机值：[7 3 5 2 1 8 4 9 0 6]
////		numRangeResInterface：[7 3 5 2 1 8 4 9 0 6]
//func (thisObj *helper) ShuffleArrayFloat64(arr []float64) ([]float64,[]interface{}) {
//	arrLen := len(arr)
//
//	//若小于等于0,则返回空数据
//	if arrLen<=0 {
//		return []float64{},[]interface{}{}
//	}
//
//	//存储两种结果
//	//一个是数组的float64
//	//一个是数组的interface{}
//	arrInterface := make([]interface{},arrLen)
//
//	//若大于1,则进行随机处理
//	if arrLen>1 {
//		for i := arrLen - 1; i > 0; i-- {
//			//获取随机数
//			num := thisObj.RandIntn(i)
//			//位置交换
//			arr[i], arr[num] = arr[num], arr[i]
//
//			//保持与随机数后的数组数据一致
//			arrInterface[num] 	= arr[num]
//			arrInterface[i] 	= arr[i]
//		}
//	}else{ //若等于1,则直接赋值
//		arrInterface[0] = arr[0]
//	}
//
//	return arr,arrInterface
//}

//去掉左右两边字符
func Trim(str string, characterMask ...string) string {
	mask := ""
	if len(characterMask) == 0 {
		mask = " \\t\\n\\r\\0\\x0B"
	} else {
		mask = characterMask[0]
	}
	return strings.Trim(str, mask)
}
// 去掉左边字符
func LTrim(str string, characterMask ...string) string {
	mask := ""
	if len(characterMask) == 0 {
		mask = " \\t\\n\\r\\0\\x0B"
	} else {
		mask = characterMask[0]
	}
	return strings.TrimLeft(str, mask)
}
//去掉右边字符
func RTrim(str string, characterMask ...string) string {
	mask := ""
	if len(characterMask) == 0 {
		mask = " \\t\\n\\r\\0\\x0B"
	} else {
		mask = characterMask[0]
	}
	return strings.TrimRight(str, mask)
}

//字符串反转
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

//将f1=m&f2=n解析成map[string]interface{}
//
//调用示例:
//	requestData := make(map[string]interface{})
//	parseErr := util.ParseStr("f1=m&f2=n", requestData)
//	if parseErr!=nil {
//		panic(parseErr)
//	}
//	fmt.Println("requestData:",requestData)
//
// 模拟php的parse_str():
// f1=m&f2=n 			-> map[f1:m f2:n]
// f[a]=m&f[b]=n 		-> map[f:map[a:m b:n]]
// f[a][a]=m&f[a][b]=n 	-> map[f:map[a:map[a:m b:n]]]
// f[]=m&f[]=n 			-> map[f:[m n]]
// f[a][]=m&f[a][]=n 	-> map[f:map[a:[m n]]]
// f[][]=m&f[][]=n 		-> map[f:[map[]]] // Currently does not support nested slice.
// f=m&f[a]=n 			-> error // This is not the same as PHP.
// a .[[b=c 			-> map[a___[b:c]
func ParseStr(encodedString string, result map[string]interface{}) error {
	// build nested map.
	var build func(map[string]interface{}, []string, interface{}) error

	build = func(result map[string]interface{}, keys []string, value interface{}) error {
		length := len(keys)
		// trim ',"
		key := strings.Trim(keys[0], "'\"")
		if length == 1 {
			result[key] = value
			return nil
		}

		// The end is slice. like f[], f[a][]
		if keys[1] == "" && length == 2 {
			// todo nested slice
			if key == "" {
				return nil
			}
			val, ok := result[key]
			if !ok {
				result[key] = []interface{}{value}
				return nil
			}
			children, ok := val.([]interface{})
			if !ok {
				return fmt.Errorf("expected type '[]interface{}' for key '%s', but got '%T'", key, val)
			}
			result[key] = append(children, value)
			return nil
		}

		// The end is slice + map. like f[][a]
		if keys[1] == "" && length > 2 && keys[2] != "" {
			val, ok := result[key]
			if !ok {
				result[key] = []interface{}{}
				val = result[key]
			}
			children, ok := val.([]interface{})
			if !ok {
				return fmt.Errorf("expected type '[]interface{}' for key '%s', but got '%T'", key, val)
			}
			if l := len(children); l > 0 {
				if child, ok := children[l-1].(map[string]interface{}); ok {
					if _, ok := child[keys[2]]; !ok {
						_ = build(child, keys[2:], value)
						return nil
					}
				}
			}
			child := map[string]interface{}{}
			_ = build(child, keys[2:], value)
			result[key] = append(children, child)

			return nil
		}

		// map. like f[a], f[a][b]
		val, ok := result[key]
		if !ok {
			result[key] = map[string]interface{}{}
			val = result[key]
		}
		children, ok := val.(map[string]interface{})
		if !ok {
			return fmt.Errorf("expected type 'map[string]interface{}' for key '%s', but got '%T'", key, val)
		}

		return build(children, keys[1:], value)
	}

	// split encodedString.
	parts := strings.Split(encodedString, "&")
	for _, part := range parts {
		pos := strings.Index(part, "=")
		if pos <= 0 {
			continue
		}
		key, err := url.QueryUnescape(part[:pos])
		if err != nil {
			return err
		}
		for key[0] == ' ' {
			key = key[1:]
		}
		if key == "" || key[0] == '[' {
			continue
		}
		value, err := url.QueryUnescape(part[pos+1:])
		if err != nil {
			return err
		}

		// split into multiple keys
		var keys []string
		left := 0
		for i, k := range key {
			if k == '[' && left == 0 {
				left = i
			} else if k == ']' {
				if left > 0 {
					if len(keys) == 0 {
						keys = append(keys, key[:left])
					}
					keys = append(keys, key[left+1:i])
					left = 0
					if i+1 < len(key) && key[i+1] != '[' {
						break
					}
				}
			}
		}
		if len(keys) == 0 {
			keys = append(keys, key)
		}
		// first key
		first := ""
		for i, chr := range keys[0] {
			if chr == ' ' || chr == '.' || chr == '[' {
				first += "_"
			} else {
				first += string(chr)
			}
			if chr == '[' {
				first += keys[0][i+1:]
				break
			}
		}
		keys[0] = first

		// build nested map
		if err := build(result, keys, value); err != nil {
			return err
		}
	}

	return nil
}

//类似php.urldecode()
func URLDecode(str string) (string, error) {
	return url.QueryUnescape(str)
}

//类似php.htmlentities()
func HtmlEntities(str string) string {
	return html.EscapeString(str)
}

//类似php.html_entity_decode()
func HtmlEntityDecode(str string) string {
	return html.UnescapeString(str)
}

// 类似php.is_numeric()
// Numeric strings consist of optional sign, any number of digits, optional decimal part and optional exponential part.
// Thus +0123.45e6 is a valid numeric value.
// In PHP hexadecimal (e.g. 0xf4c3b00c) is not supported, but IsNumeric is supported.
func IsNumeric(val interface{}) bool {
	switch val.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return true
	case float32, float64, complex64, complex128:
		return true
	case string:
		str := val.(string)
		if str == "" {
			return false
		}
		// Trim any whitespace
		str = strings.TrimSpace(str)
		if str[0] == '-' || str[0] == '+' {
			if len(str) == 1 {
				return false
			}
			str = str[1:]
		}
		// hex
		if len(str) > 2 && str[0] == '0' && (str[1] == 'x' || str[1] == 'X') {
			for _, h := range str[2:] {
				if !((h >= '0' && h <= '9') || (h >= 'a' && h <= 'f') || (h >= 'A' && h <= 'F')) {
					return false
				}
			}
			return true
		}
		// 0-9, Point, Scientific
		p, s, l := 0, 0, len(str)
		for i, v := range str {
			if v == '.' { // Point
				if p > 0 || s > 0 || i+1 == l {
					return false
				}
				p = i
			} else if v == 'e' || v == 'E' { // Scientific
				if i == 0 || s > 0 || i+1 == l {
					return false
				}
				s = i
			} else if v < '0' || v > '9' {
				return false
			}
		}
		return true
	}

	return false
}
