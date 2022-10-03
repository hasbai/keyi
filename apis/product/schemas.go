package product

import (
	. "keyi/models"
	"keyi/utils"
)

type CreateModel struct {
	Name        string   `json:"name" validate:"required,max=32"`
	Description string   `json:"description" validate:"required,max=256"`
	Images      []string `json:"images"`
	Price       float64  `json:"price" validate:"required,min=0"`
	// -1 is to buy, 1 is to sell
	Type      ProductType `json:"type" validate:"required,oneof=-1 1"`
	Condition float64     `json:"condition" validate:"min=0,max=10"`
	Location  string      `json:"location" validate:"max=32"`
	Contact   string      `json:"contact" validate:"max=32"`
}

func (m *CreateModel) ToProduct(p *Product) error {
	return utils.Copy(m, p)
}

type ModifyModel struct {
	Name        string      `json:"name" validate:"max=32"`
	Description string      `json:"description" validate:"max=256"`
	Images      []string    `json:"images"`
	Price       float64     `json:"price" validate:"min=0"`
	Type        ProductType `json:"type"`
	Condition   float64     `json:"condition" validate:"min=0,max=10"`
	Location    string      `json:"location" validate:"max=32"`
	Contact     string      `json:"contact" validate:"max=32"`
	CategoryID  int         `json:"category_id"`
	Closed      *bool       `json:"closed"`
}

func (m *ModifyModel) ToProduct(p *Product) error {
	p.CategoryID = m.CategoryID
	return utils.Copy(m, p, "Closed")
}
