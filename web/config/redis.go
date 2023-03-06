package config

import (
	"context"
	"github.com/go-redis/redis/v9"
	"log"
	"time"
	"web/common"
)

// 这个必须配置正确否则无法连接,也不会报错
// var rdb *redis.ClusterClient //集群
var rdb *redis.Client

func readsInit() {
	log.Println("redis:初始化！")
	//redis集群用法
	config := common.AppConfig.RedisConfig
	if len(config.List) == 1 {
		addr := config.List[0]
		//redis单机用法
		rdb = redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: config.Pass, // no password set
			Username: config.User,
		})
		RedisSet("test", "Hello Redis")
		log.Println("Redis数据:" + RedisGetString("test"))
	}
}

// RedisSet 2021/4/23 过期时间
func RedisSet(k string, v interface{}) {
	if err := rdb.Set(context.Background(), k, v, 0).Err(); err != nil {
		log.Panic(err)
	}
}

// RedisSetExpiration 2021/4/23 过期时间
func RedisSetExpiration(k string, v interface{}, expiration time.Duration) {
	if err := rdb.Set(context.Background(), k, v, expiration).Err(); err != nil {
		log.Panic(err)
	}
}

// RedisGetString 2021/4/23 获取数据
func RedisGetString(k string) string {
	v, err := rdb.Get(context.Background(), k).Result()
	if err == redis.Nil {
		return ""
	} else if err != nil {
		log.Panic(err)
	}
	return v
}

// RedisGetInt 2021/4/23 获取数据
func RedisGetInt(k string) int64 {
	v, err := rdb.Get(context.Background(), k).Int64()
	if err == redis.Nil {
		return 0
	} else if err != nil {
		log.Panic(err)
	}
	return v
}

// RedisLock 2021/4/23 分布式锁,上锁成功为true
func RedisLock(k string, expiration time.Duration) bool {
	nx := rdb.SetNX(context.Background(), k, k, expiration)
	result, err := nx.Result()
	if err != nil {
		log.Panic(err)
	}
	return result
}

func RedisDel(k string) {
	rdb.Del(context.Background(), k)
}
