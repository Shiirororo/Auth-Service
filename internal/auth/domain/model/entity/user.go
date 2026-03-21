package entity

import "time"

type User struct {
	ID        []byte    `gorm:"column:id;type:binary(16);primaryKey"`
	Username  string    `gorm:"column:username;size:50;uniqueIndex;not null"`
	State     uint8     `gorm:"column:state;not null;default:0"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (u *User) TableName() string {
	return "users"
}

const (
	UserActive  = 0
	UserBanned  = 1
	UserSuspend = 2
)
