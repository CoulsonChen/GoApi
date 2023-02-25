package controllers

import (
	"net/http"

	"github.com/CoulsonChen/GoApi/internal/app/services"
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
