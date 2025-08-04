package converter

import (
	"github.com/guiziin227/CRUDgo/src/model"
	"github.com/guiziin227/CRUDgo/src/model/repository/entity"
)

func ConvertDomainToEntity(userDomain model.UserDomainInterface) *entity.UserEntity {
	return &entity.UserEntity{
		Email:    userDomain.GetEmail(),
		Password: userDomain.GetPassword(),
		Username: userDomain.GetUsername(),
		Age:      userDomain.GetAge(),
	}
}
