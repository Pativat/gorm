package controller

import (
	"gorm/database"
	"gorm/models/userModel"
	"log"

	"github.com/labstack/echo/v4"
)

func GetAllOrder(ctx echo.Context) error {
	Filuser := []userModel.UserFill{}

	if err := ctx.Bind(&Filuser); err != nil {
		log.Println("Error Bind")
		return err
	}

	userModelHelper := userModel.UserModelHelper{DB: database.DBMYSQL}

	user, count, err := userModelHelper.GetUserOrder(Filuser)
	if err != nil {
		log.Println("Error retrieving user orders:", err)
		return err
	}
	OrderNew := make([]userModel.NewOrder, len(user))

	for idx, i := range user {
		newOrder := userModel.NewOrder{
			Id:         i.Id,
			User_id:    i.User_id,
			Firstname:  i.Firstname,
			Quantity:   i.Quantity,
			Name:       i.Name,
			Order_date: i.Order_date,
		}
		OrderNew[idx] = newOrder
	}
	// for _, i := range user {
	// 	newOrder := userModel.OrderList{
	// 		Id:         i.Id,
	// 		Firstname:  i.Firstname,
	// 		Name:       i.Name,
	// 		Order_date: i.Order_date,
	// 		Quantity:   i.Quantity,
	// 		User_id:    i.User_id,
	// 	}
	// 	OrderNew = append(OrderNew, newOrder)
	// }

	// OrderNew := make([]userModel.NewOrder, len(user))

	// for idx, u := range user {

	// 	OrderNew[idx] = userModel.NewOrder{
	// 		Id:         u.Id,
	// 		Firstname:  u.Firstname,
	// 		Name:       u.Name,
	// 		Order_date: u.Order_date,
	// 		Quantity:   u.Quantity,
	// 		User_id:    u.User_id,
	// 	}
	// }

	return ctx.JSON(200, map[string]interface{}{
		"Order": user,
		"total": count,
	})
}
