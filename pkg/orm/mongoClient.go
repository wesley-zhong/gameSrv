package orm

import (
	"context"
	"gameSrv/pkg/log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client

func Init() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx,
		options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(options.Credential{Username: "admin", Password: "password"}))
	if err != nil {
		panic(" connect mong db failed")
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Error(err)
	}
	mongoClient = client
}

func GetDBConnTable(dbName string, tableName string) *mongo.Collection {
	return mongoClient.Database(dbName).Collection(tableName)
}
