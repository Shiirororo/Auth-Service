package persistence

import (
	"context"

	"github.com/user_service/internal/auth/domain/model/entity"
	"github.com/user_service/internal/auth/domain/repository"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateNewUser(ctx context.Context, user *entity.User) error {
	var model entity.User = *user
	err := r.db.WithContext(ctx).Create(&model).Error
	if err != nil {
		return err
	}
	return nil
}
