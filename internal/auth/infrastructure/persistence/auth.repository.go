package persistence

import (
	"context"

	"github.com/user_service/internal/auth/domain/model/entity"
	"github.com/user_service/internal/auth/domain/repository"
	"github.com/user_service/internal/auth/domain/vo"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.AuthRepository {
	return &userRepository{db: db}
}

// func (r *userRepository) FindByUsername(ctx context.Context, username string) (*entity.Auth, error) {
// 	var model entity.Auth

// 	err := r.db.
// 		WithContext(ctx).
// 		Where("username = ?", username).
// 		First(&model).Error

// 	if err != nil {
// 		return nil, err
// 	}

// 	return model.ToDomain()
// }

func (r *userRepository) GetUserByUserID(ctx context.Context, userID string) (*entity.Auth, error) {
	var model entity.Auth
	err := r.db.
		WithContext(ctx).
		Where("id = ?", userID).
		First(&model).Error

	if err != nil {
		return nil, err
	}

	return model.ToDomain()
}

func (r *userRepository) CreateNewUser(ctx context.Context, user *entity.Auth) error {
	model := entity.FromDomain(user)
	return r.db.WithContext(ctx).Create(model).Error
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email vo.Email) (*entity.Auth, error) {
	var model entity.Auth

	err := r.db.
		WithContext(ctx).
		Where("email = ?", email.String()).
		First(&model).Error

	if err != nil {
		return nil, err
	}

	return model.ToDomain()
}

func (r *userRepository) UpdateLastLogin(ctx context.Context, userID string) error {
	// Only update what is strictly required from DB
	return r.db.
		WithContext(ctx).
		Model(&entity.Auth{}).
		Where("id = ?", userID).
		Update("last_login", gorm.Expr("NOW()")).
		Error
}
