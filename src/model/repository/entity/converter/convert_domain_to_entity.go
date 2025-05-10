package converter

import (
	"github.com/rafabcanedo/api-golang/src/model"
	"github.com/rafabcanedo/api-golang/src/model/repository/entity"
)

func ConvertDomainToEntity(
	domain model.UserDomainInterface,
) *entity.UserEntity {
	return &entity.UserEntity{
		Name: domain.GetName(),
		Email: domain.GetEmail(),
		Password: domain.GetPassword(),
		Age: domain.GetAge(),
	}
}