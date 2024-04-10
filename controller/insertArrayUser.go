package controller

import (
	"gorm/database"
	"gorm/helper"
	"gorm/models/userModel"
	"log"
	"time"

	"github.com/labstack/echo/v4"
)

func InsertArrayUser(ctx echo.Context) error {
	userModelHelper := userModel.UserModelHelper{DB: database.DBMYSQL}
	now := time.Now()

	data := []userModel.User{}

	if err := ctx.Bind(&data); err != nil {
		log.Println("Error bind data:", err)
		return err
	}

	users := []userModel.User{}

	for _, user := range data {
		user.Id = helper.GenerateUUID()
		user.CreatedAt = &now
		user.UpdatedAt = &now

		//users[index] = user
		users = append(users, user)
	}

	if err := userModelHelper.InsertArray(users); err != nil {
		log.Println("Error insert users:", err)
		return err
	}

	return ctx.JSON(200, map[string]interface{}{
		"users":   users,
		"message": "success",
	})
}

// func InsertArrayUser2(ctx echo.Context) error {
// 	userModel := userModel.UserModelHelper{DB: database.DBMYSQL}

// 	return nil
// }
