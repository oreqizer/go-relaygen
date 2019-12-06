//go:generate go run github.com/oreqizer/go-relaygen User LocalID

package example

import (
	"github.com/oreqizer/go-relaygen/relay"
)

const UserType = "User"

type User struct {
	LocalID string
	Name    string
}

func (u *User) ID() string {
	return relay.ToGlobalID(UserType, u.LocalID)
}
