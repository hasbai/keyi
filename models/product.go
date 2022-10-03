package models

type Product struct {
	BaseModel
	Name        string      `json:"name" gorm:"size:32;not null"`
	Description string      `json:"description" gorm:"size:256;not null"`
	Images      StringSlice `json:"images"  gorm:"not null"`
	Price       float64     `json:"price" gorm:"not null"`
	Type        ProductType `json:"type" gorm:"not null"`
	Condition   float64     `json:"condition" gorm:"not null"` // 0 - 10
	Location    string      `json:"location" gorm:"size:32;not null"`
	Contact     string      `json:"contact" gorm:"size:32;not null"`
	Closed      bool        `json:"closed" gorm:"not null"`
	UserID      int         `json:"user_id" gorm:"not null"`
	CategoryID  int         `json:"category_id" gorm:"not null"`
}

type ProductType = int8

const (
	ProductTypeBid ProductType = -1 // buy
	ProductTypeAsk ProductType = 1  // sell
)
