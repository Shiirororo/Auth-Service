package vo

import (
	"errors"
	"regexp"
)

var (
	ErrInvalidEmail = errors.New("invalid email format")
	emailRegex      = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
)

type Email struct {
	value string
}

func NewEmail(val string) (Email, error) {
	if !emailRegex.MatchString(val) {
		return Email{}, ErrInvalidEmail
	}
	return Email{value: val}, nil
}

func (e Email) String() string {
	return e.value
}
