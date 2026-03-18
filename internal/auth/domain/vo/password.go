package vo

import (
	"errors"
)

var (
	ErrPasswordTooShort = errors.New("password must be at least 6 characters")
)

type Password struct {
	value string
}

func NewPassword(val string) (Password, error) {
	if len(val) < 6 {
		return Password{}, ErrPasswordTooShort
	}
	return Password{value: val}, nil
}

// RestorePassword is used when loading an already hashed password from the persistence layer
func RestorePassword(val string) Password {
	return Password{value: val}
}

func (p Password) String() string {
	return p.value
}
