package controller

import (
	"gorm/database"
	"gorm/models/userModel"
	"log"
	"strconv"

	"github.com/labstack/echo/v4"
)

type FillGetUser struct {
	NextPage        int `json:"next_page"`
	PrevPage        int `json:"prev_page"`
	TotalPages      int `json:"total_pages"`
	TotalRows       int `json:"total_rows"`
	Total_next_page int `json:"total_next_page"`
	Total_prev_page int `json:"total_prev_page"`
}

// @Tags User
// @Summary Get users Limit Offset
// @Description Get users from the database limit
// @Accept json
// @Produce json
// @Param limit query int false "limit"
// @Param page query int false "page"
// @response 200 {object} helper.SuccessResponse "Success response"
// @Router /users/p/ [get]
func GetUserp(ctx echo.Context) error {
	userModelHelper := userModel.UserModelHelper{DB: database.DBMYSQL}
	limitstr := ctx.QueryParam("limit")
	pageStr := ctx.QueryParam("page")
	page, _ := strconv.Atoi(pageStr)
	limit, _ := strconv.Atoi(limitstr)

	users, count, err := userModelHelper.GetAllUsersP(limit, page)

	if err != nil {
		log.Println("Error")
	}

	totalPages := int(count / int64(limit))

	// user, err := userModelHelper.GetUserLimit(page)

	// if err != nil {
	// 	log.Println("Error")
	// }

	return ctx.JSON(200, map[string]interface{}{
		"name": users,
		"meta": FillGetUser{
			NextPage:        page + 1,
			PrevPage:        page - 1,
			TotalPages:      totalPages,
			Total_next_page: totalPages - page,
			Total_prev_page: page - 1,
			TotalRows:       int(count),
		},
	})
}
