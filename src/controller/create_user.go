package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rafabcanedo/api-golang/src/configuration/logger"
	"github.com/rafabcanedo/api-golang/src/configuration/validation"
	"github.com/rafabcanedo/api-golang/src/controller/model/request"
	"github.com/rafabcanedo/api-golang/src/model"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func CreateUser(c *gin.Context) {

	logger.Info("Init CreateUser controller",
		zap.String("journey", "createUser"),
	)
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {

		logger.Error("Error trying to validate user info", err,
		zap.String("journey", "createUser"))
		errRest := validation.ValidateUserError(err)

			c.JSON(errRest.Code, errRest)

			return
	}

	domain := model.NewUserDomain(
		userRequest.Name,
		userRequest.Email,
		userRequest.Password,
		userRequest.Age,
	)

	if err := domain.CreateUser(); err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User created successfully",
	zap.String("journey", "createUser"))

	c.String(http.StatusOK, "")
}