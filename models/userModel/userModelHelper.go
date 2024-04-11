package userModel

import (
	"log"

	"gorm.io/gorm"
)

type UserModelHelper struct {
	DB *gorm.DB
}

func NewUserModelHelper(db *gorm.DB) *UserModelHelper {
	return &UserModelHelper{db}
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

func (u *UserModelHelper) GetUserLimit(limit int) ([]User, error) {

	user := []User{}

	result := u.DB.Debug().Limit(limit).Offset(0).Find(&user)

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

func (u *UserModelHelper) UpdatedArray(User_id string, users []User) ([]User, error) {

	var updatedUsers []User
	tx := u.DB.Begin()

	for _, user := range users {

		if result := tx.Debug().Model(&User{}).Where("id = ?", User_id).Updates(&user); result.Error != nil {

			log.Println("Error updating user")
			tx.Rollback()
			return nil, result.Error

		}
		updatedUsers = append(updatedUsers, user)
	}
	tx.Commit()
	return updatedUsers, nil
}

func (u *UserModelHelper) GetUserFilter(filterData FilterData) ([]*User, error) {
	user := []*User{}
	tx := u.DB.Begin()
	result := tx.Debug().Where("firstname LIKE ? AND lastname LIKE ?", "%"+filterData.Firstname+"%", "%"+filterData.Lastname+"%").Limit(filterData.Limit).Offset(filterData.Page).Find(&user)

	if result.Error != nil {
		tx.Rollback()
		return nil, result.Error
	}
	tx.Commit()
	return user, nil
}

func (u *UserModelHelper) GetCountUser([]User) (int64, error) {

	user := []User{}
	var c int64
	tx := u.DB.Begin()
	result := tx.Debug().Model(user).Count(&c)

	if result.Error != nil {
		return 0, result.Error
	}
	return c, nil

}

func (u *UserModelHelper) GetUserByPage(page int) ([]*User, error) {

	user := []*User{}
	tx := u.DB.Begin()

	result := tx.Limit(10).Offset((page - 1) * 10).Find(&user)

	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil

}

func (u *UserModelHelper) GetAllUsersP(limit, page int) ([]User, int64, error) {
	var users []User
	var count int64

	tx := u.DB.Begin()

	if limit == 0 {
		limit = 10
	}

	offset := (page - 1) * limit
	err := tx.Debug().Table("user").Limit(limit).Offset(offset).Find(&users).Error
	if err != nil {
		return nil, 0, err
	}

	err = u.DB.Model(&User{}).Count(&count).Error
	if err != nil {
		return nil, 0, err
	}

	return users, count, nil
}

func (u *UserModelHelper) GetUserOrder([]UserFill) ([]OrderList, int, error) {

	var order []OrderList

	var count int64

	pipeline := "SELECT `order`.id ,user.firstname ,product_order.quantity ,product.name ,  user.id  AS user_id , order_date FROM user inner join `order` on user.id = `order`.user_id inner join product_order on  `order`.id = product_order.order_id inner join product on product_order.product_id = product.id"

	if err := u.DB.Debug().Raw(pipeline).Find(&order).Error; err != nil {
		return nil, 0, err
	}

	if err := u.DB.Model(&User{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	return order, int(count), nil

}

func (u *UserModelHelper) CreateTableBank() error {

	newbank := []Bank{}
	err := u.DB.AutoMigrate(newbank)

	if err != nil {
		return err
	}
	return nil

}

func (u *UserModelHelper) InsertArrayBank([]Bank) error {

	bank := []Bank{}
	tx := u.DB.Begin()

	if err := tx.Create(&bank).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil

}
