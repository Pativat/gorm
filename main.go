package main

import (
	"gorm/configs"
	"gorm/database"
	"gorm/routes"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()

	routes.InitRouter(e)
	configs.Init()
	database.Init()

	//gormQuery.CreateUser()
	//gormQuery.GetAll()

	//gormQuery.UpdatedById()

	e.Logger.Fatal(e.Start(":1323"))

}
