package orm

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var mongoClient *mongo.Client

func InitMongoDB(addr, userName, pwd string) error {
	mongoUrl := fmt.Sprintf("mongodb://%s", addr)

	opts := options.Client().
		ApplyURI(mongoUrl).
		SetAuth(options.Credential{
			Username: userName,
			Password: pwd,
		}).
		SetTimeout(10 * time.Second)

	client, err := mongo.Connect(opts)
	if err != nil {
		return fmt.Errorf("connect mongodb failed: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := client.Ping(ctx, nil); err != nil {
		return fmt.Errorf("ping mongodb failed: %w", err)
	}

	mongoClient = client
	return nil
}

func GetDBConnTable(dbName, tableName string) *mongo.Collection {
	if mongoClient == nil {
		return nil
	}
	return mongoClient.Database(dbName).Collection(tableName)
}

func CloseMongoDB() error {
	if mongoClient == nil {
		return nil
	}
	return mongoClient.Disconnect(context.Background())
}
