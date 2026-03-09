package commons_persistence

import (
	"context"

	"github.com/user_service/internal/commons/domains/model/entity"
	common_repository "github.com/user_service/internal/commons/domains/repository"
	"gorm.io/gorm"
)

type roleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) common_repository.RoleRepository {
	return &roleRepository{db: db}
}

func (rr *roleRepository) AssignRoleToUser(ctx context.Context, userRole *entity.UserRole) error {
	return rr.db.WithContext(ctx).Create(userRole).Error
}
