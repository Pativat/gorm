package gormQuery

import (
	"gorm/database"
	"gorm/helper"
	"gorm/models/userModel"
	"log"
	"time"
)

func CreateUser() {
	now := time.Now()
	user := userModel.User{
		Id:        helper.GenerateUUID(),
		Firstname: "Piraya",
		Lastname:  "tasha",
		Age:       21,
		CreatedAt: &now,
		UpdatedAt: &now,
	}

	userModelHelper := userModel.UserModelHelper{DB: database.DBMYSQL}
	if err := userModelHelper.Insert(&user); err != nil {
		log.Println("Error inserting user: ", err)
	}
	log.Println("Created user: ", user)
}

func GetUserById() {
	id := "58db5c65-d01d-471d-b7dc-a2c84de3fd79"
	userModelHelper := userModel.UserModelHelper{DB: database.DBMYSQL}
	user, err := userModelHelper.FindById(id)
	if err != nil {
		log.Println("Error inserting user: ", err)
	}

	log.Println("user:", user)
}

func GetAll() {
	userModelHelper := userModel.UserModelHelper{DB: database.DBMYSQL}
	user, err := userModelHelper.GetAllUsers()

	if err != nil {
		log.Println("Error")
	} else {
		for _, v := range user {
			log.Println("user  :  ", v)
		}
	}

}

func UpdatedById() {
	userModelHelper := userModel.UserModelHelper{DB: database.DBMYSQL}
	user, err := userModelHelper.UpdatedUserById()

	if err != nil {
		log.Println("Error")
	} else {

		log.Print(user)
	}

}
