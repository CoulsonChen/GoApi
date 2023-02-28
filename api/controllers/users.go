package controllers

import (
	"net/http"

	"github.com/CoulsonChen/GoApi/pkg/models"
	"github.com/CoulsonChen/GoApi/pkg/services"
	"github.com/gin-gonic/gin"
)

type UsersController struct {
	service    services.IUsersService
	jwtservice services.IJwtService
}

func UsersControllerProvider(s services.IUsersService,
	js services.IJwtService) *UsersController {
	return &UsersController{
		service:    s,
		jwtservice: js,
	}
}

// @Summary Get all users
// @Schemes
// @Description Get all user
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {array} string
// @Router /users [get]
func (uc *UsersController) GetAllUsers(context *gin.Context) {
	users, err := uc.service.GetAllUsers()
	if err != nil {
		panic(err)
	}

	context.JSON(http.StatusOK, users)
	return
}

// @Summary Get user by fullname
// @Schemes
// @Description Get user by fullname
// @Tags users
// @Accept json
// @Produce json
// @Param fullname path string true "user'sfullname"
// @Success 200 {array} string
// @Router /users/byfullname/{fullname} [get]
func (uc *UsersController) GetUserByFullName(context *gin.Context) {
	fullname := context.Param("fullname")
	user, err := uc.service.GetUserByFullName(fullname)
	if err != nil {
		panic(err)
	}

	context.JSON(http.StatusOK, user)
	return
}

// @Summary Create user
// @Schemes
// @Description Create user
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.User true "user's info"
// @Success 200 {array} string
// @Router /users [post]
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

// @Summary User login
// @Schemes
// @Description User login
// @Tags users
// @Accept json
// @Produce json
// @Param login body models.Login true "login info"
// @Success 200 {array} string
// @Router /users/login [post]
func (uc *UsersController) Login(context *gin.Context) {
	var login models.Login
	err := context.BindJSON(&login)
	if err != nil {
		panic(err)
	}

	isSuccess, err := uc.service.Login(login)
	if err != nil {
		panic(err)
	}

	token := ""
	if *isSuccess {
		t, err := uc.jwtservice.GenToken(login.Acct)
		if err != nil {
			panic(err)
		}
		token = t
	}

	context.JSON(http.StatusOK, token)
	return
}
