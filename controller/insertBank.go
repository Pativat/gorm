package controller

import (
	"gorm/database"
	"gorm/helper"
	"gorm/models/userModel"
	"log"

	"github.com/labstack/echo/v4"
)

// @Tags Bank
// @Summary Insert Array Bank
// @Description Insert all users from the database
// @Accept json
// @Produce json
// @Param Request body []userModel.BankInsert true "Array User to insert"
// @Response 200 {object} helper.SuccessResponse "Success response insert"
// @Router /user/insert [post]
func InsertBankArray(ctx echo.Context) error {

	databank := []userModel.BankInsert{}
	userModelHelper := userModel.UserModelHelper{DB: database.DBMYSQL}

	if err := ctx.Bind(&databank); err != nil {
		log.Println("Error Bind")
	}
	bank := []userModel.Bank{}

	for _, bankData := range databank {
		bankinfo := userModel.Bank{
			Id:           helper.GenerateUUID(),
			Bank_name:    bankData.Bank_name,
			Bank_address: bankData.Bank_address,
			Bank_image:   bankData.Bank_image,
			Bank_type:    bankData.Bank_type,
			Bank_status:  bankData.Bank_status,
		}
		bank = append(bank, bankinfo)
	}

	if err := userModelHelper.InsertArrayBank(bank); err != nil {
		log.Println("Error Insert")
	}
	return ctx.JSON(200, map[string]interface{}{
		"bank":    bank,
		"message": "success",
	})

}
