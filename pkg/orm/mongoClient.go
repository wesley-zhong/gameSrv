package orm

import (
	"context"
	"gameSrv/pkg/log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client

func Init(addr string, userName string, pwd string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoUrl := "mongodb://" + addr
	client, err := mongo.Connect(ctx,
		options.Client().ApplyURI(mongoUrl).SetAuth(options.Credential{Username: userName, Password: pwd}))
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
