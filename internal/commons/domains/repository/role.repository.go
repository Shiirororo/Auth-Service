package commons_repository

import (
	"context"

	"github.com/user_service/internal/commons/domains/model/entity"
)

type RoleRepository interface {
	AssignRoleToUser(ctx context.Context, userRole *entity.UserRole) error
}
