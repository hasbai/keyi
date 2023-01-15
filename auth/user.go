package auth

import (
	"fmt"
	"keyi/models"
)

type User struct {
	models.BaseModel
	Username     string     `json:"username" gorm:"size:32;not null"`
	Description  string     `json:"description" gorm:"size:256;not null"`
	Email        string     `json:"email" gorm:"size:64;unique;not null"`
	OpenID       string     `json:"openid" gorm:"size:64;unique;not null"`
	Permission   Permission `json:"permission" gorm:"not null"`
	Avatar       string     `json:"avatar" gorm:"size:256;not null"`
	TenantID     int        `json:"tenant_id" gorm:"not null"`
	Tenant       Tenant     `json:"tenant"`
	TenantAreaID int        `json:"tenant_area_id" gorm:"not null"` // 0 is default area
	TenantArea   TenantArea `json:"tenant_area"`
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
