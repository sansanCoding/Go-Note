package util
//
//import (
//	"bufio"
//	"database/sql"
//	"encoding/json"
//	"errors"
//	"fmt"
//	"io"
//	"os"
//	"strings"
//)
//
////消息前缀
//var ipv6regionLibraryMsgPrefix = "util_ipv6regionLibrary-"
//
////php导出ipv6数据的json串文件路径
//var phpIpv6ToRegionJsonFilePath = "./util/phpIpv6ToRegionJson.log"
//
////ipv6ToRegionDB文件路径
////注意:实际文件路径地址是/xxx/xxx/src/项目目录/util/ipv6ToRegion.db
//var ipv6ToRegionDBFilePath = "./util/ipv6ToRegion.db"
//
//type ipv6regionLibrary struct {
//	sqLiteDB *sql.DB
//	sqLiteDBErr error
//}
//
////创建对象
//func NewIpv6regionLibrary(params map[string]interface{}) *ipv6regionLibrary {
//
//	//是否需要创建表操作,true 需要,false 不需要
//	isCreateTable := false
//	if params["isCreateTable"]!=nil {
//		isCreateTable = params["isCreateTable"].(bool)
//	}
//
//	obj := new(ipv6regionLibrary)
//
//	var err error
//	//初始化打开db
//	err = obj.openDB()
//	if err!=nil {
//		panic(err.Error())
//	}
//
//	//如果需要创建表操作
//	if isCreateTable {
//		//创建表操作
//		err = obj.createTable()
//		if err!=nil {
//			panic(err.Error())
//		}
//	}
//
//	return obj
//}
//
////*** 外部调用入口 ***
////查询ipv6-ip解析地址
//func (cthis *ipv6regionLibrary) GetIPv6Address(params map[string]interface{}) (actionResult map[string]interface{},actionError error){
//
//	//ipv6-ip地址,如240e:82:f000:adef:b9a3:c1eb:d555:badd
//	ipv6 := params["ipv6"].(string)
//
//	//查询一条数据
//	return cthis.selectDataByOne(map[string]interface{}{
//		"ipv6":ipv6,
//	})
//}
//
////关闭资源
//func (cthis *ipv6regionLibrary) Close(){
//	if cthis.sqLiteDB!=nil {
//		cthis.sqLiteDB.Close()
//	}
//}
//
////---------------------------- IPv6 以文件形式插入sqLiteDB数据操作 start ----------------------------
////调用示例:
////			////1.若有需要新增ipv6的数据操作,打开注释,本地执行!
////			////创建对象
////			//obj:= util.NewIpv6regionLibrary(map[string]interface{}{
////			//	"isCreateTable":true,
////			//})
////			////最后才关闭资源
////			//defer obj.Close()
////			////IPv6 以文件形式插入sqLiteDB数据操作
////			//// a.具体文件内容形式参考obj.InsertSqLiteDBDataByFile()方法里的操作流程.
////			//// b.该操作只需要本地执行后,将ipv6ToRegion.db上传到git管理,再更新到线上环境即可,记得重启go服务!
////			//insertSqLiteDBDataByFileRes,insertSqLiteDBDataByFileResErr := obj.InsertSqLiteDBDataByFile(map[string]interface{}{})
////			//fmt.Println("insertSqLiteDBDataByFileRes:",insertSqLiteDBDataByFileRes)
////			//fmt.Println("insertSqLiteDBDataByFileResErr:",insertSqLiteDBDataByFileResErr)
////			//
////			////调试操作
////			////obj.Test()
////
////以文件形式获取数据插入sqLiteDB
////	1.phpIpv6ToRegionJson.log该文件每一行都是json字符串!!!
////		json格式如:
////		{"ipv6":"240e:82:f000:adef:b9a3:c1eb:d555:badd","address":"\u4e2d\u56fd\u5317\u4eac\u5e02 \u4e2d\u56fd\u7535\u4fe1CTNET\u7f51\u7edc"}
////		翻译到go就是(或者go反解析json后的数据结构):
////		map[string]interface{}{
////			"ipv6":"240e:82:f000:adef:b9a3:c1eb:d555:badd",
////			"address":"中国北京市 中国电信CTNET网络",
//// 		}
////	2.重复的数据会覆盖处理!!!
////	3.该方法循环处理文件内容时，每次循环以fmt.Println()输出统计数据,sql等告知.
////	4.若以后想添加数据,只需要本地执行phpIpv6ToRegionJson.log这个文件操作将文件内容的数据插入到sqLiteDB文件中即可,然后将ipv6ToRegion.db这样的sqLiteDB文件上传更新到线上!
////	  最后别忘了将 phpIpv6ToRegionJson.log 文件内容置为空白的,即一个空文件(该文件只作添加sqLiteDB数据用,用完就清空)!
//func (cthis *ipv6regionLibrary) InsertSqLiteDBDataByFile(params map[string]interface{}) (actionResult map[string]interface{},actionError error){
//	actionResult = map[string]interface{}{
//
//	}
//	actionError = nil
//
//	//该文件名是自定义的,可根据需求自行更改
//	//	1.实际文件路径地址是/xxx/xxx/src/项目目录/util/phpIpv6ToRegionJson.log
//	//	2.但凡需要添加数据,只要将phpIpv6ToRegionJson.log该文件内容写入需要添加的数据即可!!!
//	filePath := phpIpv6ToRegionJsonFilePath
//
//	//打开文件
//	file,fileErr := os.Open(filePath)
//	if fileErr!=nil {
//		actionError = errors.New(ipv6regionLibraryMsgPrefix+"insertSqLiteDBDataByFile_Open_error:"+fileErr.Error())
//		return actionResult,actionError
//	}
//	//执行完毕后关闭资源
//	defer file.Close()
//
//	//创建缓冲区
//	fileReader := bufio.NewReader(file)
//
//	//总页数(即走了多少页,也可以看做执行了多少次sql操作)
//	pageTotal := 0
//	//每页最大数据量(即每页处理数 达到 每页最大数据量,则执行一次sql插入操作)
//	pageMax := 100
//	//每页处理数
//	pageSum := 0
//	//插入数据计数
//	insertTotal := 0
//
//	//准备插入sql
//	// ipv6是唯一键索引,若数据重复,则会使用REPLACE进行替换操作【该REPLACE操作会删除原数据再按新数据插入,id也是按新增id累计,即原id是10,REPLACE操作后同样的数据不变,但是id会变成11】
//	insertField 	:= "INSERT OR REPLACE INTO "+cthis.tableName()+" (`ipv6`,`address`) VALUES"
//	//准备插入sql值
//	insertValArr 	:= make([]string,0)
//	//循环读取每一行数据
//	for {
//		//该行字节的数据
//		lineBytes,_,err := fileReader.ReadLine()
//		//如果已没有可读取的数据,则停止循环
//		if err==io.EOF{
//			break
//		}
//
//		//将字节转成字符串
//		//lineStr := string(lineBytes)
//		//调试输出:
//		//fmt.Println("lineStr:",lineStr)
//		//输出结果:
//		//lineStr: {"ipv6":"240e:82:f000:adef:b9a3:c1eb:d555:badd","address":"\u4e2d\u56fd\u5317\u4eac\u5e02 \u4e2d\u56fd\u7535\u4fe1CTNET\u7f51\u7edc"}
//
//		//每行json数据反解析
//		var lineData map[string]interface{}
//		jsonErr := json.Unmarshal(lineBytes,&lineData)
//		if jsonErr!=nil {
//			actionError = errors.New(ipv6regionLibraryMsgPrefix+"fileReader_for_lineDataToJsonUnmarshalErr:"+jsonErr.Error())
//			return actionResult,actionError
//		}
//		//ipv6-ip地址
//		ipv6,_ 		:= JsonValToStr(lineData["ipv6"])
//		//ipv6-ip解析地址
//		address,_ 	:= JsonValToStr(lineData["address"])
//		//拼接效果如:("ff00::ffff:ffff:ffff:ffef","IETF 多播地址")
//		insertValArr = append(insertValArr,"(\""+ipv6+"\",\""+address+"\")")
//
//		//每页处理数计数
//		pageSum++
//		//插入数据总计
//		insertTotal++
//
//
//		//循环输出-统计数据:
//		fmt.Println(map[string]interface{}{
//			"___0.pageSum___":pageSum,
//			"___1.pageTotal___":pageTotal,
//			"___2.insertTotal___":insertTotal,
//		})
//
//
//		//如果达到指定的数据上限,则进行sql操作
//		if pageSum>=pageMax {
//			//1.进行sql操作
//			//拼接sql
//			insertSql := insertField+strings.Join(insertValArr,",")
//
//
//			//循环输出-执行sql:
//			fmt.Println("___insertSql___:",insertSql)
//
//
//			//执行sql
//			_,execErr := cthis.execSql(map[string]interface{}{
//				"sqlStr":insertSql,
//			})
//			//若产生错误,则直接返回操作结果,不再继续往下执行(如果改成break,则insertValArr还是有值时,最后还会执行一次sql操作)
//			if execErr!=nil {
//				actionError = errors.New(ipv6regionLibraryMsgPrefix+"fileReader_for_insertSql_execErr:"+execErr.Error())
//				return actionResult,actionError
//			}
//
//			//2.进行重置
//			//重置插入值
//			insertValArr = make([]string,0)
//			//重置后重新计数
//			pageSum = 0
//			//总页数计数
//			pageTotal++
//		}//if end!
//
//	}//for end!
//
//	//若循环结束,最后还有剩余插入的值,则再进行扫尾数据的sql操作
//	if len(insertValArr)>0 {
//		//拼接sql
//		lastInsertSql := insertField+strings.Join(insertValArr,",")
//
//
//		//执行sql输出
//		fmt.Println("___lastInsertSql___:",lastInsertSql)
//
//
//		//执行sql
//		_,execErr := cthis.execSql(map[string]interface{}{
//			"sqlStr":lastInsertSql,
//		})
//		//若产生错误,则直接返回操作结果
//		if execErr!=nil {
//			actionError = errors.New(ipv6regionLibraryMsgPrefix+"fileReader_for_insertSql_lastExecErr:"+execErr.Error())
//			return actionResult,actionError
//		}
//
//		//总页数计数
//		pageTotal++
//	}
//
//	//最终处理结束后的重置(其实也可以不用再重置,因为也没什么使用地方了)
//	pageSum = 0
//	insertValArr = make([]string,0)
//
//	//操作成功存储
//	actionResult["pageTotal"] = pageTotal
//	actionResult["insertTotal"] = insertTotal
//
//	return actionResult,actionError
//}
////---------------------------- IPv6 以文件形式插入sqLiteDB数据操作 end ----------------------------
//
////调试操作
//func (cthis *ipv6regionLibrary) Test(){
//
//	//1.删除表操作,删除表操作只能单独执行
//	//dropTableErr := cthis.dropTable()
//	//fmt.Println("dropTableErr:")
//	//fmt.Println(dropTableErr)
//
//	////2.增删改查操作,不能在删除表操作之后执行
//	////先查询数据存在不在
//	//oneRes,oneResErr := cthis.selectDataByOne(map[string]interface{}{
//	//	"ipv6":"240e:82:f000:adef:b9a3:c1eb:d555:badd",
//	//})
//	//if oneResErr!=nil {
//	//	panic("selectError:"+oneResErr.Error())
//	//}
//	////数据存在则输出,否则就插入一条
//	//if len(oneRes)>0 {
//	//	fmt.Println("oneRes:",oneRes)
//	//}else{
//	//	lastId,err := cthis.insertData(map[string]interface{}{
//	//		"insertData":map[string]interface{}{
//	//			"ipv6":"240e:82:f000:adef:b9a3:c1eb:d555:badd",
//	//			"address":"中国北京市 中国电信CTNET网络",
//	//		},
//	//	})
//	//	if err!=nil {
//	//		panic("insertError:"+err.Error())
//	//	}
//	//	fmt.Println("lastId:",lastId)
//	//}
//	//
//	////删除数据
//	//deleteRes,err := cthis.deleteData(map[string]interface{}{
//	//	"ipv6":"240e:82:f000:adef:b9a3:c1eb:d555:badd",
//	//})
//	//if err!=nil {
//	//	panic("deleteError:"+err.Error())
//	//}
//	//fmt.Println("deleteRes:",deleteRes)
//
//	////修改数据
//	//updateDataRes,updateDataResErr := cthis.updateData(map[string]interface{}{
//	//	"updateData":map[string]interface{}{"address":"test"},
//	//	"updateWhere":map[string]interface{}{"ipv6":"240e:82:f000:adef:b9a3:c1eb:d555:badd",},
//	//})
//	//if updateDataResErr!=nil {
//	//	panic("updateError:"+updateDataResErr.Error())
//	//}
//	//
//	//fmt.Println("updateDataRes:",updateDataRes)
//
//}
//
////============================== 以下是私有方法相关操作集合 ==============================
//
////---------------------------- IPv6 sqLiteDB操作 start ----------------------------
////打开数据库
//func (cthis *ipv6regionLibrary) openDB() error {
//	if cthis.sqLiteDB==nil {
//		dbFileName := ipv6ToRegionDBFilePath
//		cthis.sqLiteDB, cthis.sqLiteDBErr = sql.Open("sqlite3", dbFileName)
//		if cthis.sqLiteDBErr!=nil {
//			return errors.New(ipv6regionLibraryMsgPrefix+"sqlOpen_dbFile["+dbFileName+"]_is_failed_error:"+cthis.sqLiteDBErr.Error())
//		}
//	}
//	return nil
//}
//
////统一表名
//func (cthis *ipv6regionLibrary) tableName() string {
//	return "ipv6Address"
//}
//
////创建表
//func (cthis *ipv6regionLibrary) createTable() error {
//	sqlStr := " CREATE TABLE IF NOT EXISTS "+cthis.tableName()+" ( " +
//			  "		`id` 		INTEGER		 PRIMARY KEY AUTOINCREMENT, " +
//			  "		`ipv6` 		VARCHAR(500) UNIQUE 		NOT NULL, " +	//ipv6的ip地址,值如240e:82:f000:adef:b9a3:c1eb:d555:badd
//			  " 	`address` 	TEXT						NOT NULL " +	//ipv6的解析后地址,值如中国北京市 中国电信CTNET网络
//		      " ) "
//	_,err := cthis.sqLiteDB.Exec(sqlStr)
//	if err!=nil {
//		return errors.New(ipv6regionLibraryMsgPrefix+"createTable_Exec_error:"+err.Error())
//	}
//
//	return nil
//}
//
////删除表
//func (cthis *ipv6regionLibrary) dropTable() error {
//	sqlStr := "DROP TABLE IF EXISTS "+cthis.tableName()
//	_,err := cthis.sqLiteDB.Exec(sqlStr)
//	if err!=nil {
//		return errors.New(ipv6regionLibraryMsgPrefix+"dropTable_Exec_error:"+err.Error())
//	}
//
//	return nil
//}
//
////执行sql操作-增删改
//func (cthis *ipv6regionLibrary) execSql(params map[string]interface{}) (actionResult map[string]interface{},actionError error){
//	actionResult = map[string]interface{}{
//	}
//	actionError = nil
//
//	//必传-执行sql
//	sqlStr 		:= params["sqlStr"].(string)
//
//	//预解析sql
//	stmt,err := cthis.sqLiteDB.Prepare(sqlStr)
//	if err!=nil {
//		actionError = errors.New(ipv6regionLibraryMsgPrefix+"execSql_Prepare_error:"+err.Error())
//		return actionResult,actionError
//	}
//
//	//最主要释放的是锁
//	defer stmt.Close()
//
//	//执行sql
//	_,execErr := stmt.Exec()
//	if execErr!=nil {
//		actionError = errors.New(ipv6regionLibraryMsgPrefix+"execSql_Exec_error:"+execErr.Error())
//		return actionResult,actionError
//	}
//
//	return actionResult,actionError
//}
//
////插入数据
//func (cthis *ipv6regionLibrary) insertData(params map[string]interface{}) (actionResult map[string]interface{},actionError error){
//	actionResult = map[string]interface{}{
//	}
//	actionError = nil
//
//	//必传-插入数据
//	insertData := params["insertData"].(map[string]interface{})
//
//	//准备相关插入数据
//	fieldArr 		:= make([]string,0)
//	valArr 			:= make([]interface{},0)
//	valPrepareArr 	:= make([]string,0)
//	for field,val := range insertData {
//		fieldArr 		= append(fieldArr,field)
//		valArr 			= append(valArr,val)
//		valPrepareArr 	= append(valPrepareArr,"?")
//	}
//	fieldStr 		:= strings.Join(fieldArr,",")
//	valPrepareStr 	:= strings.Join(valPrepareArr,",")
//
//	//预解析sql
//	stmt,err := cthis.sqLiteDB.Prepare("INSERT INTO "+cthis.tableName()+"("+fieldStr+") VALUES("+valPrepareStr+")")
//	if err!=nil {
//		actionError = errors.New(ipv6regionLibraryMsgPrefix+"insertData_Prepare_error:"+err.Error())
//		return actionResult,actionError
//	}
//
//	//最主要释放的是锁
//	defer stmt.Close()
//
//	//执行sql
//	res,execErr := stmt.Exec(valArr...)
//	if execErr!=nil {
//		actionError = errors.New(ipv6regionLibraryMsgPrefix+"insertData_Exec_error:"+execErr.Error())
//		return actionResult,actionError
//	}
//
//	//获取最后插入id
//	lastId,lastIdErr := res.LastInsertId()
//	if lastIdErr!=nil {
//		actionError = errors.New(ipv6regionLibraryMsgPrefix+"insertData_LastInsertId_error:"+lastIdErr.Error())
//		return actionResult,actionError
//	}
//
//	//操作成功存储
//	actionResult["lastId"] = lastId
//
//	return actionResult,actionError
//}
//
////修改数据
//func (cthis *ipv6regionLibrary) updateData(params map[string]interface{}) (actionResult map[string]interface{},actionError error){
//	actionResult = map[string]interface{}{
//	}
//	actionError = nil
//
//	//必传-修改数据
//	updateData 	:= params["updateData"].(map[string]interface{})
//	//可选-修改条件
//	updateWhere := params["updateWhere"].(map[string]interface{})
//
//	//-------------------- 准备相关修改数据 start --------------------
//	fieldArr 		:= make([]string,0)
//	valArr 			:= make([]interface{},0)
//	for field,val := range updateData {
//		fieldArr 		= append(fieldArr,field+"=?")
//		valArr 			= append(valArr,val)
//	}
//	fieldStr 		:= strings.Join(fieldArr,",")
//	//-------------------- 准备相关修改数据 end --------------------
//
//	//-------------------- 准备相关修改条件 start --------------------
//	//sql条件组
//	sqlWhereArr := make([]string,0)
//
//	//可选-ipv6的ip地址
//	if updateWhere["ipv6"]!=nil {
//		sqlWhereArr = append(sqlWhereArr," `ipv6` = '"+updateWhere["ipv6"].(string)+"' ")
//	}
//
//	//sql条件字符串
//	sqlWhere := strings.Join(sqlWhereArr," AND ")
//
//	//若有条件存在,则根据条件删除
//	if sqlWhere!="" {
//		sqlWhere = " WHERE "+sqlWhere
//	}
//	//-------------------- 准备相关修改条件 end --------------------
//
//	//预解析sql
//	stmt,err := cthis.sqLiteDB.Prepare("UPDATE "+cthis.tableName()+" SET "+fieldStr+sqlWhere)
//	if err!=nil {
//		actionError = errors.New(ipv6regionLibraryMsgPrefix+"updateData_Prepare_error:"+err.Error())
//		return actionResult,actionError
//	}
//
//	//最主要释放的是锁
//	defer stmt.Close()
//
//	//执行sql
//	_,execErr := stmt.Exec(valArr...)
//	if execErr!=nil {
//		actionError = errors.New(ipv6regionLibraryMsgPrefix+"updateData_Exec_error:"+execErr.Error())
//		return actionResult,actionError
//	}
//
//	//修改影响行就不判断,有可能是重复数据
//
//	return actionResult,actionError
//}
//
////删除数据
//func (cthis *ipv6regionLibrary) deleteData(params map[string]interface{}) (actionResult map[string]interface{},actionError error){
//	actionResult = map[string]interface{}{
//	}
//	actionError = nil
//
//	//sql条件组
//	sqlWhereArr := make([]string,0)
//
//	//可选-ipv6的ip地址
//	if params["ipv6"]!=nil {
//		sqlWhereArr = append(sqlWhereArr," `ipv6` = '"+params["ipv6"].(string)+"' ")
//	}
//
//	//sql条件字符串
//	sqlWhere := strings.Join(sqlWhereArr," AND ")
//
//	//若有条件存在,则根据条件删除
//	if sqlWhere!="" {
//		sqlWhere = " WHERE "+sqlWhere
//	}
//
//	//删除sql
//	sqlStr := "DELETE FROM "+cthis.tableName()+sqlWhere
//
//	//预解析sql
//	stmt,err := cthis.sqLiteDB.Prepare(sqlStr)
//	if err!=nil {
//		actionError = errors.New(ipv6regionLibraryMsgPrefix+"deleteData_Prepare_error:"+err.Error())
//		return actionResult,actionError
//	}
//
//	//执行sql
//	_,execErr := stmt.Exec()
//	if execErr!=nil {
//		actionError = errors.New(ipv6regionLibraryMsgPrefix+"deleteData_Exec_error:"+execErr.Error())
//		return actionResult,actionError
//	}
//
//	//最主要释放的是锁
//	defer stmt.Close()
//
//	return actionResult,actionError
//}
//
////统计该表的数据总条数
//func (cthis *ipv6regionLibrary) countTableDataTotal(params map[string]interface{}) (actionResult map[string]interface{},actionError error){
//	//查询sql
//	sqlStr := "SELECT COUNT(*) AS total FROM "+cthis.tableName()+" LIMIT 1"
//
//	//查询数据
//	rows, err := cthis.sqLiteDB.Query(sqlStr)
//	if err!=nil {
//		actionError = errors.New(ipv6regionLibraryMsgPrefix+"countTableDataTotal_Query_error:"+err.Error())
//		return actionResult,actionError
//	}
//
//	//最主要释放锁
//	defer rows.Close()
//
//	//循环获取数据
//	for rows.Next() {
//		var total int
//		scanErr := rows.Scan(&total)
//		if scanErr!=nil {
//			actionError = errors.New(ipv6regionLibraryMsgPrefix+"countTableDataTotal_rowsScan_error:"+scanErr.Error())
//			return actionResult,actionError
//		}
//
//		//操作成功存储数据
//		actionResult["total"] 		= total
//
//		//数据只有1条时,只循环一次后就停止
//		break
//	}
//
//	return actionResult,actionError
//}
//
////根据条件查询1条数据
//func (cthis *ipv6regionLibrary) selectDataByOne(params map[string]interface{}) (actionResult map[string]interface{},actionError error){
//	actionResult = map[string]interface{}{
//
//	}
//	actionError = nil
//
//	//sql条件组
//	sqlWhereArr := make([]string,0)
//
//	//可选-ipv6的ip地址
//	if params["ipv6"]!=nil {
//		sqlWhereArr = append(sqlWhereArr," `ipv6` = '"+params["ipv6"].(string)+"' ")
//	}
//
//	//若一个条件也没有,则提示错误
//	if len(sqlWhereArr)<=0 {
//		actionError = errors.New(ipv6regionLibraryMsgPrefix+"selectDataByOne_sqlWhereArr_len_elt_0")
//		return actionResult,actionError
//	}
//	//sql条件字符串
//	sqlWhere := strings.Join(sqlWhereArr," AND ")
//
//	//查询sql
//	sqlStr := "SELECT * FROM "+cthis.tableName()+" WHERE "+sqlWhere+" LIMIT 1"
//
//	//查询数据
//	rows, err := cthis.sqLiteDB.Query(sqlStr)
//	if err!=nil {
//		actionError = errors.New(ipv6regionLibraryMsgPrefix+"selectDataByOne_Query_error:"+err.Error())
//		return actionResult,actionError
//	}
//
//	//注意:如果这里不加rows.Close(),下面一旦循环里使用break,就会形成database is locked报错!
//	// 往上文章说明:https://blog.csdn.net/LOVETEDA/article/details/82690498
//	// 看go的标准库代码的时候发现，rows.Close()这个方法是幂等的，当rows.Next()返回false，即所有行数据都已经遍历结束后，会自动调用rows.Close()方法。而我的代码里面，有的地方，当rows.Next()返回true的时候，在循环体当中便有break代码，导致没有调用rows.Close()方法。
//	defer rows.Close()
//
//	//循环获取数据
//	for rows.Next() {
//		var id int
//		var ipv6 string
//		var address string
//		scanErr := rows.Scan(&id, &ipv6, &address)
//		if scanErr!=nil {
//			actionError = errors.New(ipv6regionLibraryMsgPrefix+"selectDataByOne_rowsScan_error:"+scanErr.Error())
//			return actionResult,actionError
//		}
//
//		//操作成功存储数据
//		actionResult["id"] 		= id
//		actionResult["ipv6"] 	= ipv6
//		actionResult["address"] = address
//
//		//数据只有1条时,只循环一次后就停止
//		break
//	}
//
//	return actionResult,actionError
//}
////---------------------------- IPv6 sqLiteDB操作 start ----------------------------
