package controller

import (
	"gorm/database"
	"gorm/models/userModel"
	"log"

	"github.com/labstack/echo/v4"
)

func DeleteArrayUser(ctx echo.Context) error {

	Id_user := ctx.Param("User_id")

	userModelHelper := userModel.UserModelHelper{DB: database.DBMYSQL}
	user, err := userModelHelper.DeletedArray(Id_user)

	if err != nil {
		log.Println("Error Delete array")
	}

	return ctx.JSON(200, map[string]interface{}{
		"name":    user,
		"message": "success",
	})

}
