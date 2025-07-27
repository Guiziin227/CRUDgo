package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/guiziin227/CRUDgo/src/configuration/validation"
	"github.com/guiziin227/CRUDgo/src/controller/model/request"
)

func CreateUser(c *gin.Context) {

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		cErr := validation.ValidateUserError(err)
		c.JSON(cErr.Code, cErr)
		return
	}

	fmt.Println(userRequest)
}
