package MySQL_Redis

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
////@todo 如main.go文件中写入CreateRedisAndMysqlConnectionPool(),相当于提前创建好链接池,后面直接用,无需再重复创建
//func CreateRedisAndMysqlConnectionPool() {
//	//循环注册mysql和redis
//	for key, value := range config.DBS {
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
























