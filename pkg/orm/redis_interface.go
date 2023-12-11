package orm

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
)

type RedisDAOInterface[T any] struct {
	Rdb *redis.Client
}

func (redisDAO *RedisDAOInterface[T]) getByKey(key string) *T {
	ctx := context.Background()
	result, err := redisDAO.Rdb.Get(ctx, key).Bytes()
	if err != nil {
		return nil
	}
	obj := new(T)
	json.Unmarshal(result, obj)
	return obj
}

func (redisDAO *RedisDAOInterface[T]) setByKey(key string, obj *T) error {
	ctx := context.Background()
	strBody, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	return redisDAO.Rdb.Set(ctx, key, strBody, 0).Err()
}

/**
to do other options
*/
