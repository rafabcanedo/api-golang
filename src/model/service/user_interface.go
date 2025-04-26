package service

import (
	"github.com/rafabcanedo/api-golang/src/configuration/rest_errors"
	"github.com/rafabcanedo/api-golang/src/model"
)

func NewUserDomainService() UserDomainService {
	return &userDomainService{}
}

type userDomainService struct {

}

// Why use the interface?
// If we change the database we don't change anything in controller
// because the layers are connection by interfaces
// Use interface, the tests turns be easy
type UserDomainService interface {
	CreateUser(model.UserDomainInterface) *rest_errors.RestErrors
	UpdateUser(string, model.UserDomainInterface)  *rest_errors.RestErrors
	FindUser(string) (*model.UserDomainInterface, *rest_errors.RestErrors)
	DeleteUser(string) *rest_errors.RestErrors
}