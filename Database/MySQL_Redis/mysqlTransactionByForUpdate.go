package MySQL_Redis
//
//import (
//	"Go-Note/util"
//	"database/sql"
//	"strconv"
//)
//
////余额公用处理示例,仅供参考!(mysql事物+悲观锁处理)
//
//type Account struct {
//	id             int
//	UserID         int
//	Balance        float64
//}
//
//type BalanceCommon struct {
//	//DB回调函数检测
//	DbCallbackCheck func(tx *sql.Tx, account *Account) interface{}
//	//DB回调函数检测错误信息
//	DbCallbackCheckErrorMsg string
//	//账户明细变更表日志数据插入前的回调函数
//	BalanceLogDataInsertBeforeCallback func(tx *sql.Tx, balanceLogData map[string]interface{}) bool
//}
//
////构造函数
//func NewBalanceCommon(platform string) *BalanceCommon {
//	return &BalanceCommon{
//
//	}
//}
//
////外部调用入口
////余额修改操作
//func (cthis *BalanceCommon) BalanceUpdate(info map[string]interface{}, dbCallback func(txDb *sql.Tx)(interface{}, error) ) map[string]interface{} {
//	//记录本函数产生错误时所需要排查错误的参数
//	localErrParams := map[string]interface{}{
//		"1_info": info,
//	}
//
//	//操作结果
//	actionRes := map[string]interface{}{
//		"status": 0,
//		"msg": "操作失败~!!!",
//		"exception": make(map[string]interface{}),
//	}
//
//	//获取写库
//	writeDB := MysqlClientMap["test_write"]
//	//开启事物
//	tx, txErr := writeDB.Begin()
//	if txErr!=nil {
//		//错误排查参数记录
//		localErrParams["2_msg"] = "开启事物失败,检查数据库链接~!"
//		localErrParams["3_txErr"] = txErr
//		//操作结果记录
//		actionRes["exception"] = localErrParams
//		//记录错误日志
//		cthis.errLog(localErrParams)
//		return actionRes
//	}
//
//	//是否需要检测该用户的余额:true 需要,false 不需要
//	isCheckAccountBalance := true
//	if info["isCheckAccountBalance"]!=nil {
//		isCheckAccountBalance = info["isCheckAccountBalance"].(bool)
//	}
//	//检测账户余额的标记
//	checkAccountBalanceTag := ""
//	if info["checkAccountBalanceTag"]!=nil {
//		checkAccountBalanceTag = info["checkAccountBalanceTag"].(string)
//	}
//
//	//因配置了读写数据库,所以这里也只能是读取写库数据库的数据保持余额一致
//	//@todo 这里使用的是FOR UPDATE 悲观锁
//	balance := new(Account)
//	sqlString := "SELECT `balance` FROM Account WHERE `UserID` = ? FOR UPDATE"
//	err1 := tx.QueryRow(sqlString, info["user_id"]).Scan(&balance.Balance)
//	if err1 != nil {
//		//事物回滚
//		tx.Rollback()
//		//错误排查参数记录
//		localErrParams["2_msg"] = "查询余额信息失败~!"
//		localErrParams["3_txErr"] = err1
//		//操作结果记录
//		actionRes["exception"] = localErrParams
//		//写入错误日志
//		cthis.errLog(localErrParams)
//		return actionRes
//	}
//
//	//若有自定义检测回调函数存在
//	if cthis.DbCallbackCheck != nil {
//		//将 当前的db 和 当前的用户余额数据 传入到检测回调函数中
//		checkRes := cthis.DbCallbackCheck(tx, balance)
//		//返回的结果必须是引用类型的
//		if checkRes == nil {
//			//事物回滚
//			tx.Rollback()
//			//操作结果记录
//			if cthis.DbCallbackCheckErrorMsg != "" {
//				actionRes["msg"] = cthis.DbCallbackCheckErrorMsg
//			} else {
//				actionRes["msg"] = "操作失败1~!"
//			}
//			return actionRes
//		}
//	}
//
//	//计算余额(要么是加余额的值，要么是减少余额的值)
//	balanceUpdate := map[string]interface{}{
//		"Balance": balance.Balance + float64(100),
//	}
//
//	//@update 2019.06.21 检测账户余额标记
//	if checkAccountBalanceTag!="" {
//		switch checkAccountBalanceTag {
//		//若是余额变更后小于0时,则置为0
//		case "balance_negative_change_zero":
//			if balanceUpdate["Balance"].(float64)<0 {
//				//原赋值逻辑
//				//balanceUpdate["Balance"].(float64) = float64(0)
//				//现赋值逻辑
//				//先删除后赋值,原赋值逻辑会报: cannot assign to balanceUpdate["Balance"].(float64)
//				delete(balanceUpdate,"Balance")
//				balanceUpdate["Balance"] = float64(0)
//			}
//		}
//	}
//
//	//检测余额 且 余额小于0,则提示异常
//	if isCheckAccountBalance && balanceUpdate["Balance"].(float64) < 0 {
//		//事物回滚
//		tx.Rollback()
//		//操作结果记录
//		actionRes["msg"] = "用户余额不足~!"
//		return actionRes
//	}
//
//	//账户明细变更表日志数据
//	balanceLogData := map[string]interface{}{
//		"UserID":           info["user_id"],
//		"Amount":           info["amount"],
//		"Msg":              info["msg"],
//		"Balance":          balanceUpdate["Balance"],
//		"CreateTime":       util.Helper.CurrTimeInt(),
//	}
//
//	//若写入记录之前有记录数据的回调函数需要处理,则将回调函数处理后的数据插入到记录中
//	if cthis.BalanceLogDataInsertBeforeCallback != nil {
//		//这个地方在代理佣金的地方会用到,如果操作成功返回true
//		//	balanceLogData是map,在BalanceLogDataInsertBeforeCallback回调函数中直接对该balanceLogData变量修改其值即可
//		isBLDInsertBeforeCallbackOk := cthis.BalanceLogDataInsertBeforeCallback(tx, balanceLogData)
//		if !isBLDInsertBeforeCallbackOk {
//			//事物回滚
//			tx.Rollback()
//			//操作结果记录
//			actionRes["msg"] = "操作失败2~!"
//			return actionRes
//		}
//	}
//
//	//检测记录中是否已存在该条数据,防止重复
//	//@todo 这里根据用户id进行分表数据存储处理
//	balanceLogTable := "BalanceLog_" + strconv.Itoa(info["user_id"].(int)%10)
//	var bldIsExists int
//	sqlString = "select exists(select * from " + balanceLogTable + " where UserID = ? Msg = ?) as isExists"
//	bldIsExistsErr := tx.QueryRow(sqlString, balanceLogData["UserID"], balanceLogData["Msg"]).Scan(&bldIsExists)
//	if bldIsExistsErr != nil {
//		//事物回滚
//		tx.Rollback()
//		//错误排查参数记录
//		localErrParams["2_msg"] = "检测记录中是否已存在该条数据,防止重复-查询失败~!"
//		localErrParams["3_txErr"] = bldIsExistsErr
//		//操作结果记录
//		actionRes["msg"] = "交易失败1~!"
//		actionRes["exception"] = localErrParams
//		//记录错误日志
//		cthis.errLog(localErrParams)
//		return actionRes
//	}
//	//若该数据存在
//	if bldIsExists > 0 {
//		//事物回滚
//		tx.Rollback()
//		//操作结果记录
//		actionRes["msg"] = "此订单已存在~"
//		return actionRes
//	}
//
//	//修改用户余额数据
//	balanceUpdateSetString, balanceUpdateSetValue := MysqlFiledUpdate(balanceUpdate)
//	balanceUpdateSetValue = append(balanceUpdateSetValue, info["user_id"])
//	sqlString = "update Account set " + balanceUpdateSetString + " where UserID = ?"
//	_, balanceUpdateErr := tx.Exec(sqlString, balanceUpdateSetValue...)
//	if balanceUpdateErr != nil {
//		//事物回滚
//		tx.Rollback()
//		//错误排查参数记录
//		localErrParams["2_msg"] = "修改余额信息失败~!"
//		localErrParams["3_txErr"] = balanceUpdateErr
//		//操作结果记录
//		actionRes["msg"] = "交易失败2~!"
//		actionRes["exception"] = localErrParams
//		//写入错误日志
//		cthis.errLog(localErrParams)
//		return actionRes
//	}
//
//	//写入记录
//	balanceLogDataInsertString, balanceLogDataInsertValue := MysqlFiledInsertOneRow(balanceLogData)
//	sqlString = "insert into " + balanceLogTable + balanceLogDataInsertString
//	_, balanceLogDataInsertErr := tx.Exec(sqlString, balanceLogDataInsertValue...)
//	if balanceLogDataInsertErr != nil {
//		//事物回滚
//		tx.Rollback()
//		//错误排查参数记录
//		localErrParams["2_msg"] = "插入余额变更记录失败~!"
//		localErrParams["3_txErr"] = balanceLogDataInsertErr
//		//操作结果记录
//		actionRes["msg"] = "交易失败3~!"
//		actionRes["exception"] = localErrParams
//		//写入错误日志
//		cthis.errLog(localErrParams)
//		return actionRes
//	}
//
//	//执行回调函数
//	if dbCallback != nil {
//		//将当前的db传入到回调函数中
//		_, errRet3 := dbCallback(tx)
//		//返回的结果必须是nil或者error对象
//		if errRet3 != nil {
//			//事物回滚
//			tx.Rollback()
//			//操作结果记录
//			actionRes["msg"] = "交易失败4~!"
//			return actionRes
//		}
//	}
//
//	//事物提交
//	err3 := tx.Commit()
//	//事物提交判断
//	if err3 != nil {
//		//事物回滚
//		tx.Rollback()
//		//错误排查参数记录
//		localErrParams["2_msg"] = "事物提交失败~!"
//		localErrParams["3_txErr"] = err3
//		//操作结果记录
//		actionRes["msg"] = "交易失败5~!"
//		actionRes["exception"] = localErrParams
//		//写入错误日志
//		cthis.errLog(localErrParams)
//		return actionRes
//	}
//
//	actionRes["status"] = 1
//	actionRes["msg"] = "交易成功~"
//
//	return actionRes
//}
//
////错误日志记录
//func (cthis *BalanceCommon) errLog(info map[string]interface{}) {
//	//捕获异常
//	defer func() {
//		if err := recover(); err != nil {
//			//记录异常栈
//			util.LogFile.PanicTrace(map[string]interface{}{
//				"1.func":"BalanceCommon.errLog-panic!",
//				"4.info":info,
//			},4,err)
//		}
//	}()
//
//	//转成json串(可根据需求转成json字符串存储到日志文件中)
//	//jsonStr, _ := json.Marshal(info)
//
//	//异常日志
//	util.LogFile.ErrorLog(info,"BalanceCommon.errLogDoing...")
//}