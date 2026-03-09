package commons_repository

import (
	"context"

	"github.com/user_service/internal/commons/domains/model/entity"
)

type UserRepository interface {
	CreateNewUser(ctx context.Context, user *entity.User) error
}
