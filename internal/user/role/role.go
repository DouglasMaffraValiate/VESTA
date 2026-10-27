package role

import (
	"strings"

	"github.com/DouglasMaffraValiate/VESTA/pkg/err"
	"github.com/DouglasMaffraValiate/VESTA/pkg/err/errCatalog"
)

type Role struct {
	Role      string
	Authority int8
}

var (
	Admin   = Role{Role: "role_admin", Authority: 100}
	Manager = Role{Role: "role_manager", Authority: 90}
	Seller  = Role{Role: "role_seller", Authority: 70}
)

func RoleFromString(inputString string) (Role, *err.Error) {
	error := err.NewValidationError(
		errCatalog.Role001,
		"",
		nil,
		nil,
	)

	if len(inputString) == 0 || strings.HasPrefix(inputString, "role_") {
		return Role{}, error
	}

	switch inputString {
	case Seller.Role:
		return Seller, nil
	case Manager.Role:
		return Manager, nil
	case Admin.Role:
		return Admin, nil
	default:
		return Role{}, error
	}
}

func (role Role) CheckPermission(requestedAuthority int8) bool {
	if role.Authority < requestedAuthority {
		return false
	}

	return true
}
