package entity

import (
	"github.com/google/uuid"
)

type UserRole struct {
	User_role uuid.UUID `gorm:"type:binary(16);primaryKey"`
	Role_ID   int       `gorm:"type:int;primaryKey"`
}

func (u *UserRole) TableName() string {
	return "user_roles"
}
