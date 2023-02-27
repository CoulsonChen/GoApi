package main

// @title Gin swagger
// @version 1.0
// @description Gin swagger
// @contact.name nathan.lu
// @contact.url https://CoulsonChen.github.io/
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8081
// schemes http
func main() {
	routing := InitHost()

	routing.SetupRouter()
	routing.SetupSwagger()
	routing.Run()
}
