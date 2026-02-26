package orm

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

var Rdb *redis.Client

func InitRedis(addr string, password string) error {
	Rdb = redis.NewClient(&redis.Options{
		Addr:         addr,
		Password:     password,        // Redis 服务器没有设置密码
		DB:           0,               // 使用默认数据库
		PoolSize:     100,             // 连接池大小
		MinIdleConns: 10,              // 最小空闲连接
		DialTimeout:  5 * time.Second, // 建立连接超时
		ReadTimeout:  3 * time.Second, // 读超时
		WriteTimeout: 3 * time.Second, // 写超时
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := Rdb.Ping(ctx).Result()
	return err
}
