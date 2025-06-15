package view

import (
	"github.com/rafabcanedo/api-golang/src/controller/model/response"
	"github.com/rafabcanedo/api-golang/src/model"
)

func ConvertDomainToResponse(
	userDomain model.UserDomainInterface,
) response.UserResponse{
	return response.UserResponse{
	ID: userDomain.GetID(),
	Name: userDomain.GetName(),
	Email: userDomain.GetEmail(),
	Age: userDomain.GetAge(),
	}
}

func ConvertDomainListToResponse(usersDomain []model.UserDomainInterface) []response.UserResponse {
	var usersResponse []response.UserResponse
	for _, user := range usersDomain {
		usersResponse = append(usersResponse, ConvertDomainToResponse(user))
	}
	return usersResponse
}