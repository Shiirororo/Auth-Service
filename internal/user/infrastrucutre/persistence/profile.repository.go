package persistence

import (
	"context"

	"github.com/user_service/internal/user/domain/model/entity"
	"github.com/user_service/internal/user/domain/repository"
	"gorm.io/gorm"
)

type profileRepository struct {
	db *gorm.DB
}

func NewProfileRepository(db *gorm.DB) repository.ProfileRepository {
	return &profileRepository{db: db}
}

func (pr *profileRepository) GetUserInfor(ctx context.Context, userID []byte) (*entity.UserProfile, error) {
	var profile entity.UserProfile
	err := pr.db.WithContext(ctx).Where("user_id", userID).First(&profile).Error
	if err != nil {
		return nil, err
	}
	return &profile, err
}

func (pr *profileRepository) CreateProfile(ctx context.Context, profile *entity.UserProfile) error {
	return pr.db.WithContext(ctx).
		Create(profile).Error
}

func (pr *profileRepository) UpdateUser(ctx context.Context, userID []byte, data map[string]interface{}) error {

	return pr.db.WithContext(ctx).
		Model(&entity.UserProfile{}).
		Where("user_id = ?", userID).
		Updates(data).Error
}
