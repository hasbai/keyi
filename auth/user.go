package auth

import (
	"fmt"
	"keyi/models"
)

type User struct {
	models.BaseModel
	Username     string     `json:"username" gorm:"size:32;unique;not null"`
	Description  string     `json:"description" gorm:"size:256"`
	Email        string     `json:"email" gorm:"size:64;unique;not null"`
	Password     string     `json:"password" gorm:"size:64;not null"`
	IsValid      bool       `json:"is_valid" gorm:"default:false"`
	TEL          string     `json:"tel" gorm:"size:16"`
	Avatar       string     `json:"avatar" gorm:"size:256"`
	TenantID     int        `json:"tenant_id"`
	Tenant       Tenant     `json:"tenant"`
	TenantAreaID int        `json:"tenant_area_id"` // 0 is default area
	TenantArea   TenantArea `json:"tenant_area"`
}

func (u *User) GetTokenInfo() *TokenInfo {
	return &TokenInfo{
		UserID:       u.ID,
		IsValid:      u.IsValid,
		TenantID:     u.TenantID,
		TenantAreaID: u.TenantAreaID,
	}
}

type Tenant struct {
	ID          int    `gorm:"primarykey" json:"id"`
	Name        string `json:"name" gorm:"size:32;unique;not null"`
	Domains     string `json:"domains" gorm:"size:256"` // separate by comma
	TenantAreas []TenantArea
}

type TenantArea struct {
	ID       int    `gorm:"primarykey" json:"id"`
	TenantID int    `json:"tenant_id"`
	Name     string `json:"name" gorm:"size:32;unique;not null"`
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
