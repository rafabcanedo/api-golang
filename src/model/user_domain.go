package model

import (
	"crypto/md5"
	"encoding/hex"
)

type UserDomainInterface interface {
	GetName() string
	GetEmail() string
	GetPassword() string
	GetAge() int8

	EncryptPassword()
}

func NewUserDomain(
	name, email, password string,
	age int8,
) UserDomainInterface {
	return &userDomain{
		name, email, password, age,
	}
}

type userDomain struct {
	name string
	email string
	password string
	age int8
}

func(ud *userDomain) GetName() string {
	return ud.name
}
func(ud *userDomain) GetEmail() string {
	return ud.email
}
func(ud *userDomain) GetPassword() string {
	return ud.password
}
func(ud *userDomain) GetAge() int8 {
	return ud.age
}

// Encrypting user password with hash
func (ud *userDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.password))
	ud.password = hex.EncodeToString(hash.Sum(nil))
}