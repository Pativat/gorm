package controller

import (
	"gorm/database"
	"gorm/helper"
	"gorm/models/userModel"
	"log"
	"time"

	"github.com/labstack/echo/v4"
)

// @Tags User
// @Summary Insert Array User
// @Description Insert all users from the database
// @Accept json
// @Produce json
// @Param Request body []userModel.UserInsert true "Array User to insert"
// @Response 200 {object} helper.SuccessResponse "Success response insert"
// @Router /user/insert [post]
func InsertArrayUser(ctx echo.Context) error {
	userModelHelper := userModel.UserModelHelper{DB: database.DBMYSQL}
	now := time.Now()

	data := []userModel.UserInsert{}

	if err := ctx.Bind(&data); err != nil {
		log.Println("Error bind data:", err)
		return err
	}

	users := []userModel.User{}

	for _, userData := range data {
		user := userModel.User{
			Id:        helper.GenerateUUID(),
			Firstname: userData.Firstname,
			Lastname:  userData.Lastname,
			Age:       userData.Age,
			CreatedAt: &now,
			UpdatedAt: &now,
			DeletedAt: nil,
			Status:    userData.Status,
		}

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
