package main

import (
	"github.com/CoulsonChen/GoApi/internal/app/controllers"
	"github.com/gin-gonic/gin"
)

func main() {

	// http api
	userController := controllers.CreateUsersController()

	r := gin.Default()
	r.GET("/users", userController.GetAllUsers)
	r.GET("/users/byfullname/:fullname", userController.GetUserByFullName)
	r.POST("/users", userController.CreateUser)

	ginerr := r.Run(":8081")
	if ginerr != nil {
		panic(ginerr)
	}
}
