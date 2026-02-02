package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/opaulosaa/crud-go/src/controller/model/request"
	"github.com/opaulosaa/crud-go/src/rest_err"
)

func CreateUser(c *gin.Context) {
	var UserRequest request.UserRequest

	if err := c.ShouldBindJSON(&UserRequest); err != nil {
		restErr := rest_err.NewBadRequestError(
			fmt.Sprintf("There are some incorrect fields, error=%s\n", err.Error()))
		c.JSON(restErr.Code, restErr)
		return
	}
	fmt.Println(UserRequest)
}
