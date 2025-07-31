package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/guiziin227/CRUDgo/src/model/service"
)

func NewUserController(service service.UserDomainService) UserControllerInterface {
	return &userControllerInterface{
		service: service,
	}
}

type UserControllerInterface interface {
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)

	FindUserByEmail(c *gin.Context)
	FindUserById(c *gin.Context)
}

type userControllerInterface struct {
	service service.UserDomainService
}
