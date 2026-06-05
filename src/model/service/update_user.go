package service

import (
	"github.com/amon-olimpio-brisa/meu-primeiro-crud-go/src/configuration/rest_err"
	"github.com/amon-olimpio-brisa/meu-primeiro-crud-go/src/model"
)

func (*userDomainService) UpdateUser(
	userID string,
	userDomain model.UserDomainInterface) *rest_err.RestErr {
	return nil
}
