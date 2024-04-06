package routes

import (
	"gorm/controller"

	"github.com/labstack/echo/v4"
)

func InitRouter(e *echo.Echo) {
	e.GET("/", controller.Hello)
	e.GET("/user", controller.GetAll)

	e.GET("/user/:limit/:off", controller.GetLimitUser)

}
