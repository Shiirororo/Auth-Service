package service

import (
	"context"
	"encoding/hex"
	"errors"
	"strings"

	"github.com/user_service/internal/user/controller/dto"
	"github.com/user_service/internal/user/domain/repository"
)

type UserServiceInterface interface {
	GetUserInfo(ctx context.Context, userID string) (*dto.UserProfileResponse, error)
}

type UserService struct {
	profileRepo repository.ProfileRepository
}

func NewUserService(profileRepo repository.ProfileRepository) UserServiceInterface {
	return &UserService{profileRepo: profileRepo}
}

func (s *UserService) GetUserInfo(ctx context.Context, userID string) (*dto.UserProfileResponse, error) {
	// Clean up hex string prefix if present (e.g., "0x019CD...")
	cleanID := strings.TrimPrefix(userID, "0x")
	cleanID = strings.TrimPrefix(cleanID, "0X")

	idBytes, err := hex.DecodeString(cleanID)
	if err != nil {
		return nil, errors.New("invalid user ID format")
	}

	if len(idBytes) != 16 {
		return nil, errors.New("invalid user ID length")
	}

	profile, err := s.profileRepo.GetUserInfor(ctx, idBytes)
	if err != nil {
		return nil, err
	}

	return &dto.UserProfileResponse{
		UserID:      userID,
		ProfileName: profile.ProfileName,
		Mobile:      profile.Mobile,
		Gender:      profile.Gender,
		Birthday:    profile.Birthday,
	}, nil
}
