package entity

import (
	"time"
)

type UserProfile struct {
	UserID []byte `gorm:"column:user_id;type:binary(16);primaryKey"`

	ProfileName string     `gorm:"column:profile_name;size:50;not null;index:idx_profile_name"`
	Mobile      *string    `gorm:"column:mobile;size:20"`
	Gender      *uint8     `gorm:"column:gender"`
	Birthday    *time.Time `gorm:"column:birthday"`
}
type UserUpdateEntity struct {
	ProfileName string     `json:"profile_name"`
	Mobile      *string    `json:"mobile,omitempty"`
	Gender      *uint8     `json:"gender,omitempty"`
	Birthday    *time.Time `json:"birthday,omitempty"`
}

func (UserProfile) TableName() string {
	return "user_profile"
}
