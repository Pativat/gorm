package routes

import (
	"gorm/controller"

	"github.com/labstack/echo/v4"
)

func InitRouter(e *echo.Echo) {
	e.POST("/user/insert", controller.InsertArrayUser)
	e.DELETE("/user/delete/:User_id", controller.DeleteArrayUser)
	e.GET("/user", controller.GetAll)

	e.GET("/users/:limit", controller.GetLimitUser)

	e.PUT("/user/update/:User_id", controller.UpdatedArrayUser)

	e.GET("/users/filter", controller.GetAllUsersFilter)
	e.GET("/users/p/", controller.GetUserp)

	e.GET("/users/order/", controller.GetAllOrder)

	e.POST("/bank/create", controller.CreateTableBank)

}
