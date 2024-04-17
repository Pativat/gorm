package main

import (
	"gorm/configs"
	"gorm/database"
	"gorm/routes"

	_ "gorm/docs"

	echoSwagger "github.com/swaggo/echo-swagger"

	"github.com/labstack/echo/v4"
)

//	@title			Your-Project Document
//	@version		1.0
//	@description	This is a sample server.
//	@termsOfService	http://swagger.io/terms/
//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io
//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html
//	@host			localhost:1323
//	@BasePath		/
//	@schemes		https

func main() {

	e := echo.New()

	routes.InitRouter(e)
	configs.Init()
	database.Init()
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	//gormQuery.CreateUser()
	//gormQuery.GetAll()

	//gormQuery.UpdatedById()

	e.Logger.Fatal(e.Start(":1323"))

}
