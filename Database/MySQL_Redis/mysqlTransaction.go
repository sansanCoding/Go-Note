package MySQL_Redis

//简单DB事物提交示例,仅供参考!

func Test(){

	////---------------------------- DB事务处理 start ----------------------------
	//writeDB := MysqlClientMap["test_write"]
	//
	////开启事物
	//tx,txErr := writeDB.Begin()
	//if txErr!=nil {
	//	//抛出异常
	//	panic("writeDB.BeginErr:"+txErr.Error())
	//}
	//
	////先查询后删除
	//sqlStr :=`select id from xxx where id=xxx limit 1`
	//sqlRows,sqlRowsErr := tx.Query(sqlStr)
	//if sqlRowsErr!=nil {
	//	panic("tx.QueryErr:"+sqlRowsErr.Error())
	//}
	//sqlData := GetMysqlResultSets(sqlRows)
	//if len(sqlData) > 0{
	//	deleteSql :=`DELETE FROM xx WHERE id=xx`
	//	_, deleteErr := tx.Exec(deleteSql)
	//	if deleteErr != nil {
	//		//事物回滚
	//		tx.Rollback()
	//		//抛出异常
	//		panic("tx.ExecErr:"+deleteErr.Error())
	//	}
	//}
	//
	////修改和删除可获取影响行
	////执行修改操作
	////		updateRes, updateResErr := tx.Exec(updateSql, setValue...)
	////		//获取修改数据影响行
	////		updateRowsAffectedInt64,updateRowsAffectedErr := updateRes.RowsAffected()
	////		if updateRowsAffectedErr!=nil {
	////			panic("tx.updateRowsAffectedErr:"+updateRowsAffectedErr.Error())
	////		}
	////		//若修改数据没有影响行
	////		if updateRowsAffectedInt64<=0 {
	////			panic("tx.updateRowsAffectedInt64_elt_0")
	////		}
	//
	////插入获取的最后插入id
	////执行插入操作
	////		//执行sql
	////		insertRes,insertResErr := tx.Exec(insertSql, insertValue...)
	////		//若有错误返回
	////		if insertResErr!=nil {
	////			panic("tx.insertResErr:"+insertResErr.Error())
	////		}
	////		//如果最终插入有影响行存在
	////		insertRows,insertRowsErr := insertRes.RowsAffected()
	////		if insertRowsErr!=nil {
	////			panic("tx.insertRowsErr:"+insertRowsErr.Error())
	////		}
	////		if insertRows>0 {
	////			//获取插入的最后id
	////			tempInsertId,_ := insertRes.LastInsertId()
	////		}
	//
	////事务提交
	//commitErr := tx.Commit()
	//if commitErr != nil {
	//	//事物回滚
	//	tx.Rollback()
	//	//抛出异常
	//	panic("tx.CommitErr:"+commitErr.Error())
	//}
	////---------------------------- DB事务处理 end ----------------------------

}