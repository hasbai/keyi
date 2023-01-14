package models

type ProductRecord struct {
	BaseModel
	UserID    int `json:"user_id"    gorm:"index:uni,unique"`
	ProductID int `json:"product_id" gorm:"index:uni,unique"`
}
