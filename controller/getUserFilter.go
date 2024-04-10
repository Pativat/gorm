package controller

import (
	"gorm/database"
	"gorm/models/userModel"
	"log"

	"github.com/labstack/echo/v4"
)

func GetAllUsersFilter(ctx echo.Context) error {

	userModelHelper := userModel.UserModelHelper{DB: database.DBMYSQL}
	filterUser := userModel.FilterData{}
	// getlimit := ctx.QueryParam("limitpage")
	// getpage := ctx.QueryParam("page")

	// getfname := ctx.QueryParam("fname")
	// getlname := ctx.QueryParam("lname")

	err := ctx.Bind(&filterUser)

	if err != nil {
		log.Println("Error by data")
	}

	user, _ := userModelHelper.GetUserFilter(filterUser)

	// if user, err := userModelHelper.GetUserFilter(getfname, getlname); user != nil && err != nil {
	// 	log.Println("Error getting user by Filter")
	// }

	return ctx.JSON(200, map[string]interface{}{
		"name":    user,
		"message": "success",
	})

}
