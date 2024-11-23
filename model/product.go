package model

import "database/sql"

type Product struct {
	ID          int                 `json:"id" gorm:"primaryKey"`
	Name        string              `json:"name"`
	Price       float64             `json:"price"`
	Category    string              `json:"category"`
	Description string              `json:"description"`
	ImagePath   sql.NullString      `json:"image"`
	Inventory   []InventoryResponse `json:"inventory" gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE"`
	Order       []OrderResponse     `json:"orders" gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE"`
}

type ProductResponse struct {
	ID          int                 `json:"id"`
	Name        string              `json:"name"`
	Price       float64             `json:"price"`
	Category    string              `json:"category"`
	Description string              `json:"description"`
	ImagePath   *string             `json:"image"`
	Inventory   []InventoryResponse `json:"inventory"`
	Order       []OrderResponse     `json:"orders"`
}
type ProductResponseRelations struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Category    string  `json:"category"`
	Description string  `json:"description"`
	ImagePath   *string `json:"image"`
}

func (p *ProductResponse) FillFromModel(model Product) {
	p.ID = model.ID
	p.Name = model.Name
	p.Price = model.Price
	p.Category = model.Category
	p.Description = model.Description
	if model.ImagePath.Valid {
		p.ImagePath = &model.ImagePath.String
	} else {
		p.ImagePath = nil
	}
	p.Inventory = model.Inventory
	p.Order = model.Order
}

func (ProductResponseRelations) TableName() string {
	return "products"
}
