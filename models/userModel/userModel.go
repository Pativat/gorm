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

type FilterData struct {
	Limit     int    `json:"limit" gorm:"limit" query:"limit"`
	Page      int    `json:"page"  gorm:"column:page" query:"page"`
	Firstname string `json:"firstname" gorm:"column:firstname" query:"firstname"`
	Lastname  string `json:"lastname" gorm:"lastname" query:"lastname"`
}
