package productmodel

type Product struct {
	ProductID         int     `json:"product_id" gorm:"column:product_id"`
	ProductCategoryID int     `json:"product_category_id" gorm:"column:product_category_id"`
	ProductName       string  `json:"product_name" gorm:"column:product_name"`
	ProductPrice      float64 `json:"product_price" gorm:"column:product_price"`
}

func (Product) TableName() string {
	return "product"
}
