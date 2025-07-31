package mongodb

import (
	"context"
	"github.com/guiziin227/CRUDgo/src/configuration/logger"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func InitConnection() {

	ctx := context.Background()

	client, _ := mongo.Connect(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err := client.Ping(ctx, nil); err != nil {
		panic(err)
	}

	logger.Info("Connected to MongoDB")
}
