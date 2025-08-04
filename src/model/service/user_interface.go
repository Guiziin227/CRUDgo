package service

import (
	"github.com/guiziin227/CRUDgo/src/configuration/c_err"
	"github.com/guiziin227/CRUDgo/src/model"
	"github.com/guiziin227/CRUDgo/src/model/repository"
)

func NewUserDomainService(userRepository repository.UserRepository) UserDomainService {
	return &userDomainService{
		userRepository: userRepository,
	}
}

type userDomainService struct {
	userRepository repository.UserRepository
}

type UserDomainService interface {
	CreateUser(model.UserDomainInterface) (model.UserDomainInterface, *c_err.CErr)
	UpdateUser(string, model.UserDomainInterface) *c_err.CErr
	FindUser(string) (*model.UserDomainInterface, *c_err.CErr)
	DeleteUser(string) *c_err.CErr
}
