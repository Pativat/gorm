package controller

import (
	"gorm/database"
	"gorm/models/userModel"
	"log"
	"time"

	"github.com/labstack/echo/v4"
)

func UpdatedArrayUser(ctx echo.Context) error {
	userModelHelper := userModel.UserModelHelper{DB: database.DBMYSQL}
	now := time.Now()
	data := []userModel.User{}

	getid := ctx.Param("User_id")

	if err := ctx.Bind(&data); err != nil {
		log.Println("Error Bind")
		return err
	}

	users := []userModel.User{}

	for _, user := range data {
		user.UpdatedAt = &now
		users = append(users, user)
	}

	if err, use := userModelHelper.UpdatedArray(getid, users); err != nil && use != nil {
		log.Println("Error Update")
		return ctx.JSON(400, map[string]interface{}{
			"name":    users,
			"message": "Error update user",
		})

	}

	return ctx.JSON(200, map[string]interface{}{
		"name":    users,
		"message": "successfully updated",
	})
}
