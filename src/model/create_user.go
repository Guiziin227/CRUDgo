package model

import (
	"fmt"
	"github.com/guiziin227/CRUDgo/src/configuration/c_err"
	"github.com/guiziin227/CRUDgo/src/configuration/logger"
	"go.uber.org/zap"
)

func (ud *userDomain) CreateUser() *c_err.CErr {

	logger.Info("Creating user in UserDomain", zap.String("journey", "CreateUser"))

	ud.EncryptPassword()

	fmt.Println(ud)

	return nil
}
