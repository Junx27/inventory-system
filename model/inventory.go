package model

type Inventory struct {
	ID        int                      `json:"id" gorm:"primaryKey"`
	ProductID int                      `json:"product_id"`
	Quantity  int                      `json:"qty"`
	Location  string                   `json:"location"`
	Product   ProductResponseRelations `json:"product" gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE"`
}

type InventoryResponse struct {
	ID        int    `json:"id"`
	ProductID int    `json:"product_id"`
	Quantity  int    `json:"qty"`
	Location  string `json:"location"`
}

func (InventoryResponse) TableName() string {
	return "inventories"
}
