package repository

import (
	"context"

	"github.com/user_service/internal/auth/domain/entity"
	"github.com/user_service/internal/auth/domain/vo"
)

type UserRepository interface {
	FindByUsername(ctx context.Context, username string) (*entity.User, error)
	UpdateLastLogin(ctx context.Context, userID string) error
	GetUserByEmail(ctx context.Context, email vo.Email) (*entity.User, error)
	CreateNewUser(ctx context.Context, user *entity.User) error
	GetUserByUserID(ctx context.Context, userID string) (*entity.User, error)
}
