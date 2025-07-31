package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/guiziin227/CRUDgo/src/configuration/logger"
	"github.com/guiziin227/CRUDgo/src/configuration/validation"
	"github.com/guiziin227/CRUDgo/src/controller/model/request"
	"github.com/guiziin227/CRUDgo/src/model"
	"go.uber.org/zap"
	"net/http"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func CreateUser(c *gin.Context) {
	logger.Info("Initializing CreateUser controller", zap.String("method", "CreateUser"))
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user request", err, zap.String("method", "CreateUser"))
		cErr := validation.ValidateUserError(err)
		c.JSON(cErr.Code, cErr)
		return
	}

	domain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Username,
		userRequest.Age,
	)

	if err := domain.CreateUser(); err != nil {
		logger.Error("Error trying to create user", err, zap.String("method", "CreateUser"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("Create user successfully", zap.String("method", "CreateUser"))

	c.String(http.StatusOK, "")
}
