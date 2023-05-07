package orm

import (
	"context"
	"gameSrv/pkg/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type DAOIterface struct {
	Collection *mongo.Collection
	Object     interface{}
}

func (dao *DAOIterface) FindOneById(id int64) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.D{{"_id", id}}
	dao.Collection.FindOne(ctx, filter)
}

func (dao *DAOIterface) FindOne(filter interface{}) interface{} {
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
