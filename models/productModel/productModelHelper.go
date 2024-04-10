package productmodel

import "time"

type Product struct {
	Id         int        `json:"id" gorm:"id"`
	Name       string     `json:"name" gorm:"name"`
	Price      int        `json:"price" gorm:"price"`
	Status     string     `json:"status" gorm:"status"`
	Created_at *time.Time `json:"created_at " gorm:"created_at"`
	Created_by string     `json:"created_by" gorm:"created_by"`
	Updated_by *time.Time `json:"updated_by" gorm:"updated_by"`
	Deleted_at *time.Time `json:"deleted_at" gorm:"deleted_at"`
}

func (Product) TableName() string {
	return "product"
}
