package service

import (
	"fmt"
	"github.com/guiziin227/CRUDgo/src/configuration/c_err"
	"github.com/guiziin227/CRUDgo/src/configuration/logger"
	"github.com/guiziin227/CRUDgo/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(userDomain model.UserDomainInterface) *c_err.CErr {

	logger.Info("Creating user in UserDomain", zap.String("Model", "CreateUser"))

	userDomain.EncryptPassword()

	fmt.Println(userDomain.GetPassword())

	return nil
}
