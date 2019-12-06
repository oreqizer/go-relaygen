//go:generate go run github.com/oreqizer/go-relaygen Admin User.LocalID

package example

import (
	"github.com/oreqizer/go-relaygen/relay"
)

const AdminType = "Admin"

type Admin struct {
	User
	Nickname string
}

func (u *Admin) ID() string {
	return relay.ToGlobalID(AdminType, u.User.LocalID)
}
