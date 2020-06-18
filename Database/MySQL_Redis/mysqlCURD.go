package MySQL_Redis

//import(
//	"Go-Note/util"
//	"database/sql"
//	"fmt"
//	"strconv"
//	"strings"
//	"sync"
//)
//
////@todo mysql增删改查
//
//var readDB = config.MysqlClientMap["test_read"].(*sql.DB)
//var writeDB = config.MysqlClientMap["test_read"].(*sql.DB)
//
////1.查询sql
//userSql := "select * from Users limit 1"
//userRows, err := readDB.Query(userSql)
//if err != nil {
//	panic(fmt.Sprintf("%+v",map[string]interface{}{
//		"0.userSql": userSql,
//		"1.sqlErr": err,
//	})
//}
////获取结果集
//userResults := GetMysqlResultSets(userRows)
//
//
////2.插入和修改数据
////若数据存在,则修改
//if len(findRes)>0 {
//	//获取数据主键id
//	findId := findRes["id"].(int)
//	//修改数据
//	dataResults["UpdateTime"] = util.Helper.CurrTimeInt()
//	updateField,updateValue := MysqlFiledUpdate(dataResults)
//	updateValue = append(updateValue, findId)
//	updateSqlStr := "UPDATE `Users` SET "+updateField+" WHERE `id`=?"
//	_,sqlErr := writeDB.Exec(updateSqlStr,updateValue...)
//	if sqlErr!=nil {
//		panic(map[string]interface{}{
//			"0.updateSqlStr":updateSqlStr,
//			"1.updateValue":updateValue,
//			"2.sqlErr":sqlErr,
//		})
//	}
////若数据不存在则插入
//}else{
//	//插入数据
//	dataResults["CreateTime"] = util.Helper.CurrTimeInt()
//	insertField,insertValue := MysqlFiledInsertOneRow(dataResults)
//	insertSqlStr := "INSERT INTO `Users` "+insertField
//	_,sqlErr := writeDB.Exec(insertSqlStr,insertValue...)
//	if sqlErr!=nil {
//		panic(map[string]interface{}{
//			"0.insertSqlStr":insertSqlStr,
//			"1.insertValue":insertValue,
//			"2.sqlErr":sqlErr,
//		})
//	}
//}
//
////3.删除
////删除条件参数
//deleteArgs := []interface{}{
//	findId,
//}
////拼接删除sql
//deleteSql := " DELETE FROM Users WHERE id=?"
//
////执行sql
//_,err := writeDB.Exec(deleteSql,deleteArgs...)
//if err!=nil {
//	panic(map[string]interface{}{
//		"0.sqlErr":err,
//		"1.deleteSql":deleteSql,
//		"2.deleteArgs":deleteArgs,
//	})
//}
//
////@todo sql增删改查工具集
////------------------------------------- sql增删改查工具集 start -------------------------------------
//import (
//	"database/sql"
//	"strconv"
//	"strings"
//	"sync"
//)
//
////获取结果集
//func GetMysqlResultSets(rows *sql.Rows) (container []map[string]interface{}) {
//	//捕获异常
//	defer func() {
//		if err := recover(); err != nil {
//			//记录异常栈
//			util.LogFile.PanicTrace(map[string]interface{}{
//				"1.func":"GetMysqlResultSets-panic!",
//			},4,err)
//		}
//	}()
//	columnTypes, _ := rows.ColumnTypes()
//	container = make([]map[string]interface{}, 0)
//	//values是每个列的值，这里获取到byte里
//	var values = make([]sql.RawBytes, len(columnTypes))
//	//query.Scan的参数，因为每次查询出来的列是不定长的，用len(cols)定住当次查询的长度
//	var scans = make([]interface{}, len(columnTypes))
//	//让每一行数据都填充到[][]byte里面
//	for i := range values {
//		scans[i] = &values[i]
//	}
//	for rows.Next() { //循环，让游标往下推
//		if err := rows.Scan(scans...); err != nil {
//			//query.Scan查询出来的不定长值放到scans[i] = &values[i],也就是每行都放在values里
//			rows.Close()
//			panic(err)
//		}
//		parseRow := make(map[string]interface{})
//		for k, val := range values { //每行数据是放在values里面，现在把它挪到row里
//			fieldObject := *columnTypes[k]
//			columnName := fieldObject.Name()
//			columnType := fieldObject.DatabaseTypeName()
//			var parseValue interface{}
//			valString := string(val)
//			switch columnType {
//			case "INT":
//				parseValue, _ = strconv.Atoi(valString)
//			case "VARCHAR":
//				parseValue = valString
//			case "FLOAT":
//				parseValue, _ = strconv.ParseFloat(valString, 64)
//			case "CHAR":
//				parseValue = valString
//			case "TINYINT":
//				parseValue, _ = strconv.Atoi(valString)
//			case "MEDIUMINT":
//				parseValue, _ = strconv.Atoi(valString)
//			case "DECIMAL":
//				parseValue, _ = strconv.ParseFloat(valString, 64)
//			default:
//				parseValue = valString
//			}
//			parseRow[columnName] = parseValue
//		}
//		container = append(container, parseRow)
//	}
//	return container
//}
//
////update字段组装
//func MysqlFiledUpdate(data map[string]interface{}) (string, []interface{}) {
//	var fields string
//	var updateValue []interface{}
//	for k, v := range data {
//		fields += k + "= ?,"
//		updateValue = append(updateValue, v)
//
//	}
//	fields = strings.TrimRight(fields, ",")
//	return fields, updateValue
//}
//
////insert字段组装
//func MysqlFiledInsertOneRow(data map[string]interface{}) (string, []interface{}) {
//	var fields string = " ("
//	var mark string = "("
//	var insertValue []interface{}
//	for k, v := range data {
//		fields += k + ","
//		mark += "?,"
//		insertValue = append(insertValue, v)
//
//	}
//	fields = strings.TrimRight(fields, ",")
//	mark = strings.TrimRight(mark, ",")
//	fields += ") values " + mark + ")"
//	return fields, insertValue
//}
//
////insert批量数据字段组装
////调用示例:
////	insertData := []map[string]interface{}{
////		{"name":"test1","age":21,"address":"a1","where":"w1"},
////		{"name":"test2","age":22,"address":"a2","where":"w222"},
////		{"name":"test3","age":23,"address":"a3","where":"w33"},
////		{"name":"test4","age":24,"address":"a4","where":"w44444"},
////	}
////	fieldStr,preAllStr,valAllArr := MysqlFiledInsertMoreRow(insertData)
////最终输出:
////	fieldStr 	=> (`where`,`address`,`age`,`name`)
////	preAllStr 	=> (?,?,?,?),(?,?,?,?),(?,?,?,?),(?,?,?,?)
////	valAllArr	=> [w1 a1 21 test1 w222 a2 22 test2 w33 a3 23 test3 w44444 a4 24 test4]
//func MysqlFiledInsertMoreRow(insertData []map[string]interface{}) (fieldStr string,preAllStr string,valAllArr []interface{}) {
//	//存储字段位置
//	fieldMap := make(map[string]int)
//	//第1次循环:确认字段位置
//	for _,v := range insertData {
//		if len(fieldMap)==0 {
//			//索引标记位
//			indexTag := len(v) - 1
//			//字段名称为key,索引标记位为值
//			for fieldName := range v {
//				fieldMap[fieldName] = indexTag
//				indexTag--
//			}
//		}
//		break
//	}
//	//处理完后fieldMap输出:map[address:1 age:2 name:3 where:0]
//
//	//统计总共有多少个字段
//	fieldMapLen := len(fieldMap)
//
//	//第2步:处理字段名的顺序位
//	fieldStrArr := make([]string,fieldMapLen)
//	for fieldName,fieldIndexTag := range fieldMap {
//		fieldStrArr[fieldIndexTag] = fieldName
//	}
//	//处理完后fieldStr输出:(`where`,`address`,`age`,`name`)
//	fieldStr = "(`"+strings.Join(fieldStrArr,"`,`")+"`)"
//
//	//第3步:处理字段值的顺序位
//	//预占位的数组
//	preAllStr = ""
//	//值占位的数组
//	valAllArr = make([]interface{},0)
//	for _,v := range insertData {
//		preArr := make([]string,fieldMapLen)
//		valArr := make([]interface{},fieldMapLen)
//		for fieldName,fieldVal := range v {
//			fieldIndexTag := fieldMap[fieldName]
//			preArr[fieldIndexTag] = "?"
//			valArr[fieldIndexTag] = fieldVal
//		}
//		//preArr是[? ? ? ?]
//		preStr := strings.Join(preArr,",")
//		//拼接效果如:preStr是?,?,?,? preAllStr输出效果是,(?,?,?,?),(?,?,?,?)
//		preAllStr = poolStringJoin("",preAllStr,",(", preStr, ")")
//		//拼接效果如:valArr是[w1 a1 21 test1] valAllArr输出效果是[w1 a1 21 test1 w222 a2 22 test2]
//		valAllArr = append(valAllArr,valArr...)
//	}
//
//	return fieldStr,strings.Trim(preAllStr,","),valAllArr
//}
//
//// insert字段组装 字段加上`，避免特殊字段导致插入错误 (`aa`,`bb`,`cc`) value (?,?,?)
//func MysqlFiledInsertOneRowEscape(data map[string]interface{}) (string, []interface{}) {
//	var fields string = " (`"
//	var mark string = "("
//	var insertValue []interface{}
//	for k, v := range data {
//		fields += k + "`,`"
//		mark += "?,"
//		insertValue = append(insertValue, v)
//	}
//	// TrimRight只能去掉一个字符,这里需要去掉两次不同的字符
//	fields = strings.TrimRight(fields, "`")
//	fields = strings.TrimRight(fields, ",")
//	mark = strings.TrimRight(mark, ",")
//	fields += ") values " + mark + ")"
//	return fields, insertValue
//}
////------------------------------------- sql增删改查工具集 end -------------------------------------
//
////------------------------------------ 工具集 start ------------------------------------
//
////strings.Builder字符串拼接处理-对象池
//var stringBuilderPool = &sync.Pool{
//	New: func() interface{} {
//		return new(strings.Builder)
//	},
//}
////拼接字符串,效果如poolStringJoin(",","a","b") 输出如 a,b
//func poolStringJoin(sep string,strMore ...string) string {
//	strMoreLen := len(strMore)
//
//	//获取该对象
//	b := stringBuilderPool.Get().(*strings.Builder)
//	//每次都将该对象的数据重置清空
//	b.Reset()
//
//	//循环将每个字符串压入到对象中
//	for i:=0;i<strMoreLen;i++ {
//		b.WriteString(strMore[i])
//		//若是最后一个则不需要拼接符号
//		if i<strMoreLen-1 {
//			b.WriteString(sep)
//		}
//	}
//	s := b.String()
//
//	//存储到对象池中
//	stringBuilderPool.Put(b)
//	return s
//}
//
////------------------------------------ 工具集 end ------------------------------------