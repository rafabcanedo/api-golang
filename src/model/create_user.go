package model

import (
	"fmt"

	"github.com/rafabcanedo/api-golang/src/configuration/logger"
	"github.com/rafabcanedo/api-golang/src/configuration/rest_errors"
	"go.uber.org/zap"
)

func (ud *UserDomain) CreateUser() *rest_errors.RestErrors {
	
	logger.Info("Initializing createUser model", zap.String("journey", "createUser"))

	ud.EncryptPassword()

	fmt.Println(ud)
    return nil
}