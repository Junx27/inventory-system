package model

type Product struct {
	ID          uint              `json:"id" gorm:"primaryKey"`
	Name        string            `json:"name"`
	Price       float64           `json:"price"`
	Category    string            `json:"category"`
	Description string            `json:"description"`
	ImagePath   string            `json:"image"`
	Inventory   InventoryResponse `json:"inventory" gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE"`
	Order       []OrderResponse   `json:"orders" gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE"`
}

type ProductResponse struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Category    string  `json:"category"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
}

func (ProductResponse) TableName() string {
	return "products"
}
