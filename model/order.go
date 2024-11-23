package model

type Order struct {
	ID        int             `json:"id" gorm:"primaryKey"`
	ProductID int             `json:"product_id"` // Tipe data sesuai dengan ProductID di Product
	Quantity  int             `json:"qty"`
	Product   ProductResponse `json:"products" gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE"` // Pastikan foreignKey sesuai
}

type OrderResponse struct {
	ID        int `json:"id" gorm:"primaryKey"`
	ProductID int `json:"-"` // Tipe data sesuai dengan ProductID di Product
	Quantity  int `json:"qty"`
}

func (OrderResponse) TableName() string {
	return "orders"
}
