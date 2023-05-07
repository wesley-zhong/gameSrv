package orm

import (
	"context"
	"time"

	"gameSrv/pkg/log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client

func Init() {
	//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//defer cancel()
	client, err := mongo.Connect(context.Background(),
		options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(options.Credential{Username: "admin", Password: "password"}))
	//defer func() {
	//	if err = client.Disconnect(ctx); err != nil {
	//		panic(err)
	//	}
	//}()
	if err != nil {
		panic(" connect mong db failed")
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Error(err)
	}
	mongoClient = client
}

var result struct {
	Value float64
}

func getObj(pid int64) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	filter := bson.D{{"id", pid}}
	collection := mongoClient.Database("testing").Collection("numbers")
	err := collection.FindOne(ctx, filter).Decode(&result)
	if err == mongo.ErrNoDocuments {
		// Do something when no record was found
		log.Infof("record does not exist")
	} else if err != nil {
		log.Error(err)
	}
}

func GetDBConnTable(dbName string, tableName string) *mongo.Collection {
	return mongoClient.Database(dbName).Collection(tableName)
}
