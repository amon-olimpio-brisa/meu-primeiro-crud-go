package controller

import (
	"github.com/amon-olimpio-brisa/meu-primeiro-crud-go/src/model/service"
	"github.com/gin-gonic/gin"
)

func NewUserControllerInterface(
	serviceInterface service.UserDomainService,
) UserControllerInterface {
	return UserControllerInterface{
		service: serviceInterface,
	}
}

type userControllerInterface interface {
	FindUserByID(c *gin.Context)
	FindUserByEmail(c *gin.Context)

	DeleteUser(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
}

type UserControllerInterface struct {
	service service.UserDomainService
}
