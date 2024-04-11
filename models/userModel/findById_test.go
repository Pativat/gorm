package userModel_test

import (
	"fmt"
	"gorm/database"
	"gorm/models/userModel"
	"testing"
)

func TestFindBy(t *testing.T) {
	t.Run("findBy", func(t *testing.T) {
		db := database.DBMYSQL
		userRepo := userModel.NewUserModelHelper(db)
		result, err := userRepo.FindById("0b908c89-45f3-41ed-aeb6-d4d7baa8928c")

		if err != nil {
			fmt.Printf("err: %v\n", err)
		}
		fmt.Printf("user: %v\n", result)
	})
}
