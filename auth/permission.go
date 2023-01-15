package auth

type Permission = uint8

const (
	PUser Permission = 1 << iota
	PAdmin
	POperator
)

func PermissionToString(p Permission) string {
	switch p {
	case PUser:
		return "user"
	case PAdmin:
		return "admin"
	case POperator:
		return "operator"
	default:
		return "unknown"
	}
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
