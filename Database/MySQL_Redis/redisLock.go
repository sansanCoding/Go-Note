package MySQL_Redis

//调用示例:
//	// -------------------------------------------------- redis锁 START ------------------------------------------------
//	lockKey := "testKey"
//	redisClient := RedisClientMap[currPlatform]
//	isLock := Lock(redisClient, lockKey, 30)
//
//	// 如果加锁处理失败,直接返回失败
//	if !isLock {
//		fmt.Println(ctx, "正在处理中!", 0)
//		return
//	}
//
//	// 加锁成功了，不管何种情况都要解锁
//	defer redis.UnLock(redisClient, lockKey)
//
//	// -------------------------------------------------- redis锁 END ------------------------------------------------


//redis单机锁-原子操作
///**
//* 获取锁
//* @param  String  $key    锁标识
//* @param  Int     $expire 锁过期时间(单位/秒)
//* @return Boolean
//* @desc 加锁与解锁示例
//       // 定义锁标识
//       lockKey = 'mylock';
//
//       // 获取锁,锁过期时间10秒
//       isOkLock := Lock(redisDb, lockKey, 10)
//
//       // 获取锁成功
//       if !isOkLock {
//           //停顿5秒,代表需要处理的的程序时间需要5秒
//           time.Sleep(5)
//           //删除锁
//           UnLock(redisDb, lockKey)
//       // 获取锁失败
//       }else{
//           fmt.Println("加锁失败")
//       }
//*/
//func Lock(redis *redis.Client, key string, expire int) bool {
//	// setnx 具有原子性,并发时,进程中只有1个操作成功的
//	var nowTime = int(time.Now().Unix())    // 当前时间
//	var val     = nowTime + expire          // 存储值
//
//	// 过期时间 转换成Duration类型
//	var expTime = time.Duration(expire)
//
//	// 并发时,只有一个请求加锁成功
//	var isLock  = redis.SetNX(key, val, expTime*time.Second)
//
//	// 不能获取锁
//	if !isLock.Val() {
//		// 判断锁是否过期
//		lockStr := redis.Get(key).Val()
//		if lockStr != "" {
//			// 锁已过期，删除锁，重新获取
//			lockTime, _ := strconv.Atoi(lockStr)
//			if nowTime > lockTime {
//				// 若获取锁成功,防止死锁,给锁加过期时间
//				redis.Del(key)
//				isLock = redis.SetNX(key, val, expTime*time.Second)
//			}
//		}
//	}
//
//	// 若获取锁成功,防止死锁,给锁加过期时间
//	if isLock.Val() {
//		redis.Expire(key, expTime*time.Second)
//	}
//
//	return isLock.Val()
//}
//
///**
// * 释放锁
// * @param  String  $key 锁标识
// * @return Boolean
// */
//func UnLock(redis *redis.Client, key string) {
//	redis.Del(key)
//}
