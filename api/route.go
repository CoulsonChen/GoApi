package route

import (
	"fmt"

	"github.com/CoulsonChen/GoApi/api/controllers"
	"github.com/CoulsonChen/GoApi/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Route struct {
	port            int
	router          *gin.Engine
	userscontroller *controllers.UsersController
}

func RouteProvider(userscontroller *controllers.UsersController) *Route {
	return &Route{
		port:            8081,
		router:          gin.Default(),
		userscontroller: userscontroller,
	}
}

func (r *Route) SetupRouter() {
	r.router.GET("/users", r.userscontroller.GetAllUsers)
	r.router.GET("/users/byfullname/:fullname", r.userscontroller.GetUserByFullName)
	r.router.POST("/users", r.userscontroller.CreateUser)
}

func (r *Route) SetupSwagger() {
	if mode := gin.Mode(); mode == gin.DebugMode {
		docs.SwaggerInfo.BasePath = "/"
		url := ginSwagger.URL(fmt.Sprintf("http://localhost:%d/swagger/doc.json", r.port)) // The url pointing to API definition
		r.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	}
}

func (r *Route) Run() {
	ginerr := r.router.Run(fmt.Sprintf(":%d", r.port))
	if ginerr != nil {
		panic(ginerr)
	}
}