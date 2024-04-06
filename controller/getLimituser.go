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

	getoff := ctx.Param("off")

	getoff1, _ := strconv.Atoi(getoff)

	getlimit1, _ := strconv.Atoi(getlimit)

	userModelHelper := userModel.UserModelHelper{DB: database.DBMYSQL}

	user, err := userModelHelper.GetUserLimit(getlimit1, getoff1)

	if err != nil {
		log.Println("Error getting user limit")
	} else {

		for _, v := range user {

			log.Println("User :", v)
		}

	}

	return ctx.JSON(200, map[string]interface{}{
		"name":    user,
		"message": "success",
	})
}
