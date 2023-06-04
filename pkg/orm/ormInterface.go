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

type DAOIterface struct {
	Collection *mongo.Collection
	Object     interface{}
}

var Upsert = true

type SaveFun func()

var replaceOneOptions = &options.ReplaceOptions{Upsert: &Upsert}

var workerPool = gopool.StartNewWorkerPool(16, 256)

func (dao *DAOIterface) FindOneById(id int64) any {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.D{{"_id", id}}
	return dao.Collection.FindOne(ctx, filter)
}

func (dao *DAOIterface) FindOne(filter interface{}) any {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	singleResult := dao.Collection.FindOne(ctx, filter)
	if singleResult.Err() != nil {
		log.Error(singleResult.Err())
		return nil
	}
	newObject := dao.Object //this must be a new object instance
	err := singleResult.Decode(newObject)
	if err != nil {
		log.Error(err)
		return nil
	}
	return newObject
}

func (dao *DAOIterface) Insert(obj interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	dao.Collection.InsertOne(ctx, obj)
	return nil
}

func (dao *DAOIterface) Save(id int64, obj interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.D{{"_id", id}}
	ret, err := dao.Collection.ReplaceOne(ctx, filter, obj, replaceOneOptions)
	if ret.ModifiedCount > 0 {
		return nil
	}
	return err
}

func (dao *DAOIterface) AsynSave(id int64, obj interface{}) error {
	return workerPool.SubmitTask(func() {
		dao.Save(id, obj)
	})
}
