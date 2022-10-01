package product

import (
	. "keyi/models"
)

type CreateModel struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Images      JSON        `json:"images"`
	Price       float64     `json:"price"`
	Type        ProductType `json:"type"`
}

type ModifyModel struct {
	CreateModel
	CategoryID int `json:"category_id"`
}
