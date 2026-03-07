package entity

import (
	"time"

	"github.com/user_service/internal/auth/domain/vo"
)

type User struct {
	ID           string
	Username     string
	Email        vo.Email
	PasswordHash vo.Password
	LastLogin    *time.Time
	LockedUntil  *time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// NewUser creates a new Domain User enforcing all invariants correctly.
func NewUser(id string, username string, email vo.Email, passwordHash vo.Password) *User {
	now := time.Now()
	return &User{
		ID:           id,
		Username:     username,
		Email:        email,
		PasswordHash: passwordHash,
		CreatedAt:    now,
		UpdatedAt:    now,
	}
}

func (u *User) UpdateLastLogin() {
	now := time.Now()
	u.LastLogin = &now
	u.UpdatedAt = now
}
