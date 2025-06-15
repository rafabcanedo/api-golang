package controller

import (
	"net/http"
	"net/mail"

	"github.com/gin-gonic/gin"
	"github.com/rafabcanedo/api-golang/src/configuration/logger"
	"github.com/rafabcanedo/api-golang/src/configuration/rest_errors"
	"github.com/rafabcanedo/api-golang/src/view"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) FindAllUsers(c *gin.Context) {

	logger.Info("Init findAllUsers controller",
	 zap.String("journey", "findAllUsers"),
    )

	usersDomain, err := uc.service.FindAllUsersServices()

	if err != nil {
		logger.Error("Error trying to call findAllUsers services",
		 err,
	     zap.String("journey", "findAllUsers"),
        )
		c.JSON(err.Code, err)
		return
	}

	logger.Info("FindAllUsers controller executed successfully",
	 zap.String("journey", "findAllUsers"),
    )

	c.JSON(http.StatusOK, view.ConvertDomainListToResponse(usersDomain))
}

func (uc *userControllerInterface) FindUserById(c *gin.Context) {

	logger.Info("Init findUserByID controller",
	 zap.String("journey", "findUserByID"),
    )

	userId := c.Param("userId")

	if _, err := primitive.ObjectIDFromHex(userId); err != nil {

		logger.Error("Error trying to validate userId",
		 err,
	     zap.String("journey", "findUserByID"),
        )

		errorMessage := rest_errors.NewBadRequestError(
			"UserID is not a valid id",
		)

		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserByIDServices(userId)

	if err != nil {
		logger.Error("Error trying to call findUserByID services",
		 err,
	     zap.String("journey", "findUserByID"),
        )
		c.JSON(err.Code, err)
		return
	}

	logger.Info("FindUserByID controller executed successfully",
	 zap.String("journey", "findUserByID"),
    )

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(
		userDomain,
	))
}

func (uc *userControllerInterface) FindUserByEmail(c *gin.Context) {

	logger.Info("Init findUserByEmail controller",
	 zap.String("journey", "findUserByEmail"),
    )

	userEmail := c.Param("userEmail")

	if _, err := mail.ParseAddress(userEmail); err != nil {

		logger.Error("Error trying to validate userEmail",
		 err,
	     zap.String("journey", "findUserByEmail"),
        )

		errorMessage := rest_errors.NewBadRequestError(
			"UserEmail is not a valid id",
		)

		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserByEmailServices(userEmail)

	if err != nil {
		logger.Error("Error trying to call findUserByEmail services",
		 err,
	     zap.String("journey", "findUserByEmail"),
        )
		c.JSON(err.Code, err)
		return
	}

	logger.Info("FindUserByEmail controller executed successfully",
	 zap.String("journey", "findUserByEmail"),
    )

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(
		userDomain,
	))
}