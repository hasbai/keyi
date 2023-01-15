package auth

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"keyi/utils"
)

type Permission = uint8

const (
	PUser Permission = 1 << iota
	PAdmin
	POperator
)

func PermissionToString(p Permission) string {
	out := ""
	if p&PUser != 0 {
		out += "User "
	}
	if p&PAdmin != 0 {
		out += "Admin "
	}
	if p&POperator != 0 {
		out += "Operator "
	}
	if out == "" {
		out = "None"
	}
	return out
}

type PermissionOwner interface {
	GetPermission() Permission
}

// CheckPermission checks user permission
// Example:
//  CheckPermission(u, Admin)
//  CheckPermission(u, Admin | Operator)
func CheckPermission(u PermissionOwner, p Permission) bool {
	return u.GetPermission()&p != 0
}

func AddPermission(u PermissionOwner, p Permission) Permission {
	return u.GetPermission() | p
}

func RemovePermission(u PermissionOwner, p Permission) Permission {
	return u.GetPermission() &^ p
}

// PermOnly allows the user with the permission
func PermOnly(c *fiber.Ctx, p Permission) error {
	claims := GetClaims(c)
	if !CheckPermission(claims, p) {
		return utils.Unauthorized(fmt.Sprintf(
			"permission denied, required: %s, got: %s",
			PermissionToString(p),
			PermissionToString(claims.Permission),
		))
	}
	return nil
}

// OwnerOrPerm allows the user with the permission or the owner of the object
func OwnerOrPerm(c *fiber.Ctx, p Permission, ownerID int) error {
	claims := GetClaims(c)
	if !CheckPermission(claims, p) && claims.UID != ownerID {
		return utils.Unauthorized(fmt.Sprintf(
			"permission denied, required: %s, got: %s",
			PermissionToString(p),
			PermissionToString(claims.Permission),
		))
	}
	return nil
}
