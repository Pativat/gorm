package controller

import (
	"gorm/database"
	"gorm/models/userModel"
	"log"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetLimitUser(ctx echo.Context) error {

	getlimit := ctx.Param("limit")

	getlimit1, _ := strconv.Atoi(getlimit)

	userModelHelper := userModel.UserModelHelper{DB: database.DBMYSQL}

	user, err := userModelHelper.GetUserLimit(getlimit1)

	if err != nil {
		log.Println("Cannot Get User")
	}

	count, err := userModelHelper.GetCountUser(user)
	if err != nil {
		log.Println("Cannot get Count User")
	}
	return ctx.JSON(200, map[string]interface{}{
		"name":    user,
		"count":   count,
		"message": "success",
	})
}
