package entity

import (
	"time"

	"github.com/user_service/internal/auth/domain/vo"
)

type Auth struct {
	UserID       []byte     `gorm:"column:user_id;type:binary(16);primaryKey"`
	Email        string     `gorm:"column:email;size:100;uniqueIndex"`
	PasswordHash string     `gorm:"column:password_hash;size:72;not null"`
	LastLogin    *time.Time `gorm:"column:last_login"`
	LockedUntil  *time.Time `gorm:"column:locked_until"`
}

func (a *Auth) TableName() string {
	return "user_auth"
}

// NewUser creates a new Domain User enforcing all invariants correctly.
func NewAuth(id []byte, email string, passwordHash vo.Password) *Auth {
	return &Auth{
		UserID:       id,
		Email:        email,
		PasswordHash: passwordHash.String(),
	}
}

func (u *Auth) UpdateLastLogin() {
	now := time.Now()
	u.LastLogin = &now
}
