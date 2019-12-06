//go:generate go run github.com/oreqizer/go-relaygen User Person.ID

package example

type Person struct {
	ID   string
	Name string
}

type User struct {
	Person
	Nickname string
}

func (u *User) ID() string {
	return u.Person.ID
}
