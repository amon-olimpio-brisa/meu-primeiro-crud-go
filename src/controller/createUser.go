package controller

import (
	"fmt"
	"net/http"

	"github.com/amon-olimpio-brisa/meu-primeiro-crud-go/src/configuration/logger"
	"github.com/amon-olimpio-brisa/meu-primeiro-crud-go/src/configuration/rest_err"
	"github.com/amon-olimpio-brisa/meu-primeiro-crud-go/src/controller/model/request"
	"github.com/amon-olimpio-brisa/meu-primeiro-crud-go/src/model"
	"github.com/amon-olimpio-brisa/meu-primeiro-crud-go/src/view"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap/zapcore"
)

var (
	userDomainInterface model.UserDomainInterface
)

func (uc *UserControllerInterface) CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller",
		zapcore.Field{
			Key:    "journey",
			String: "createUser",
		},
	)
	var userResquest request.UserResquest

	if err := c.ShouldBindJSON(&userResquest); err != nil {
		logger.Error("Error trying to validate user info", err)
		restErr := rest_err.NewBadRequestError(
			fmt.Sprintf("There are some incorrects fields, error=%s", err.Error()))
		c.JSON(restErr.Code, restErr)
		return
	}

	domain := model.NewUserDomain(
		userResquest.Email,
		userResquest.Password,
		userResquest.Name,
		userResquest.Age,
	)

	if err := uc.service.CreateUser(domain); err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User created successfully",
		zapcore.Field{
			Key:    "journey",
			String: "createUser",
		})

	c.JSON(http.StatusOK, view.ConverseDomainToResponse(
		domain,
	))

}
