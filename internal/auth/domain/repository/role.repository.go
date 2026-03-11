package repository

import (
	"context"

	"github.com/user_service/internal/auth/domain/model/entity"
)

type RoleRepository interface {
	AssignRoleToUser(ctx context.Context, userRole *entity.UserRole) error
}
