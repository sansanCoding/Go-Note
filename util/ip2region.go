package util
//
////******************************** IPV4-IP地址解析 ********************************
////lionsoul2014/ip2region-golang(也有对应的php版本,java版本等)使用说明地址:
////https://github.com/lionsoul2014/ip2region/tree/master/binding/golang
////********************************************************************************
//
//import (
//	"errors"
//	//@todo 这里是基于Iris-Go框架写的代码,可供参考使用!
//	//@todo 需要引用如下扩展包
//	"github.com/lionsoul2014/ip2region/binding/golang/ip2region"
//)
//
////消息前缀
//var ip2regionLibraryMsgPrefix = "util_ip2regionLibrary-"
//
////ip2regionDB文件路径
////注意:这里的文件地址,是以项目目录下开始计算的,即./项目目录/util/ip2region.db 文件路径为 xxx/src/项目目录/util/ip2region.db
//var ip2regionDBFilePath = "./util/ip2region.db"
//
////ip2regionLibrary结构体
//type ip2regionLibrary struct {
//	ip2regionObj *ip2region.Ip2Region
//	ip2regionObjErr error
//}
//
////创建对象
//func NewIp2regionLibrary() *ip2regionLibrary {
//	obj := new(ip2regionLibrary)
//
//	//创建ip地址解析资源
//	createError := obj.createIpAddressAnalyze()
//	if createError!=nil {
//		panic(createError.Error())
//	}
//
//	return obj
//}
//
////创建ip地址解析资源
//func (cthis *ip2regionLibrary) createIpAddressAnalyze() error {
//	if cthis.ip2regionObj==nil {
//		//创建ip地址解析资源
//		cthis.ip2regionObj,cthis.ip2regionObjErr = ip2region.New(ip2regionDBFilePath)
//		if cthis.ip2regionObjErr!=nil {
//			return errors.New(ip2regionLibraryMsgPrefix+"ip2regionNewErr:"+cthis.ip2regionObjErr.Error())
//		}
//	}
//	return nil
//}
//
////*** 外部调用入口 ***
////执行ip地址解析
//func (cthis *ip2regionLibrary) DoIpAddressAnalyze(ipStr string) (actionResult map[string]interface{},actionError error) {
//	actionResult = map[string]interface{}{
//
//	}
//	actionError = nil
//
//	//解析ip地址
//	ip, err := cthis.ip2regionObj.MemorySearch(ipStr)
//	if err!=nil {
//		actionError = errors.New(ip2regionLibraryMsgPrefix+"ip2regionMemorySearchErr:"+err.Error())
//		return actionResult,actionError
//	}
//
//	//调试输出
//	//fmt.Println(ip, err)
//	//fmt.Println(ip.Country)
//	//fmt.Println(ip.Province)
//	//fmt.Println(ip.City)
//	//输出结果:
//	//ipStr := "58.97.161.14"
//	//0|孟加拉|0|0|0|0 <nil>
//	//	ip.Country 孟加拉
//	//	ip.Province 0
//	//	ip.City 0
//	//ipStr := "42.90.126.71"
//	//3110|中国|0|甘肃省|兰州市|电信 <nil>
//	//	ip.Country 中国
//	//	ip.Province 甘肃省
//	//	ip.City 兰州市
//
//	//操作成功存储
//	//城市Id
//	actionResult["CityId"] = ip.CityId
//	//国家
//	actionResult["Country"] = ip.Country
//	//地区
//	actionResult["Region"] = ip.Region
//	//省
//	actionResult["Province"] = ip.Province
//	//市
//	actionResult["City"] = ip.City
//	//ISP
//	actionResult["ISP"] = ip.ISP
//
//	return actionResult,actionError
//}
//
////关闭ip地址解析资源
//func (cthis *ip2regionLibrary) Close(){
//	if cthis.ip2regionObj!=nil {
//		//关闭资源
//		cthis.ip2regionObj.Close()
//		//同时将成员属性置为null,下次再重新创建资源对象
//		cthis.ip2regionObj = nil
//	}
//}
