//go:generate go run github.com/oreqizer/go-relaygen User LocalID

package example

type User struct {
	LocalID string
	Name    string
}

func (u *User) ID() string {
	return "User:" + u.LocalID
}
