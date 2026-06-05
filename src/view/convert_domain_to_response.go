package view

import (
	"github.com/amon-olimpio-brisa/meu-primeiro-crud-go/src/controller/model/response"
	"github.com/amon-olimpio-brisa/meu-primeiro-crud-go/src/model"
)

func ConverseDomainToResponse(
	userDomain model.UserDomainInterface,
) response.UserResponse {
	return response.UserResponse{
		ID:    "",
		Email: userDomain.GetEmail(),
		Name:  userDomain.GetName(),
		Age:   userDomain.GetAge(),
	}
}
