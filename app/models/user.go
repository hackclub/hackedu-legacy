package models

import (
	"fmt"
	"github.com/robfig/revel"
)

type User struct {
	UserId                               int
	FirstName, LastName, Email, Password string
	HashedPassword                       []byte
}

func (u *User) String() string {
	return fmt.Sprintf("User(%s)", u.Email)
}

func (u *User) FullName() string {
	return fmt.Sprintf("%s %s", u.FirstName, u.LastName)
}

func (user *User) Validate(v *revel.Validation) {
	v.Required(user.FirstName).Message("Your first name is required.")
	v.MaxSize(user.FirstName, 256).Message("Your first name is too long.")

	v.Required(user.LastName).Message("Your last name is required.")
	v.MaxSize(user.LastName, 256).Message("Your last name is too long.")

	v.Email(user.Email).Message("You must provide a valid email address.")

	v.Required(user.Password).Message("A password is required.")
	v.MinSize(user.Password, 5).Message("Your password must be at least 5 characters.")
	v.MaxSize(user.Password, 256).Message("Your password is too long.")
}
