package orm

import (
	"context"
	"gameSrv/pkg/gopool"
	"gameSrv/pkg/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type MongodbDAOInterface[T any] struct {
	Collection *mongo.Collection
}

var Upsert = true

type SaveFun func()

var replaceOneOptions = &options.ReplaceOptions{Upsert: &Upsert}

var workerPool = gopool.StartNewWorkerPool(16, 256)

func (dao *MongodbDAOInterface[T]) FindOneById(id int64) *T {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.D{{"_id", id}}
	singleResult := dao.Collection.FindOne(ctx, filter)
	if singleResult.Err() != nil {
		log.Error(singleResult.Err())
		return nil
	}
	obj := new(T)
	err := singleResult.Decode(obj)
	if err != nil {
		log.Error(err)
		return nil
	}
	return obj
}

func (dao *MongodbDAOInterface[T]) FindOne(filter interface{}) *T {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	singleResult := dao.Collection.FindOne(ctx, filter)
	if singleResult.Err() != nil {
		//log.Error(singleResult.Err())
		return nil
	}
	newObject := new(T) //this must be a new object instance
	err := singleResult.Decode(newObject)
	if err != nil {
		log.Error(err)
		return nil
	}
	return newObject
}

func (dao *MongodbDAOInterface[T]) Insert(obj *T) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	dao.Collection.InsertOne(ctx, obj)
	return nil
}

func (dao *MongodbDAOInterface[T]) Save(id int64, obj *T) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.D{{"_id", id}}
	ret, err := dao.Collection.ReplaceOne(ctx, filter, obj, replaceOneOptions)
	if ret.ModifiedCount > 0 {
		return nil
	}
	return err
}

func (dao *MongodbDAOInterface[T]) AsynSave(id int64, obj *T) error {
	return workerPool.SubmitTask(func() {
		dao.Save(id, obj)
	})
}
