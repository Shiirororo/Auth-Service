package repository

import (
	"context"

	"github.com/user_service/internal/user/domain/model/entity"
)

type ProfileRepository interface {
	GetUserInfor(ctx context.Context, userID []byte) (*entity.UserProfile, error)
	CreateProfile(ctx context.Context, profile *entity.UserProfile) error
	UpdateUser(ctx context.Context, userID []byte, data entity.UserUpdateEntity) error
}
