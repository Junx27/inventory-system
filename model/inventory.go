package model

type Inventory struct {
	ID        int             `json:"id" gorm:"primaryKey"`
	ProductID int             `json:"product_id"`
	Quantity  int             `json:"qty"`
	Location  string          `json:"location"`
	Product   ProductResponse `json:"product" gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE"` // Hindari preload relasi Product di dalam Inventory
}

type InventoryResponse struct {
	ID        int    `json:"id"`
	ProductID int    `json:"-"`
	Quantity  int    `json:"qty"`
	Location  string `json:"location"`
}

func (InventoryResponse) TableName() string {
	return "inventories"
}
