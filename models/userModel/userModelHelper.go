package userModel

import (
	"log"

	"gorm.io/gorm"
)

type UserModelHelper struct {
	DB *gorm.DB
}

func (u *UserModelHelper) Insert(user *User) error {
	result := u.DB.Debug().Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *UserModelHelper) FindById(id string) (*User, error) {
	user := User{}
	result := u.DB.First(&user, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (u *UserModelHelper) GetAllUsers() ([]*User, error) {
	user := []*User{}
	result := u.DB.Find(&user)

	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (u *UserModelHelper) UpdatedUserById() (*User, error) {

	user := User{Firstname: "piya", Lastname: "tacha"}
	result := u.DB.Table("user").Where("id = ?", "cd889de3-7acb-4a65-85d5-1618ef0318ae").Updates(&user)

	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (u *UserModelHelper) GetUserLimit(limit int) ([]*User, error) {

	user := []*User{}

	result := u.DB.Limit(limit).Offset(0).Find(&user)

	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (u *UserModelHelper) InsertArray(users []User) error {
	tx := u.DB.Begin()

	if err := tx.Create(&users).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (u *UserModelHelper) DeletedArray(User_id string) ([]*User, error) {

	user := []*User{}
	tx := u.DB.Begin()

	result := tx.Debug().Where("id = ?", User_id).Delete(user)

	if result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}
	tx.Commit()
	return user, nil
}

func (u *UserModelHelper) UpdatedArray(User_id string, users []User) ([]*User, error) {

	var updatedUsers []*User
	tx := u.DB.Begin()

	for _, user := range users {

		if result := tx.Debug().Model(&User{}).Where("id = ?", User_id).Updates(&user); result.Error != nil {

			log.Println("Error updating user")
			tx.Rollback()
			return nil, result.Error

		}
		updatedUsers = append(updatedUsers, &user)
	}
	tx.Commit()
	return updatedUsers, nil
}

func (u *UserModelHelper) GetUserFilter(filterData FilterData) ([]*User, error) {
	user := []*User{}
	tx := u.DB.Begin()
	result := tx.Debug().Where("firstname LIKE ? AND lastname LIKE ?", "%"+filterData.Firstname+"%", "%"+filterData.Lastname+"%").Limit(filterData.Limit).Find(&user)

	if result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}
	tx.Commit()
	return user, nil
}
