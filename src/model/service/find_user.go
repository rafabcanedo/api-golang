package service

import (
	"github.com/rafabcanedo/api-golang/src/configuration/rest_errors"
	"github.com/rafabcanedo/api-golang/src/model"
)

func (*userDomainService) FindUser(string) (
	*model.UserDomainInterface, *rest_errors.RestErrors) {
	return nil, nil
}