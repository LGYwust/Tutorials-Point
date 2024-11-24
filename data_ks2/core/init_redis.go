package core

import (
	"context"
	"data_ks2/global"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

func InitRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     global.Config.Redis.Addr(),   // Redis 服务器地址
		Password: global.Config.Redis.Password, // 没有密码
		DB:       0,
		PoolSize: global.Config.Redis.PoolSize, // 使用默认的DB
	})
	var ctx = context.Background()
	_, cancel := context.WithTimeout(ctx, 500*time.Millisecond)
	defer cancel()
	_, err := rdb.Ping().Result()
	if err != nil {
		fmt.Println("---------！！！Redis连接失败！！！---------------")
		return nil
	}
	fmt.Println("---------Redi s连接成功---------------")
	return rdb
}
