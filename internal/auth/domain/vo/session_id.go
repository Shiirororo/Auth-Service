package vo

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrInvalidSessionID = errors.New("invalid session id format")
)

type SessionID struct {
	value string
}

func NewSessionID(val string) (SessionID, error) {
	if err := uuid.Validate(val); err != nil {
		return SessionID{}, ErrInvalidSessionID
	}
	return SessionID{value: val}, nil
}

func GenerateSessionID() SessionID {
	return SessionID{value: uuid.NewString()}
}

func (s SessionID) String() string {
	return s.value
}
