package product

import (
	. "keyi/models"
)

type CreateModel struct {
	Name        string      `json:"name" validate:"required,max=32"`
	Description string      `json:"description" validate:"max=256"`
	Images      StringArray `json:"images"`
	// Price in cent, $2.70 = 270
	Price    int         `json:"price" validate:"required,min=0"`
	Type     ProductType `json:"type" validate:"required,oneof=-1 1"`
	TenantID int         `json:"tenant_id" validate:"required"`
}

type ModifyModel struct {
	Name        string      `json:"name" validate:"max=32"`
	Description string      `json:"description" validate:"max=256"`
	Images      StringArray `json:"images"`
	Price       int         `json:"price" validate:"min=0"`
	Type        ProductType `json:"type"`
	TenantID    int         `json:"tenant_id"`
	CategoryID  int         `json:"category_id"`
	Closed      *bool       `json:"closed"`
}

type ListUserProductsQuery struct {
	Query
	// 0: all, 1: not closed, -1: closed
	Closed int8 `query:"closed" validate:"min=-1,max=1"`
	// 0: all, 1: sell, -1: buy
	Type int8 `query:"type" validate:"min=-1,max=1"`
}
