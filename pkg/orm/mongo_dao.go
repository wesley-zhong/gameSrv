package orm

import (
	"context"
	"errors"
	"gameSrv/pkg/gopool"
	"gameSrv/pkg/log"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type MongodbDAO[T any] struct {
	Collection *mongo.Collection
}

var upsertOptions = options.Replace().SetUpsert(true)
var workerPool = gopool.StartNewWorkerPool(16, 256)

const defaultTimeout = 5 * time.Second

func (dao *MongodbDAO[T]) FindOneById(id int64) (*T, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	filter := bson.D{{Key: "_id", Value: id}}
	singleResult := dao.Collection.FindOne(ctx, filter)

	if err := singleResult.Err(); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		log.Error(err)
		return nil, err
	}

	obj := new(T)
	if err := singleResult.Decode(obj); err != nil {
		log.Error(err)
		return nil, err
	}
	return obj, nil
}

func (dao *MongodbDAO[T]) FindOne(filter any) (*T, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	singleResult := dao.Collection.FindOne(ctx, filter)
	if err := singleResult.Err(); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		log.Error(err)
		return nil, err
	}

	obj := new(T)
	if err := singleResult.Decode(obj); err != nil {
		log.Error(err)
		return nil, err
	}
	return obj, nil
}

func (dao *MongodbDAO[T]) Insert(obj *T) error {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	_, err := dao.Collection.InsertOne(ctx, obj)
	return err
}

func (dao *MongodbDAO[T]) Save(id int64, obj *T) error {
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()

	filter := bson.D{{Key: "_id", Value: id}}
	result, err := dao.Collection.ReplaceOne(ctx, filter, obj, upsertOptions)

	if err != nil {
		return err
	}
	if result.ModifiedCount > 0 {
		return nil
	}
	return errors.New("no document modified")
}

func (dao *MongodbDAO[T]) AsynSave(id int64, obj *T) error {
	return workerPool.SubmitTask(id, func() {
		dao.Save(id, obj)
	})
}
