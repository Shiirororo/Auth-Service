package entity

import (
	"github.com/google/uuid"
)

type UserRole struct {
	UserID uuid.UUID `gorm:"column:user_id;type:binary(16);primaryKey"`
	RoleID int       `gorm:"column:role_id;type:int;primaryKey"`
}

func (u *UserRole) TableName() string {
	return "user_roles"
}
