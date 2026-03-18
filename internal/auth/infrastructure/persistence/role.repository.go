package persistence

import (
	"context"

	"github.com/user_service/internal/auth/domain/model/entity"
	"github.com/user_service/internal/auth/domain/repository"
	"gorm.io/gorm"
)

type roleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) repository.RoleRepository {
	return &roleRepository{db: db}
}

func (rr *roleRepository) AssignRoleToUser(ctx context.Context, userRole *entity.UserRole) error {
	return rr.db.WithContext(ctx).Create(userRole).Error
}
