package persistence

import (
	"time"

	"github.com/user_service/internal/auth/domain/entity"
	"github.com/user_service/internal/auth/domain/vo"
)

// UserModel is the infrastructure representation of a user in the MySQL database.
type UserModel struct {
	ID           string     `gorm:"type:char(36);primaryKey;default:(UUID())"`
	Username     string     `gorm:"type:varchar(50);unique;not null"`
	Email        string     `gorm:"type:varchar(100);uniqueIndex"`
	PasswordHash string     `gorm:"type:varchar(72);not null"`
	LastLogin    *time.Time `gorm:"type:timestamp"`
	LockedUntil  *time.Time `gorm:"type:timestamp"`
	CreatedAt    time.Time  `gorm:"autoCreateTime"`
	UpdatedAt    time.Time  `gorm:"autoUpdateTime"`
}

func (u *UserModel) TableName() string {
	return "auth_users"
}

// ToDomain maps the GORM infrastructure model to a pure Domain Entity.
func (u *UserModel) ToDomain() (*entity.User, error) {
	emailVo, err := vo.NewEmail(u.Email)
	if err != nil {
		return nil, err
	}

	passVo := vo.RestorePassword(u.PasswordHash)

	return &entity.User{
		ID:           u.ID,
		Username:     u.Username,
		Email:        emailVo,
		PasswordHash: passVo,
		LastLogin:    u.LastLogin,
		LockedUntil:  u.LockedUntil,
		CreatedAt:    u.CreatedAt,
		UpdatedAt:    u.UpdatedAt,
	}, nil
}

// FromDomain maps a pure Domain Entity into a GORM infrastructure model for persistence.
func FromDomain(user *entity.User) *UserModel {
	return &UserModel{
		ID:           user.ID,
		Username:     user.Username,
		Email:        user.Email.String(),
		PasswordHash: user.PasswordHash.String(),
		LastLogin:    user.LastLogin,
		LockedUntil:  user.LockedUntil,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
	}
}

type ProfileUserModel struct {
	ID          string    `gorm:"type:char(36);primaryKey;default:(UUID())"`
	Username    string    `gorm:"type:varchar(50);unique;not null"`
	ProfileName string    `gorm:"type:varchar(50);not null"`
	UserState   int       `gorm:"type:tinyint; not null"`
	UserMobile  int       `gorm:"type:varchar(50);not null"`
	UserEmail   string    `gorm:"type:varchar(100);uniqueIndex"`
	UserGender  int       `gorm:"type:tinyint; not null"`
	UserBirth   time.Time `gorm:"type:date; not null"`
}

func (u *ProfileUserModel) TableName() string {
	return "profile_users"
}
