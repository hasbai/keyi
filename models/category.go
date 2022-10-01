package models

type Category struct {
	BaseModel
	Name        string    `json:"name" gorm:"size:32;not null;unique"`
	Description string    `json:"description" gorm:"size:256;not null"`
	Products    []Product `json:"products"`
}
