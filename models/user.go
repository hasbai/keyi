package models

type User struct {
	BaseModel
	Name        string    `json:"name" gorm:"size:32;unique;not null"`
	Description string    `json:"description" gorm:"size:256"`
	Email       string    `json:"-" gorm:"size:64;unique;not null"`
	Password    string    `json:"-" gorm:"size:128;not null"`
	Avatar      string    `json:"avatar" gorm:"size:256;not null"`
	Products    []Product `json:"products"`
}

func (u *User) CheckPassword(password string) bool {
	return u.Password == password
}
