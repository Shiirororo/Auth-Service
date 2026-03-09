package repository

import (
	"context"

	"github.com/user_service/internal/auth/domain/model/entity"
)

type AuthRepository interface {
	UpdateLastLogin(ctx context.Context, userID []byte) error
	GetUserByEmail(ctx context.Context, email string) (*entity.Auth, error)
	CreateAuth(ctx context.Context, user *entity.Auth) error
	GetUserByUserID(ctx context.Context, userID []byte) (*entity.Auth, error)
}
