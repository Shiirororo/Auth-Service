package repository

import (
	"context"

	"github.com/user_service/internal/auth/domain/model/entity"
)

type UserRepository interface {
	CreateNewUser(ctx context.Context, user *entity.User) error
}
