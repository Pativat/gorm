package controller

import (
	"gorm/database"
	"gorm/models/userModel"
	"log"

	"github.com/labstack/echo/v4"
)

// @Tags User
// @Summary Get user Filter data
// @Description Get users from the database Filter
// @Accept json
// @Produce json
// @Param Request body userModel.FilterData true "Get user Filter data"
// @response 200 {object} helper.SuccessResponse "Success response"
// @Router /user/filter [get]
func GetAllUsersFilter(ctx echo.Context) error {

	userModelHelper := userModel.UserModelHelper{DB: database.DBMYSQL}
	filterUser := userModel.FilterData{}
	// getlimit := ctx.QueryParam("limitpage")
	// getpage := ctx.QueryParam("page")

	// getfname := ctx.QueryParam("fname")
	// getlname := ctx.QueryParam("lname")

	if err := ctx.Bind(&filterUser); err != nil {
		log.Println("Error Bind")
	}

	user, _ := userModelHelper.GetUserFilter(filterUser)

	// if user, err := userModelHelper.GetUserFilter(getfname, getlname); user != nil && err != nil {
	// 	log.Println("Error getting user by Filter")
	// }

	return ctx.JSON(200, map[string]interface{}{
		"name":    user,
		"message": "success",
	})

}
