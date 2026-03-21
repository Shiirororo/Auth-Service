package user_service

import (
	"context"
	"encoding/hex"
	"errors"
	"strings"

	auth_repository "github.com/user_service/internal/auth/domain/repository"
	"github.com/user_service/internal/auth/domain/vo"
	"github.com/user_service/internal/event"
	"github.com/user_service/internal/user/controller/dto"
	"github.com/user_service/internal/user/domain/model/entity"
	"github.com/user_service/internal/user/domain/repository"
)

type UserServiceInterface interface {
	RegisterService(ctx context.Context, username string, password string, email string) error
	GetUserInfo(ctx context.Context, userID string) (*dto.UserProfileResponse, error)
	UpdateUserInfo(ctx context.Context, userID string, data dto.UserUpdateRequest) error
}

type UserService struct {
	profileRepo repository.ProfileRepository
	authRepo    auth_repository.AuthRepository
	dispatcher  *event.Dispatcher
}

func NewUserService(profileRepo repository.ProfileRepository, authRepo auth_repository.AuthRepository, dispatcher *event.Dispatcher) UserServiceInterface {
	return &UserService{
		profileRepo: profileRepo,
		authRepo:    authRepo,
		dispatcher:  dispatcher,
	}
}

func (s *UserService) RegisterService(ctx context.Context, username string, password string, email string) error {
	// 1. Create Domain Value Objects to validate integrity early
	// Email domain validation is handled by the presentation layer (gin binding)

	//Check duplicate email and username
	_, err := s.authRepo.GetUserByEmail(ctx, email)
	if err == nil {
		return errors.New("Email already exists")
	}

	passVo, err := vo.NewPassword(password)
	if err != nil {
		return err
	}

	// ASSUME OTP AND VERIFICATION COMPLETED HERE

	s.dispatcher.Dispatch(ctx, event.Event{
		Type: event.RegisterSuccessEvent,
		Payload: event.RegisterSuccessPayload{
			Username: username,
			Email:    email,
			Password: passVo.String(),
		},
	})
	return nil
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

func (s *UserService) UpdateUserInfo(ctx context.Context, userID string, data dto.UserUpdateRequest) error {
	cleanID := strings.TrimPrefix(userID, "0x")
	cleanID = strings.TrimPrefix(cleanID, "0X")

	idBytes, err := hex.DecodeString(cleanID)
	if err != nil {
		return errors.New("invalid user ID format")
	}
	if len(idBytes) != 16 {
		return errors.New("invalid user ID length")
	}
	updateData := entity.UserUpdateEntity{
		ProfileName: data.Data.ProfileName,
		Mobile:      data.Data.Mobile,
		Gender:      data.Data.Gender,
		Birthday:    data.Data.Birthday,
	}
	e := s.profileRepo.UpdateUser(ctx, idBytes, updateData)
	if e != nil {
		return errors.New("Fatal, cannot update")
	}
	return nil
}
