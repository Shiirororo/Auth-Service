package persistence

import (
	"context"

	"github.com/user_service/internal/auth/domain/model/entity"
	"github.com/user_service/internal/auth/domain/repository"
	"gorm.io/gorm"
)

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) repository.AuthRepository {
	return &authRepository{db: db}
}

func (r *authRepository) GetUserByUserID(ctx context.Context, userID []byte) (*entity.Auth, error) {
	var model entity.Auth
	err := r.db.
		WithContext(ctx).
		Where("user_id = ?", userID).
		First(&model).Error

	if err != nil {
		return nil, err
	}

	return &model, err
}

func (r *authRepository) CreateAuth(ctx context.Context, user *entity.Auth) error {
	return r.db.WithContext(ctx).Create(user).Error
}

func (r *authRepository) GetUserByEmail(ctx context.Context, email string) (*entity.Auth, error) {
	var model entity.Auth

	err := r.db.
		WithContext(ctx).
		Where("email = ?", email).
		First(&model).Error

	if err != nil {
		return nil, err
	}

	return &model, err
}

func (r *authRepository) UpdateLastLogin(ctx context.Context, userID []byte) error {
	// Only update what is strictly required from DB
	return r.db.
		WithContext(ctx).
		Model(&entity.Auth{}).
		Where("user_id = ?", userID).
		Update("last_login", gorm.Expr("NOW()")).
		Error
}
