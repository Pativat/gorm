package controller

import (
	"gorm/database"
	"gorm/models/userModel"
	"log"

	"github.com/labstack/echo/v4"
)

func InsertBankArray(ctx echo.Context) {

	databank := []userModel.Bank{}
	userModelHelper := userModel.UserModelHelper{DB: database.DBMYSQL}

	if err := ctx.Bind(&databank); err != nil {
		log.Println("Error Bind")
	}

	if err := userModelHelper.InsertArrayBank(databank); err != nil {
		log.Println("Error Insert")
	}

}
