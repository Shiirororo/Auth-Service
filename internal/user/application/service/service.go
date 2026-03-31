package user_service

import (
	"context"
	"encoding/hex"
	"errors"
	"strings"

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
	dispatcher  *event.Dispatcher
}

func NewUserService(profileRepo repository.ProfileRepository, dispatcher *event.Dispatcher) UserServiceInterface {
	return &UserService{
		profileRepo: profileRepo,
		dispatcher:  dispatcher,
	}
}

func (s *UserService) RegisterService(ctx context.Context, username string, password string, email string) error {
	emailCh := make(chan bool, 1)
	usernameCh := make(chan bool, 1)

	s.dispatcher.Dispatch(ctx, event.Event{
		Type:    event.CheckEmailEvent,
		Payload: event.CheckEmailPayload{Email: email, ReplyCh: emailCh},
	})
	s.dispatcher.Dispatch(ctx, event.Event{
		Type:    event.CheckUsernameEvent,
		Payload: event.CheckUsernamePayload{Username: username, ReplyCh: usernameCh},
	})

	for i := 0; i < 2; i++ {
		select {
		case exists := <-emailCh:
			if exists {
				return errors.New("email already exists")
			}
		case exists := <-usernameCh:
			if exists {
				return errors.New("username already exists")
			}
		case <-ctx.Done():
			return ctx.Err()
		}
	}

	passVo, err := vo.NewPassword(password)
	if err != nil {
		return err
	}

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
	if e := s.profileRepo.UpdateUser(ctx, idBytes, updateData); e != nil {
		return errors.New("fatal, cannot update")
	}
	return nil
}
