package service

import (
	"github.com/guiziin227/CRUDgo/src/configuration/c_err"
	"github.com/guiziin227/CRUDgo/src/model"
)

func NewUserDomainService() UserDomainService {
	return &userDomainService{}
}

type userDomainService struct {
}

type UserDomainService interface {
	CreateUser(model.UserDomainInterface) *c_err.CErr
	UpdateUser(string, model.UserDomainInterface) *c_err.CErr
	FindUser(string) (*model.UserDomainInterface, *c_err.CErr)
	DeleteUser(string) *c_err.CErr
}
