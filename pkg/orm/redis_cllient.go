package orm

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

var Rdb *redis.Client

func InitRedis(addr string, password string) error {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password, // Redis 服务器没有设置密码
		DB:       0,        // 使用默认数据库
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := Rdb.Ping(ctx).Result()
	return err
}
