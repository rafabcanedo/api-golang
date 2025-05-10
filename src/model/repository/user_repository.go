package repository

import (
	"github.com/rafabcanedo/api-golang/src/configuration/rest_errors"
	"github.com/rafabcanedo/api-golang/src/model"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	MONGODB_USER_DB = "MONGO_USER_DB"
)

func NewUserRepository(
	database *mongo.Database,
) UserRepository {
	return &userRepository{
		database,
	}
}

type userRepository struct {
	databaseConnection *mongo.Database
}

type UserRepository interface {
	CreateUser(
		userDomain model.UserDomainInterface,
	) (model.UserDomainInterface, *rest_errors.RestErrors)

	FindAllUsers() ([]model.UserDomainInterface, *rest_errors.RestErrors)

	FindUserByEmail(
		email string,
	) (model.UserDomainInterface, *rest_errors.RestErrors)

	FindUserByID(
		id string,
	) (model.UserDomainInterface, *rest_errors.RestErrors)
}