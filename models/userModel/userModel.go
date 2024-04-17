package userModel

import "time"

type User struct {
	Id        string     `json:"id" gorm:"id"`
	Firstname string     `json:"firstname" gorm:"column:firstname"`
	Lastname  string     `json:"lastname" gorm:"column:lastname"`
	Age       int        `json:"age" gorm:"column:age"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"column:deleted_at"`
	Status    string     `json:"status" gorm:"column:status"`
}

func (User) TableName() string {
	return "user"
}

type UserInsert struct {
	Firstname string `json:"firstname" example:"phat"`
	Lastname  string `json:"lastname" example:"cha"`
	Age       int    `json:"age" example:"21"`
	Status    string `json:"status" example:"active"`
}

type UserDelete struct {
	Firstname string `json:"firstname" example:"phat2"`
	Lastname  string `json:"lastname" example:"cha2"`
	Age       int    `json:"age" example:"22"`
	Status    string `json:"status" example:"unactive"`
}

type FilterData struct {
	Limit     int    `json:"limit" gorm:"limit" query:"limit"`
	Page      int    `json:"page"  gorm:"column:page" query:"page"`
	Firstname string `json:"firstname" gorm:"column:firstname" query:"firstname"`
	Lastname  string `json:"lastname" gorm:"lastname" query:"lastname"`
}

//	type Userfunc interface {
//		GetUserById(Id string) (User, error)
//	}

type UserFill struct {
	Id         string     `json:"id" `
	User_id    string     `json:"user_id"`
	Firstname  string     `json:"firstname"`
	Quantity   int        `json:"quantity"`
	Name       string     `json:"name"`
	Order_date *time.Time `json:"order_date"`
}

type UserList struct {
	Id        string `json:"id"`
	Firstname string `json:"firstname"`
}
type OrderList struct {
	Id         string     `json:"id"`
	User_id    string     `json:"user_id"`
	Firstname  string     `json:"first_name"`
	Quantity   int        `json:"quantity"`
	Name       string     `json:"product_name"`
	Order_date *time.Time `json:"order_date"`
}

type NewOrder struct {
	Id         string     `json:"id"`
	User_id    string     `json:"user_id"`
	Firstname  string     `json:"user_name"`
	Quantity   int        `json:"quantity"`
	Name       string     `json:"product_name"`
	Order_date *time.Time `json:"order_date"`
}

type Userinter interface {
	GetUserOrder(UserFill) ([]User, int, error)
}

type Bank struct {
	Id           string `json:"id" gorm:"primaryKey;autoIncrement;not null" `
	Bank_name    string `json:"bank_name" gorm:"bank_name"`
	Bank_address string `json:"bank_address" gorm:"bank_address"`
	Bank_image   string `json:"bank_image" gorm:"bank_image"`
	Bank_type    string `json:"bank_type" gorm:"bank_type"`
	Bank_status  string `json:"bank_status" gorm:"bank_status"`
}

func (Bank) TableName() string {
	return "bank"
}
