package converter

import (
	"github.com/guiziin227/CRUDgo/src/model"
	"github.com/guiziin227/CRUDgo/src/model/repository/entity"
)

func ConvertEntityToDomain(userEntity *entity.UserEntity) model.UserDomainInterface {
	domain := model.NewUserDomain(
		userEntity.Email,
		userEntity.Password,
		userEntity.Username,
		userEntity.Age,
	)

	domain.SetID(userEntity.ID.Hex())
	return domain
}
