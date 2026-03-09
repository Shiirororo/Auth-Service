package commons_persistence

import (
	"context"

	"github.com/user_service/internal/commons/domains/model/entity"
	commons_repository "github.com/user_service/internal/commons/domains/repository"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) commons_repository.UserRepository {
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
