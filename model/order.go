package model

type Order struct {
	ID        int                      `json:"id" gorm:"primaryKey"`
	ProductID int                      `json:"product_id"`
	Quantity  int                      `json:"qty"`
	DateOrder string                   `json:"date_order"`
	Product   ProductResponseRelations `json:"products" gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE"`
}

type OrderResponse struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	ProductID int    `json:"product_id"`
	Quantity  int    `json:"qty"`
	DateOrder string `json:"date_order"`
}

func (OrderResponse) TableName() string {
	return "orders"
}
