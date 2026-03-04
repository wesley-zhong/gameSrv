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

var Upsert = true

type SaveFun func()

var replaceOneOptions = options.Replace().SetUpsert(true) //{Upsert: &Upsert}

var workerPool = gopool.StartNewWorkerPool(16, 256)

func (dao *MongodbDAO[T]) FindOneById(id int64) (*T, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.D{{"_id", id}}
	singleResult := dao.Collection.FindOne(ctx, filter)
	retError := singleResult.Err()
	if errors.Is(retError, mongo.ErrNoDocuments) {
		return nil, nil
	}
	if retError != nil {
		log.Error(retError)
		return nil, retError
	}

	obj := new(T)
	err := singleResult.Decode(obj)
	if err != nil {
		log.Error(err)
		return nil, retError
	}
	return obj, nil
}

func (dao *MongodbDAO[T]) FindOne(filter interface{}) (*T, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	singleResult := dao.Collection.FindOne(ctx, filter)
	retError := singleResult.Err()
	if retError != nil {
		log.Error(retError)
		return nil, retError
	}
	if errors.Is(retError, mongo.ErrNoDocuments) {
		return nil, nil
	}
	newObject := new(T) //this must be a new object instance
	err := singleResult.Decode(newObject)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return newObject, nil
}

// TODO  check some return error
func (dao *MongodbDAO[T]) Insert(obj *T) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	dao.Collection.InsertOne(ctx, obj)
	return nil
}

func (dao *MongodbDAO[T]) Save(id int64, obj *T) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.D{{"_id", id}}

	ret, err := dao.Collection.ReplaceOne(ctx, filter, obj, replaceOneOptions)
	if ret.ModifiedCount > 0 {
		return nil
	}
	return err
}

func (dao *MongodbDAO[T]) AsynSave(id int64, obj *T) error {
	return workerPool.SubmitTask(func() {
		dao.Save(id, obj)
	})
}
