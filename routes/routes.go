package routes

import (
	"gorm/controller"

	"github.com/labstack/echo/v4"
)

func InitRouter(e *echo.Echo) {
	e.GET("/", controller.Hello)
	e.GET("/user", controller.GetAll)

	e.GET("/user/:limit", controller.GetLimitUser)
	e.POST("/user/insert", controller.InsertArrayUser)
	e.DELETE("/user/delete/:User_id", controller.DeleteArrayUser)

	e.PUT("/user/update/:User_id", controller.UpdatedArrayUser)

	e.GET("/user/fil", controller.GetAllUsersFilter)

}
