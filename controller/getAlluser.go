package controller

import (
	"gorm/database"
	"gorm/models/userModel"
	"log"

	"github.com/labstack/echo/v4"
)

// @Tags User
// @Summary Get all users
// @Description Get all users from the database
// @Accept json
// @Produce json
// @response 200 {object} helper.SuccessResponse "Success response"
// @Router /user [get]
func GetAll(ctx echo.Context) error {
	userModelHelper := userModel.UserModelHelper{DB: database.DBMYSQL}
	user, err := userModelHelper.GetAllUsers()

	if err != nil {
		log.Println("Error")
	}
	return ctx.JSON(200, map[string]interface{}{
		"data":    user,
		"message": "success",
	})
}
