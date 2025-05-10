package model

import (
	"encoding/json"
	"fmt"
)

type userDomain struct {
	ID string
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	Age int8 `json:"age"`
}

func (ud *userDomain) GetID() string {
	return ud.ID
}

func (ud *userDomain) SetID(id string) {
	ud.ID = id
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