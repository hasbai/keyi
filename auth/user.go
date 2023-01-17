package auth

import (
	"fmt"
	"keyi/models"
)

type User struct {
	models.BaseModel
	Username     string      `json:"username" gorm:"size:32;not null"`
	Description  string      `json:"description" gorm:"size:256;not null"`
	Email        string      `json:"-" gorm:"size:64;unique;not null"`
	OpenID       string      `json:"-" gorm:"size:64;unique;not null"`
	Permission   Permission  `json:"permission" gorm:"not null"`
	Avatar       string      `json:"avatar" gorm:"size:256;not null"`
	Contacts     string      `json:"contacts" gorm:"size:32;default:'';not null"`
	TenantID     int         `json:"tenant_id" gorm:"not null"`
	Tenant       *Tenant     `json:"tenant,omitempty"`
	TenantAreaID int         `json:"tenant_area_id" gorm:"not null"` // 0 is default area
	TenantArea   *TenantArea `json:"tenant_area,omitempty"`
	Follow       []User      `json:"follow,omitempty" gorm:"many2many:user_follow"`
}

func (u *User) GetTokenInfo() *TokenInfo {
	return &TokenInfo{
		UserID:       u.ID,
		Permission:   u.Permission,
		TenantID:     u.TenantID,
		TenantAreaID: u.TenantAreaID,
	}
}

func (u *User) GetPermission() Permission {
	return u.Permission
}

func init() {
	fmt.Println("init models for auth...")
	err := models.DB.AutoMigrate(&User{}, &Tenant{}, &TenantArea{})
	if err != nil {
		panic(err)
	}
	models.DB.FirstOrCreate(&Tenant{Name: "Default"})
	models.DB.Exec(`UPDATE tenant SET id = 0 WHERE name = 'Default'`)
	models.DB.FirstOrCreate(&TenantArea{Name: "Default"})
	models.DB.Exec(`UPDATE tenant_area SET id = 0 WHERE name = 'Default'`)
}
