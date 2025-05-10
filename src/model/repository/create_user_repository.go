package repository

import (
	"context"
	"os"

	"github.com/rafabcanedo/api-golang/src/configuration/logger"
	"github.com/rafabcanedo/api-golang/src/configuration/rest_errors"
	"github.com/rafabcanedo/api-golang/src/model"
	"github.com/rafabcanedo/api-golang/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

const (
	MONGODB_USER_DB = "MONGO_USER_DB"
)

func (ur *userRepository) CreateUser(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_errors.RestErrors) {

	logger.Info("Init createUser repository")
	zap.String("journey", "createUser")

	collection_name := os.Getenv(MONGODB_USER_DB)

	collection := ur.databaseConnection.Collection(collection_name)

	value := converter.ConvertDomainToEntity(userDomain)

	result, err := collection.InsertOne(context.Background(), value)
	if err != nil {
		logger.Error("Error trying to create user",
	       err,
		   zap.String("journey", "createUser"))
		return nil, rest_errors.NewInternalServerError(err.Error())
	}

	value.ID = result.InsertedID.(primitive.ObjectID)

	logger.Info(
		"CreateUser repository executed successfully",
		zap.String("userId", value.ID.Hex()),
		zap.String("journey", "createUser"))

	return converter.ConvertEntityToDomain(*value), nil

}