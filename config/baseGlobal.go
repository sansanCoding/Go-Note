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

//@todo 参考Iris-Go注册数据库链接与注册Redis链接（但凡数据库或Redis链接创建失败,则直接退出程序不再执行）
//@todo 必须要有如下包引入才可使用
//import (
//	"database/sql"
//	"github.com/go-redis/redis"
//	_ "github.com/go-sql-driver/mysql"
//)

////redis对象容器，需要用到redis直接到这里获取
//var RedisClientMap = make(map[string]*redis.Client)
//
////mysql对象容器，需要用到mysql直接到这里获取
//var MysqlClientMap = make(map[string]*sql.DB)
//
////mysql对象容器(后台admin库)，需要用到mysql直接到这里获取
//var MysqlAdminClientMap = make(map[string]*sql.DB)

////外部直接调用-1.创建mysql与redis的链接池
////	如main.go文件中写入config.CreateRedisAndMysqlConnectionPool()
//func CreateRedisAndMysqlConnectionPool() {
//	//循环注册mysql和redis
//	for key, value := range DBS {
//		//注册mysql链接
//		mysqlConf := value.(map[string]interface{})["mysql"].(map[string]DBSet)
//		registerDataBaseAction(key+"_read", mysqlConf, false)
//		registerDataBaseAction(key+"_write", mysqlConf, false)
//		registerDataBaseAction(key+"_admin", mysqlConf, true)
//
//		//注册redis链接
//		redisConf := value.(map[string]interface{})["redis"].(RedisSet)
//		registerSingleRedis(key, redisConf))
//	}
//
//	//可以在额外注册其他mysql数据库
//	//registerDataBaseAction("testDB1", mysqlConf1, false)
//	//registerDataBaseAction("testDB2", mysqlConf2, false)
//
//	//可以在额外注册其他redis
//	//registerSingleRedis("testRedis1", redisConf1))
//	//registerSingleRedis("testRedis2", redisConf2))
//
//	fmt.Println("mysql 与 redis 连接池创建完毕.........")
//}

////注册数据库连接
//func registerDataBaseAction(key string,dbConf map[string]DBSet,isAdmin bool) {
//	var dbSet DBSet
//	dbSet = dbConf[key]
//	db,err := sql.Open(dbSet.DriveName, dbSet.DriveDsn)
//	if err != nil {
//		fmt.Println("数据库链接失败,排查数据库配置是否正确以及数据库是否正常运行!")
//		fmt.Println("sqlOpenErr:"+err.Error())
//		fmt.Println("程序已停止运行!")
//		os.Exit(1)
//	}
//	//使用ping检查数据库是否链接成功
//	pingErr = db.Ping()
//	if pingErr != nil {
//		fmt.Println(key + "-数据库链接池创建失败,排查数据库配置是否正确以及数据库是否正常运行!")
//		fmt.Println("dbPingErr:"+pingErr.Error())
//		fmt.Println("程序已停止运行!")
//		os.Exit(1)
//	}
//	//最大空闲数
//	db.SetMaxIdleConns(dbSet.MaxIdle)
//	//最大链接数
//	db.SetMaxOpenConns(dbSet.MaxConn)
//	//链接生命周期
//	db.SetConnMaxLifetime(7 * time.Second)
//	//如果是后台数据库
//	if isAdmin {
//		//保存对应环境的后台admin库链接池对象
//		MysqlAdminClientMap[key] = db
//	} else {
//		//保存对应环境的链接池对象
//		MysqlClientMap[key] = db
//	}
//}

////注册redis链接
//func registerSingleRedis(key string,options RedisSet) {
//	client := redis.NewClient(&redis.Options{
//		Addr:         options.Addr,
//		Password:     options.Password,
//		DB:           options.DB,
//		MaxRetries:   options.MaxRetries,
//		PoolSize:     options.PoolSize,
//		MinIdleConns: options.MinIdleConn,
//	})
//	//通过client.Ping()来检查是否成功连接到了redis服务器
//	_, err := client.Ping().Result()
//	if err != nil {
//		fmt.Println(key + "-Redis链接池创建失败,排查Redis配置是否正确以及Redis是否正常运行!")
//		fmt.Println("pingErr:"+err.Error())
//		fmt.Println("程序已停止运行!")
//		os.Exit(1)
//	}
//	//保存对应环境的链接池对象
//	RedisClientMap[key] = client
//}
//------------------------------------------ 数据库与redis-配置数据与注册链接函数 end ------------------------------------------