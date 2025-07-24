package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/guiziin227/CRUDgo/src/configuration/c_err"
	"github.com/guiziin227/CRUDgo/src/controller/model/request"
)

func CreateUser(c *gin.Context) {

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		cErr := c_err.NewBadRequestErr(fmt.Sprintf("Invalid request body: %v\n", err.Error()))
		c.JSON(cErr.Code, cErr)
	}

	fmt.Println(userRequest)
}
