package userModel

import "gorm.io/gorm"

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

func (u *UserModelHelper) GetUserLimit(limit, off int) ([]*User, error) {

	user := []*User{}

	result := u.DB.Limit(limit).Offset(off).Find(&user)

	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

// func (u *UserModelHelper) DeleteUserById() (*User, error) {

// }
