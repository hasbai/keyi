package product

import (
	. "keyi/models"
)

type CreateModel struct {
	Name        string      `json:"name" validate:"required,max=32"`
	Description string      `json:"description" validate:"max=256"`
	Images      JSON        `json:"images"`
	Price       float64     `json:"price" validate:"required,min=0"`
	Type        ProductType `json:"type" validate:"required,oneof=-1 1"`
	TenantID    int         `json:"tenant_id" validate:"required"`
}

type ModifyModel struct {
	Name        string      `json:"name" validate:"max=32"`
	Description string      `json:"description" validate:"max=256"`
	Images      JSON        `json:"images"`
	Price       float64     `json:"price" validate:"min=0"`
	Type        ProductType `json:"type" validate:"oneof=-1 1"`
	TenantID    int         `json:"tenant_id"`
	CategoryID  int         `json:"category_id"`
	Closed      bool        `json:"closed"`
}
