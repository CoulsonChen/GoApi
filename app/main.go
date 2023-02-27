package main

import (
	"fmt"

	"github.com/CoulsonChen/GoApi/app/controllers"
	"github.com/CoulsonChen/GoApi/app/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Gin swagger
// @version 1.0
// @description Gin swagger
// @contact.name nathan.lu
// @contact.url https://CoulsonChen.github.io/
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8083
// schemes http
func main() {

	// init controller
	userController := controllers.CreateUsersController()
	// http api
	r := gin.Default()
	r.GET("/users", userController.GetAllUsers)
	r.GET("/users/byfullname/:fullname", userController.GetUserByFullName)
	r.POST("/users", userController.CreateUser)

	// swagger
	port := 8083
	if mode := gin.Mode(); mode == gin.DebugMode {
		docs.SwaggerInfo.BasePath = "/"
		url := ginSwagger.URL(fmt.Sprintf("http://localhost:%d/swagger/doc.json", port)) // The url pointing to API definition
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	}
	// listen port and run
	ginerr := r.Run(fmt.Sprintf(":%d", port))
	if ginerr != nil {
		panic(ginerr)
	}
}
