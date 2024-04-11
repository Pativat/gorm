package controller

import (
	"gorm/database"
	"gorm/models/userModel"
	"log"

	"github.com/labstack/echo/v4"
)

func CreateTableBank(ctx echo.Context) error {

	userModelHelper := userModel.UserModelHelper{DB: database.DBMYSQL}

	if err := userModelHelper.CreateTableBank(); err != nil {
		log.Println("Error creating")
	}

	return ctx.JSON(200, map[string]interface{}{

		"message": "success",
	})
}
