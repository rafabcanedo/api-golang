package converter

import (
	"github.com/rafabcanedo/api-golang/src/model"
	"github.com/rafabcanedo/api-golang/src/model/repository/entity"
)

func ConvertEntityToDomain(
	entity entity.UserEntity,
) model.UserDomainInterface {
	domain := model.NewUserDomain(
		entity.Name,
		entity.Email,
		entity.Password,
		entity.Age,
	)

	domain.SetID(entity.ID.Hex())

	return domain
}