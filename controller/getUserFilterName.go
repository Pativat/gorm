package controller

import (
	"gorm/database"
	"gorm/models/userModel"
	"log"
	"strconv"

	"github.com/labstack/echo/v4"
)

type FilterData struct {
	Totalpages      int `json:"totalpages"`
	Next_page       int `json:"next_page"`
	Prev_page       int `json:"prev_page"`
	Totalrow        int `json:"totalrow"`
	Total_next_page int `json:"total_next_page"`
	Total_prev_page int `json:"total_prev_page"`
}

// @Tags User
// @Summary Get user Filter Name data
// @Description Get users from the database Filter
// @Accept json
// @Produce json
// @Param fname query string false "fname"
// @Param lname query string false "lname"
// @Param page query int false "page"
// @Param limit query int false "limit"
// @response 200 {object} helper.SuccessResponse "Success response"
// @Router /user/filter/name [get]
func GetUserFilterName(ctx echo.Context) error {

	userModelHelper := userModel.UserModelHelper{DB: database.DBMYSQL}

	getlimit := ctx.QueryParam("limit")

	getpage := ctx.QueryParam("page")

	fname := ctx.QueryParam("fname")

	lname := ctx.QueryParam("lname")

	limit, _ := strconv.Atoi(getlimit)
	page, _ := strconv.Atoi(getpage)

	user, count, err := userModelHelper.FilterUsername(fname, lname, limit, page)

	totalpage := count / int64(limit)

	if err != nil {
		log.Println(err)
	}
	return ctx.JSON(200, map[string]interface{}{

		"meta": FilterData{
			Totalpages:      int(totalpage),
			Totalrow:        int(count),
			Total_next_page: int(totalpage) - page,
			Total_prev_page: page - 1,
			Prev_page:       page - 1,
			Next_page:       page + 1,
		},
		"user": user,
	})

}
