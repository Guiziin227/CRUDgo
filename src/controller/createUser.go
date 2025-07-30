package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/guiziin227/CRUDgo/src/configuration/logger"
	"github.com/guiziin227/CRUDgo/src/configuration/validation"
	"github.com/guiziin227/CRUDgo/src/controller/model/request"
	"github.com/guiziin227/CRUDgo/src/controller/model/response"
	"go.uber.org/zap"
	"net/http"
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

	response := response.UserResponse{
		Email:    userRequest.Email,
		Username: userRequest.Username,
		Age:      userRequest.Age,
	}

	logger.Info("Create user successfully", zap.String("method", "CreateUser"))

	c.JSON(http.StatusOK, response)
}
