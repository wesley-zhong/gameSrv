package orm

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisDAO[T any] struct {
	Rdb *redis.Client
}

func (dao *RedisDAO[T]) GetByKey(key string) *T {
	ctx := context.Background()
	result, err := dao.Rdb.Get(ctx, key).Bytes()
	if err != nil {
		return nil
	}
	obj := new(T)
	json.Unmarshal(result, obj)
	return obj
}

func (dao *RedisDAO[T]) SetByKey(key string, obj *T) error {
	ctx := context.Background()
	strBody, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	return dao.Rdb.Set(ctx, key, strBody, 0).Err()
}

func (dao *RedisDAO[T]) SetWithTTL(key string, obj *T, ttl time.Duration) error {
	data, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	ctx := context.Background()
	return dao.Rdb.Set(ctx, key, data, ttl).Err()
}

func (dao *RedisDAO[T]) Incr(key string) (int64, error) {
	ctx := context.Background()
	return dao.Rdb.Incr(ctx, key).Result()
}

/**
to dos other options
*/
