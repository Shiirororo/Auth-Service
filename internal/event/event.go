package event

import (
	"time"
)

type Event interface {
	Name() string
}

type LoginEvent struct {
	UserID    string
	SessionID string
	IPAddress string
	Timestamp time.Time
}

func (e *LoginEvent) Name() string {
	return "user.login"
}
