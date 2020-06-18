package config

import (
	"os"
)


//基础全局配置,如涉及sql数据库地址链接配置、redis链接配置等都配置在此


//是否是线上环境:true 是线上环境,false 不是线上环境
var IsRelease = false

func init(){

	//根据线上系统配置获取是否是线上环境
	if os.Getenv("MODE")=="release" {
		IsRelease = true
	}

}

//------------------------------------------ 数据库与redis-配置数据与注册链接函数 start ------------------------------------------
//Redis配置
type RedisSet struct {
	Addr         string //host:port
	Password     string //数据库密码
	DB           int    //选择的数据库
	MaxRetries   int    //重试最大次数
	PoolSize     int    //链接池大小
	MinIdleConn  int    //最小空闲数
}

//数据库配置
type DBSet struct {
	AliasName string //数据库别名
	DriveName string //驱动名称
	DriveDsn  string //数据库链接地址
	MaxIdle   int    //最大空闲数
	MaxConn   int    //最大链接数
}

//数据库与Redis配置集(当然可以配置多个与数据库类似的配置,如MongoDB之类的)
//@todo 参考Iris-Go框架的数据库配置结构构造,有可能不适用于其他框架,可自行根据需求进行改造
//@todo 这里可以配置多个环境的数据库
var DBS = map[string]interface{}{
	//本地环境测试数据库
	"test":map[string]interface{}{
		//描述
		"desc":"测试环境",
		//排序
		"order":1,
		//mysql链接配置
		"mysql":map[string]DBSet{
			//读库
			"test_read":{
				AliasName:"test_read",
				DriveName:"mysql",
				//@todo 如果是按如下这么配置的数据库链接地址、账号、密码等，则在字符串不可出现特殊有占位意义的特殊字符:
				//@todo 如 : @ 就是对应的分割符,不可重复出现
				//@todo 如下格式讲解:
				//@todo 账号:密码@tcp(链接地址,ip:端口 或 域名:端口)/数据库的库名?charset=utf8&multiStatements=true
				//密码中不管是否有#,都不影响数据库创建链接
				//DriveDsn:"root:123#456@tcp(192.168.0.225:3306)/test?charset=utf8&multiStatements=true",
				DriveDsn:"root:123456@tcp(192.168.0.225:3306)/test?charset=utf8&multiStatements=true",
				MaxIdle:1,
				MaxConn:5,
			},
			//写库(maxAllowedPacket针对一个packet数据包[该数据包的内容如insert语句insert xxTable(xxField...) values(...),(...)]的大小限制,值如1024 * 1024 *160 = 160M ,1024 * 1024 * 1024 = 1G)
			"test_write":{
				AliasName:"test_write",
				DriveName:"mysql",
				DriveDsn:"root:123456@tcp(192.168.0.225:3306)/test?charset=utf8&multiStatements=true&maxAllowedPacket=104857600",
				MaxIdle:1,
				MaxConn:5,
			},
			//admin后台库(读写库)
			"test_admin":{
				AliasName:"test_admin",
				DriveName:"mysql",
				DriveDsn:"root:123456@tcp(192.168.0.225:3306)/test_admin?charset=utf8&multiStatements=true",
				MaxIdle:1,
				MaxConn:5,
			},
		},
		//redis链接配置
		"redis":RedisSet{
			Addr:"192.168.0.225:6379",
			Password:"abcdfg@123456",
			DB:0,
			MaxRetries:2,
			PoolSize:10,
			MinIdleConn:5,
		},
		//webSocket链接配置
		"webSocket": "wss://xxx.com/websocket/test",
	},
}