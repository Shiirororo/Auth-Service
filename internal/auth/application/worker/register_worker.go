package worker

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/user_service/internal/auth/domain/model/entity"
	auth_repository "github.com/user_service/internal/auth/domain/repository"
	"github.com/user_service/internal/auth/domain/vo"
	common_entity "github.com/user_service/internal/commons/domains/model/entity"
	common_repository "github.com/user_service/internal/commons/domains/repository"
	"github.com/user_service/internal/event"
	user_entity "github.com/user_service/internal/user/domain/model/entity"
	user_repository "github.com/user_service/internal/user/domain/repository"
	"golang.org/x/crypto/bcrypt"
)

type RegisterWorker struct {
	queue       chan string
	workers     int
	authRepo    auth_repository.AuthRepository
	userRepo    common_repository.UserRepository
	profileRepo user_repository.ProfileRepository
	roleRepo    common_repository.RoleRepository
}

func NewRegisterWorker(
	authRepo auth_repository.AuthRepository,
	userRepo common_repository.UserRepository,
	profileRepo user_repository.ProfileRepository,
	roleRepo common_repository.RoleRepository,
	dispatcher *event.Dispatcher,
) *RegisterWorker {
	worker := &RegisterWorker{
		authRepo:    authRepo,
		userRepo:    userRepo,
		profileRepo: profileRepo,
		roleRepo:    roleRepo,
	}
	dispatcher.Register(event.RegisterSuccessEvent, worker)
	return worker
}

func (w *RegisterWorker) Handle(ctx context.Context, e event.Event) error {
	payload, ok := e.Payload.(event.RegisterSuccessPayload)
	if !ok {
		log.Println("Invalid payload for RegisterWorker")
		return nil
	}

	// 1. Hash password
	hashPass, err := bcrypt.GenerateFromPassword([]byte(payload.Password), 10)
	if err != nil {
		log.Println("Failed to hash password:", err)
		return err
	}

	// 2. Generate UUID
	id, err := uuid.NewV7()
	if err != nil {
		return err
	}

	idBytes := id[:]

	// 3. User Entity
	user := &common_entity.User{
		ID:       idBytes,
		Username: payload.Username,
		State:    common_entity.UserActive,
	}
	if err := w.userRepo.CreateNewUser(ctx, user); err != nil {
		log.Println("Failed to create user:", err)
		return err
	}

	// 4. User Auth
	passVo := vo.RestorePassword(string(hashPass))
	auth := entity.NewAuth(idBytes, payload.Email, passVo)
	if err := w.authRepo.CreateAuth(ctx, auth); err != nil {
		log.Println("Failed to create auth:", err)
		return err
	}

	// 5. User Profile
	profile := &user_entity.UserProfile{
		UserID:      idBytes,
		ProfileName: payload.Username,
	}
	if err := w.profileRepo.CreateProfile(ctx, profile); err != nil {
		log.Println("Failed to create profile:", err)
		return err
	}

	// 6. User Role (default user role: ID=1)
	userRole := &common_entity.UserRole{
		UserID: id,
		RoleID: 1, // Default user role
	}
	if err := w.roleRepo.AssignRoleToUser(ctx, userRole); err != nil {
		log.Println("Failed to assign role:", err)
		return err
	}

	log.Println("Register worker completed successfully for user:", payload.Username)

	return nil
}
