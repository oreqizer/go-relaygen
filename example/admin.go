//go:generate go run github.com/oreqizer/go-relaygen Admin User.LocalID

package example

type Admin struct {
	User
	Nickname string
}

func (u *Admin) ID() string {
	return "Admin:" + u.User.LocalID
}
