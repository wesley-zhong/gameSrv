package orm

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var mongoClient *mongo.Client

func InitMongoDB(addr string, userName string, pwd string) error {
	// v2 推荐直接在 options 中设置超时，不需要在 Connect 时传入 context
	mongoUrl := fmt.Sprintf("mongodb://%s", addr)

	opts := options.Client().
		ApplyURI(mongoUrl).
		SetAuth(options.Credential{
			Username: userName,
			Password: pwd,
		}).
		SetTimeout(10 * time.Second) // v2 新增：直接控制连接和操作的全局超时

	// v2 的 Connect 不再需要 context 参数
	client, err := mongo.Connect(opts)
	if err != nil {
		return fmt.Errorf("connect mongodb failed: %w", err)
	}

	// Ping 仍然需要一个 context 来控制这次探测的生命周期
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := client.Ping(ctx, nil); err != nil {
		return fmt.Errorf("ping mongodb failed: %w", err)
	}

	mongoClient = client
	return nil
}

func GetDBConnTable(dbName string, tableName string) *mongo.Collection {
	return mongoClient.Database(dbName).Collection(tableName)
}
