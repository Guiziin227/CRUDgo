package view

import (
	"github.com/guiziin227/CRUDgo/src/controller/model/response"
	"github.com/guiziin227/CRUDgo/src/model"
)

func ConvertDomainToResponse(domainInterface model.UserDomainInterface) response.UserResponse {

	return response.UserResponse{
		Id:       "id",
		Email:    domainInterface.GetEmail(),
		Username: domainInterface.GetUsername(),
		Age:      domainInterface.GetAge(),
	}
}
