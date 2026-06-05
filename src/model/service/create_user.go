package service

import (
	"fmt"

	"github.com/amon-olimpio-brisa/meu-primeiro-crud-go/src/configuration/logger"
	"github.com/amon-olimpio-brisa/meu-primeiro-crud-go/src/configuration/rest_err"
	"github.com/amon-olimpio-brisa/meu-primeiro-crud-go/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(
	userDomain model.UserDomainInterface,
) *rest_err.RestErr {

	logger.Info("Init createUser model", zap.String("journey", "createUser"))

	userDomain.EncryptPassword()

	fmt.Println(ud)

	return nil
}
