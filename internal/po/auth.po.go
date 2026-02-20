package po

import (
	"time"

	"github.com/google/uuid"
)

type AuthUser struct {
	ID           uuid.UUID  `gorm:"type:char(36);primaryKey"`
	Username     string     `gorm:"type:varchar(50);unique;not null"`
	Email        *string    `gorm:"type:varchar(100);unique"`
	PasswordHash string     `gorm:"type:text;not null"`
	LastLogin    *time.Time `gorm:"type:timestamp"`
	LockedUntil  *time.Time `gorm:"type:timestamp"`
	CreatedAt    time.Time  `gorm:"autoCreateTime"`
	UpdatedAt    time.Time  `gorm:"autoUpdateTime"`
}

func (u *AuthUser) TableName() string {
	return "auth_users"
}
