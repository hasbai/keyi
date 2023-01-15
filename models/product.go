package models

import "gorm.io/gorm"

type Product struct {
	BaseModel
	Name        string      `json:"name" gorm:"size:32;not null"`
	Description string      `json:"description" gorm:"size:256;not null"`
	Images      StringArray `json:"images"  gorm:"not null"`
	// Price in cent, $2.70 = 270
	Price      int         `json:"price" gorm:"not null"`
	Type       ProductType `json:"type" gorm:"not null"`
	Closed     bool        `json:"closed" gorm:"not null"`
	UserID     int         `json:"user_id" gorm:"index;not null"`
	CategoryID int         `json:"category_id" gorm:"not null"`
	// 由于目前租户较少，暂不添加索引
	TenantID int `json:"tenant_id" gorm:"not null"`
}

type ProductType = int8

const (
	ProductTypeBid ProductType = -1 // buy
	ProductTypeAsk ProductType = 1  // sell
)

func (p *Product) BeforeSave(tx *gorm.DB) (err error) {
	if p.Images == nil {
		p.Images = []string{}
	}
	return nil
}
