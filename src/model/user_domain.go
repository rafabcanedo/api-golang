package model

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/rafabcanedo/api-golang/src/configuration/rest_errors"
)

func NewUserDomain(
	name, email, password string,
	age int8,
) UserDomainInterface {
	return &UserDomain{
		name, email, password, age,
	}
}

type UserDomain struct {
	Name string
	Email string
	Password string
	Age int8
}

// Encrypting user password with hash
func (ud *UserDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.Password))
	ud.Password = hex.EncodeToString(hash.Sum(nil))
}

// Why use the interface?
// If we change the database we don't change anything in controller
// because the layers are connection by interfaces
// Use interface, the tests turns be easy
type UserDomainInterface interface {
	CreateUser() *rest_errors.RestErrors
	UpdateUser(string)  *rest_errors.RestErrors
	FindUser(string) (*UserDomain, *rest_errors.RestErrors)
	DeleteUser(string) *rest_errors.RestErrors
}