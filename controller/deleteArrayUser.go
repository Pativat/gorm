package controller

import (
	"gorm/database"
	"gorm/models/userModel"
	"log"

	"github.com/labstack/echo/v4"
)

// @Tags User
// @Summary Delete users
// @Description Delete users from the database
// @Accept json
// @Produce json
// @Param User_id path string true "ID ของผู้ใช้"
// @response 200 {object} helper.SuccessResponse "Success response"
// @Router /user/delete/{User_id} [delete]
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
