package service

import (
	"github.com/rafabcanedo/api-golang/src/configuration/logger"
	"github.com/rafabcanedo/api-golang/src/configuration/rest_errors"
	"github.com/rafabcanedo/api-golang/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) FindAllUsersServices() (
	[]model.UserDomainInterface, *rest_errors.RestErrors) {
	
	logger.Info("Init findAllUser services",
    zap.String("journey", "findAllUser"))
	
	return ud.userRepository.FindAllUsers()
}
func (ud *userDomainService) FindUserByIDServices(id string) (
	model.UserDomainInterface, *rest_errors.RestErrors) {
	
	logger.Info("Init findUserByID services",
    zap.String("journey", "findUserById"))
	
	return ud.userRepository.FindUserByID(id)
}

func (ud *userDomainService) FindUserByEmailServices(email string) (
	model.UserDomainInterface, *rest_errors.RestErrors) {
	
	logger.Info("Init findUserByEmail services",
    zap.String("journey", "findUserByEmail"))
	
	return ud.userRepository.FindUserByEmail(email)
}