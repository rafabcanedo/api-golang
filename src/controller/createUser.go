package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/rafabcanedo/api-golang/src/configuration/rest_errors"
	"github.com/rafabcanedo/api-golang/src/controller/model/request"
)

func CreateUser(c *gin.Context) {

	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restError := rest_errors.NewBadRequestError(
			fmt.Sprintf("There are some incorrect filds, error=%s\n", err.Error()))

			c.JSON(restError.Code, restError)

			return
	}

	fmt.Println(userRequest)
}