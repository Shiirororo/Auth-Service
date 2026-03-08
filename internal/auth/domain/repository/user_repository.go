package repository

import (
	"context"

	"github.com/user_service/internal/auth/domain/model/entity"
	"github.com/user_service/internal/auth/domain/vo"
)

type AuthRepository interface {
	//FindByUsername(ctx context.Context, username string) (*entity.Auth, error)
	UpdateLastLogin(ctx context.Context, userID string) error
	GetUserByEmail(ctx context.Context, email vo.Email) (*entity.Auth, error)
	CreateNewUser(ctx context.Context, user *entity.Auth) error
	GetUserByUserID(ctx context.Context, userID string) (*entity.Auth, error)
}
