package model

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