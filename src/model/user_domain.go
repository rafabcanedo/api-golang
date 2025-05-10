package model

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

type UserDomainInterface interface {
	GetName() string
	GetEmail() string
	GetPassword() string
	GetAge() int8

	SetID(string)

	GetJSONValue() (string, error)

	EncryptPassword()
}

func NewUserDomain(
	name, email, password string,
	age int8,
) UserDomainInterface {
	return &userDomain{
		Name: name,
		Email: email,
		Password: password,
		Age: age,
	}
}

func (ud *userDomain) SetID(id string) {
	ud.ID = id
}

type userDomain struct {
	ID string
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	Age int8 `json:"age"`
}

func (ud *userDomain) GetJSONValue() (string, error) {
	b, err := json.Marshal(ud)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return string(b), nil
}

func(ud *userDomain) GetName() string {
	return ud.Name
}
func(ud *userDomain) GetEmail() string {
	return ud.Email
}
func(ud *userDomain) GetPassword() string {
	return ud.Password
}
func(ud *userDomain) GetAge() int8 {
	return ud.Age
}

// Encrypting user password with hash
func (ud *userDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.Password))
	ud.Password = hex.EncodeToString(hash.Sum(nil))
}