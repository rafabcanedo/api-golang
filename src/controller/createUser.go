package controller

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rafabcanedo/api-golang/src/configuration/rest_errors"
	"github.com/rafabcanedo/api-golang/src/configuration/validation"
	"github.com/rafabcanedo/api-golang/src/controller/model/request"
)

func CreateUser(c *gin.Context) {

	log.Println("Init CreateUser controller")
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restError := rest_errors.NewBadRequestError(
			fmt.Sprintf("There are some incorrect filds, error=%s\n", err.Error()))

			errRest := validation.ValidateUserError(err)

			c.JSON(restError.Code, errRest)

			return
	}

	fmt.Println(userRequest)
}