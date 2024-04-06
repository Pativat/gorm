package controller

import (
	"gorm/database"
	"gorm/models/userModel"
	"log"

	"github.com/labstack/echo/v4"
)

func GetAll(ctx echo.Context) error {

	userModelHelper := userModel.UserModelHelper{DB: database.DBMYSQL}
	user, err := userModelHelper.GetAllUsers()

	if err != nil {
		log.Println("Error")
	} else {
		for _, v := range user {
			log.Println("user  :  ", v)
		}
	}
	return ctx.JSON(200, map[string]interface{}{
		"data":    user,
		"message": "success",
	})

}
