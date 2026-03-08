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

func (a *Auth) AuthTableName() string {
	return "user_auth"
}

func (a *Auth) ToDomain() (A *Auth, err error) {
	passVo := vo.RestorePassword(a.PasswordHash)

	return &Auth{
		UserID:       a.UserID,
		PasswordHash: passVo.String(),
		LastLogin:    a.LastLogin,
		LockedUntil:  a.LockedUntil,
	}, nil
}
func FromDomain(user *Auth) *Auth {
	return &Auth{
		UserID:       user.UserID,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
		LastLogin:    user.LastLogin,
		LockedUntil:  user.LockedUntil,
	}
}

// NewUser creates a new Domain User enforcing all invariants correctly.
func NewAuth(id string, email vo.Email, passwordHash vo.Password) *Auth {
	return &Auth{
		UserID:       []byte(id),
		Email:        email.String(),
		PasswordHash: passwordHash.String(),
	}
}

func (u *Auth) UpdateLastLogin() {
	now := time.Now()
	u.LastLogin = &now
}
