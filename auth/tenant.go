package auth

import (
	"github.com/gofiber/fiber/v2"
	. "keyi/models"
)

type Tenant struct {
	ID          int    `json:"id" gorm:"primarykey"`
	Name        string `json:"name" gorm:"size:32;unique;not null"`
	Domains     string `json:"domains" gorm:"size:256"` // separate by comma
	TenantAreas []TenantArea
}

type TenantArea struct {
	ID       int    `json:"id" gorm:"primarykey" `
	TenantID int    `json:"tenant_id"`
	Name     string `json:"name" gorm:"size:32;unique;not null"`
}

// ListTenants
// @Summary List Tenants
// @Description List all tenants with areas each
// @Tags Tenant
// @Produce application/json
// @Router /tenants [get]
// @Success 200 {array} Tenant
func ListTenants(c *fiber.Ctx) error {
	var tenants []Tenant
	err := DB.
		Preload("TenantAreas").
		Find(&tenants, "tenant.id > ?", 0).Error
	if err != nil {
		return err
	}

	for i := range tenants {
		if len(tenants[i].TenantAreas) == 0 {
			tenants[i].TenantAreas = []TenantArea{
				{Name: "Default"},
			}
		}
	}

	return c.JSON(&tenants)
}
