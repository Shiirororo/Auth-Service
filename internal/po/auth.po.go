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
type ProfileUser struct {
	ID          string    `gorm:"type:char(36);primaryKey;default:(UUID())"`
	Username    string    `gorm:"type:varchar(50);unique;not null"`
	ProfileName string    `gorm:"type:varchar(50);not null"`
	UserState   int       `gorm:"type:tinyint; not null"`
	UserMobile  int       `gorm:"type:varchar(50);not null"`
	UserEmail   string    `gorm:"type:varchar(100);uniqueIndex"`
	UserGender  int       `gorm:"type:tinyint; not null"`
	UserBirth   time.Time `gorm:"type:date; not null"`
}

func (u *ProfileUser) TableName() string {
	return "profile_users"
}
func (u *AuthUser) TableName() string {
	return "auth_users"
}
