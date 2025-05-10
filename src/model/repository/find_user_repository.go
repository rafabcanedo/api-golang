package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/rafabcanedo/api-golang/src/configuration/logger"
	"github.com/rafabcanedo/api-golang/src/configuration/rest_errors"
	"github.com/rafabcanedo/api-golang/src/model"
	"github.com/rafabcanedo/api-golang/src/model/repository/entity"
	"github.com/rafabcanedo/api-golang/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func (ur *userRepository) FindAllUsers() (
	[]model.UserDomainInterface, *rest_errors.RestErrors) {

    logger.Info("Init findAllUsers repository",
        zap.String("journey", "findAllUsers"))

    collection_name := os.Getenv(MONGODB_USER_DB)
    collection := ur.databaseConnection.Collection(collection_name)

    var users []entity.UserEntity

    cursor, err := collection.Find(context.Background(), bson.D{})
    if err != nil {
        errorMessage := "Error trying to find all users"
        logger.Error(errorMessage,
            err,
            zap.String("journey", "findAllUsers"))
        return nil, rest_errors.NewInternalServerError(errorMessage)
    }
    defer cursor.Close(context.Background())

    if err := cursor.All(context.Background(), &users); err != nil {
        errorMessage := "Error trying to decode users"
        logger.Error(errorMessage,
            err,
            zap.String("journey", "findAllUsers"))
        return nil, rest_errors.NewInternalServerError(errorMessage)
    }

    // If any users exist
    if len(users) == 0 {
        errorMessage := "No users found"
        logger.Error(errorMessage,
            err,
            zap.String("journey", "findAllUsers"))
        return nil, rest_errors.NewNotFoundError(errorMessage)
    }

    var domainUsers []model.UserDomainInterface
    for _, user := range users {
        domainUsers = append(domainUsers, converter.ConvertEntityToDomain(user))
    }

    logger.Info("FindAllUsers repository executed successfully",
        zap.String("journey", "findAllUsers"),
        zap.Int("users_found", len(domainUsers)))

    return domainUsers, nil
}

func (ur *userRepository) FindUserByEmail(
	email string,
) (model.UserDomainInterface, *rest_errors.RestErrors) {

	logger.Info("Init findUserByEmail repository")
	zap.String("journey", "findUserByEmail")

	collection_name := os.Getenv(MONGODB_USER_DB)

	collection := ur.databaseConnection.Collection(collection_name)

	userEntity := &entity.UserEntity{}

	filter := bson.D{{Key: "email", Value: email}}

	err := collection.FindOne(
		context.Background(),
		filter,
	).Decode(userEntity)

	// Differentiate which error gonna be return, if email not exist or if user not exist
	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf(
				"User not found with this email: %s", email)
			logger.Error(errorMessage,
				err,
	            zap.String("journey", "findUserByEmail"))
			return nil, rest_errors.NewNotFoundError(errorMessage)
		}

		errorMessage := "Error trying to find user by email"
		logger.Error(errorMessage,
				err,
	            zap.String("journey", "findUserByEmail"))
		return nil, rest_errors.NewInternalServerError(errorMessage)
	}
	
	logger.Info("FindUserByEmail repository executed successfully",
	  zap.String("journey", "findUserByEmail"),
	  zap.String("email", email),
	  zap.String("userId", userEntity.ID.Hex()))
	return converter.ConvertEntityToDomain(*userEntity), nil
}

func (ur *userRepository) FindUserByID(
	id string,
) (model.UserDomainInterface, *rest_errors.RestErrors) {

	logger.Info("Init findUserByID repository")
	zap.String("journey", "findUserByID")

	collection_name := os.Getenv(MONGODB_USER_DB)

	collection := ur.databaseConnection.Collection(collection_name)

	userEntity := &entity.UserEntity{}

	filter := bson.D{{Key: "_id", Value: id}}

	err := collection.FindOne(
		context.Background(),
		filter,
	).Decode(userEntity)

	// Check if user exist or not exist
	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf(
				"User not found with this ID: %s", id)
			logger.Error(errorMessage,
				err,
	            zap.String("journey", "findUserByID"))
			return nil, rest_errors.NewNotFoundError(errorMessage)
		}

		errorMessage := "Error trying to find user by ID"
		logger.Error(errorMessage,
				err,
	            zap.String("journey", "findUserByID"))
		return nil, rest_errors.NewInternalServerError(errorMessage)
	}
	
	logger.Info("FindUserByID repository executed successfully",
	  zap.String("journey", "findUserByID"),
	  zap.String("userId", userEntity.ID.Hex()))
	return converter.ConvertEntityToDomain(*userEntity), nil
}