package models

type User struct {
	BaseModel
	Name        string    `json:"name" gorm:"size:32;not null"`
	Description string    `json:"description" gorm:"size:256"`
	Email       string    `json:"email" gorm:"size:64;not null"`
	Password    string    `json:"password" gorm:"size:64;not null"`
	Avatar      string    `json:"avatar" gorm:"size:256;not null"`
	Products    []Product `json:"products"`
}
