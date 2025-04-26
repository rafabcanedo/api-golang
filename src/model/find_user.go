package model

import (
	"github.com/rafabcanedo/api-golang/src/configuration/rest_errors"
	"github.com/rafabcanedo/api-golang/src/model"
)

func (*UserDomain) FindUser(string) (model.UserDomainInterface, *rest_errors.RestErrors) {
	return nil, nil
}