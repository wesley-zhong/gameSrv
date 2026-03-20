package orm

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
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
	if err := json.Unmarshal(result, obj); err != nil {
		return nil
	}
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
	ctx := context.Background()
	data, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	return dao.Rdb.Set(ctx, key, data, ttl).Err()
}

func (dao *RedisDAO[T]) Incr(key string) (int64, error) {
	ctx := context.Background()
	return dao.Rdb.Incr(ctx, key).Result()
}

func (dao *RedisDAO[T]) Del(key string) error {
	ctx := context.Background()
	return dao.Rdb.Del(ctx, key).Err()
}
