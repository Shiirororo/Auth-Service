package po

import (
	"time"
)

type AuthUser struct {
	ID           string     `gorm:"type:char(36);primaryKey;default:(UUID())"`
	Username     string     `gorm:"type:varchar(50);unique;not null"`
	Email        string     `gorm:"type:varchar(100);uniqueIndex"`
	PasswordHash string     `gorm:"type:varchar(72);not null"`
	LastLogin    *time.Time `gorm:"type:timestamp"`
	LockedUntil  *time.Time `gorm:"type:timestamp"`
	CreatedAt    time.Time  `gorm:"autoCreateTime"`
	UpdatedAt    time.Time  `gorm:"autoUpdateTime"`
}

func (u *AuthUser) TableName() string {
	return "auth_users"
}
