package controller

import (
	"gorm/database"
	"gorm/models/userModel"
	"log"
	"time"

	"github.com/labstack/echo/v4"
)

// @Tags User
// @Summary Update users by Id
// @Description Update users from the database
// @Accept json
// @Produce json
// @Param User_id path string true "ID ของผู้ใช้"
// @Param Request body []userModel.UserDelete true "Delete User to insert"
// @response 200 {object} helper.SuccessResponse "Success response"
// @Router /user/update/{User_id} [put]
func UpdatedArrayUser(ctx echo.Context) error {
	userModelHelper := userModel.UserModelHelper{DB: database.DBMYSQL}
	now := time.Now()
	data := []userModel.UserDelete{}

	getid := ctx.Param("User_id")

	if err := ctx.Bind(&data); err != nil {
		log.Println("Error Bind")
		return err
	}

	users := []userModel.User{}

	for _, userdata := range data {
		user := userModel.User{

			Firstname: userdata.Firstname,
			Lastname:  userdata.Lastname,
			Age:       userdata.Age,
			UpdatedAt: &now,
			DeletedAt: nil,
			Status:    userdata.Status,
		}
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
