package models

import (
	"github.com/revel/revel"
)

type User struct {
	UserId             int
	Name               string
	Username, Password string
	HashedPassword     []byte
}

func (u *User) Validate(v *revel.Validation) {
	ValidateUsername(v, u.Username)
	ValidatePassword(v, u.Password).Key("u.Password")
	ValidateName(v, u.Name)
}

func ValidateUsername(v *revel.Validation, username string) {
	v.Check(username,
		revel.Required{},
		revel.MaxSize{15},
		revel.MinSize{4},
	)
}

func ValidateName(v *revel.Validation, name string) {
	v.Check(name,
		revel.Required{},
		revel.MaxSize{100},
	)
}

func ValidatePassword(v *revel.Validation, password string) *revel.ValidationResult {
	return v.Check(password,
		revel.Required{},
		revel.MaxSize{15},
		revel.MinSize{5},
	)
}