package repository

import (
	"github.com/guiziin227/CRUDgo/src/configuration/c_err"
	"github.com/guiziin227/CRUDgo/src/model"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func NewUserRepository(databaseConnection *mongo.Database) UserRepository {
	return &userRepository{
		databaseConnection: databaseConnection,
	}
}

type userRepository struct {
	databaseConnection *mongo.Database
}

type UserRepository interface {
	CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *c_err.CErr)
}
