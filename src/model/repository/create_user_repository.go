package repository

import (
	"context"
	"github.com/guiziin227/CRUDgo/src/configuration/c_err"
	"github.com/guiziin227/CRUDgo/src/configuration/logger"
	"github.com/guiziin227/CRUDgo/src/model"
	"go.uber.org/zap"
	"os"
)

const MONGODB_USER_COLLECTION = "MONGODB_USER_COLLECTION"

func (ur *userRepository) CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *c_err.CErr) {
	logger.Info("Init CreateUser in UserRepository", zap.String("reposity", "CreateUser"))

	collection_name := os.Getenv(MONGODB_USER_COLLECTION)

	collection := ur.databaseConnection.Collection(collection_name)

	value, err := userDomain.GetJSONValue()
	if err != nil {
		return nil, c_err.NewInternalServerErr("Error getting JSON value from user domain")
	}

	result, err := collection.InsertOne(context.Background(), value)
	if err != nil {
		return nil, c_err.NewInternalServerErr("Error getting JSON value from user domain")
	}

	userDomain.SetID(result.InsertedID.(string))
}
