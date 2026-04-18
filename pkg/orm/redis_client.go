package orm

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

var Rdb *redis.Client

func InitRedis(addr, password string) error {
	Rdb = redis.NewClient(&redis.Options{
		Addr:         addr,
		Password:     password,
		DB:           0,
		PoolSize:     100,
		MinIdleConns: 10,
		DialTimeout:  5 * time.Second,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := Rdb.Ping(ctx).Result()
	return err
}

func CloseRedis() error {
	if Rdb == nil {
		return nil
	}
	return Rdb.Close()
}

func GetRedisClient() *redis.Client {
	return Rdb
}
