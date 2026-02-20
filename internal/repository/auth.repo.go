package repository

import (
	"context"
	"time"

	"github.com/auth_service/internal/po"
	"gorm.io/gorm"
)

type authRepository struct {
	db *gorm.DB
}

type AuthRepository interface {
	FindByUsername(ctx context.Context, username string) (*po.AuthUser, error)
	UpdateLastLogin(ctx context.Context, userID string) error
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db: db}
}

func (r *authRepository) FindByUsername(ctx context.Context, username string) (*po.AuthUser, error) {
	var user po.AuthUser

	err := r.db.
		Where("username = ?", username).
		First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *authRepository) GetUser(ctx context.Context, userID string) (*po.AuthUser, error) {
	var user po.AuthUser

	err := r.db.
		Where("id = ?", userID).
		First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *authRepository) GetUserByEmail(ctx context.Context, email string) (*po.AuthUser, error) {
	var user po.AuthUser

	err := r.db.
		Where("email = ?", email).
		First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}
func (r *authRepository) UpdateLastLogin(ctx context.Context, userID string) error {
	var user po.AuthUser
	now := time.Now()
	user.LastLogin = &now
	return r.db.Save(&user).Error
}

// func(r *authRepository) UpdateUser (ctx context.Context, u *po.AuthUser) (*po.AuthUser, error) {
// 	_, err := r
// }
