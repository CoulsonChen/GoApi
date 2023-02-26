package controllers

import (
	"net/http"

	"github.com/CoulsonChen/GoApi/internal/app/services"
	"github.com/CoulsonChen/GoApi/internal/pkg/models"
	"github.com/gin-gonic/gin"
)

type UsersController struct {
	service services.IUsersService
}

func UsersControllerProvider(s services.IUsersService) *UsersController {
	return &UsersController{
		service: s,
	}
}

func (uc *UsersController) GetAllUsers(context *gin.Context) {
	users, err := uc.service.GetAllUsers()
	if err != nil {
		panic(err)
	}

	context.JSON(http.StatusOK, users)
	return
}

func (uc *UsersController) GetUserByFullName(context *gin.Context) {
	fullname := context.Param("fullname")
	user, err := uc.service.GetUserByFullName(fullname)
	if err != nil {
		panic(err)
	}

	context.JSON(http.StatusOK, user)
	return
}

func (uc *UsersController) CreateUser(context *gin.Context) {
	var user models.User
	err := context.BindJSON(&user)
	if err != nil {
		panic(err)
	}

	useracct, err := uc.service.CreateUser(user)
	if err != nil {
		panic(err)
	}

	context.JSON(http.StatusOK, useracct)
	return
}
