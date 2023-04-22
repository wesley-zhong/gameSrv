package orm

import (
	"context"
	"time"

	"gateway/pkg/log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func init() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	if err != nil {
		panic(" connetct mong db failed")
	}
	MongoClient = client
}

var result struct {
	Value float64
}

func getObj(pid int64) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.D{{"id", pid}}
	collection := MongoClient.Database("testing").Collection("numbers")
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		// Do something when no record was found
		log.Infof("record does not exist")
	} else if err != nil {
		log.Error(err)
	}
}
