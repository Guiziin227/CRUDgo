package mongodb

import (
	"context"
	"github.com/guiziin227/CRUDgo/src/configuration/logger"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"os"
)

var (
	MONGODB_URI     = "MONGODB_URI"
	MONGODB_USER_DB = "MONGODB_USER_DB"
)

func NewMongoDBConnection(ctx context.Context) (*mongo.Database, error) {
	mongodb_uri := os.Getenv(MONGODB_URI)
	mongodb_user := os.Getenv(MONGODB_USER_DB)

	client, err := mongo.Connect(options.Client().ApplyURI(mongodb_uri))
	if err != nil {
		return nil, err
	}

	if err = client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	logger.Info("Connected to MongoDB")
	return client.Database(mongodb_user), nil
}
