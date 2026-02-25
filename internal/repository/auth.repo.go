package repository

import (
	"context"
	"time"

	"github.com/auth_service/internal/po"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type authRepository struct {
	db *gorm.DB
}

type AuthRepository interface {
	FindByUsername(ctx context.Context, username string) (*po.AuthUser, error)
	UpdateLastLogin(ctx context.Context, userID string) error
	CreateNewUser(ctx context.Context, usename string, password string, email string) error
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db: db}
}

func (r *authRepository) FindByUsername(ctx context.Context, username string) (*po.AuthUser, error) {
	var user po.AuthUser

	err := r.db.
		Where("username = ?", username).
		First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *authRepository) GetUser(ctx context.Context, userID uuid.UUID) (*po.AuthUser, error) {
	var user po.AuthUser
	err := r.db.
		Where("id = ?", userID).
		First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}
func (r *authRepository) CreateNewUser(ctx context.Context, username string, password string, email string) error {

	hashPass, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return err
	}

	NewUser := po.AuthUser{
		ID:           uuid.NewString(),
		Username:     username,
		PasswordHash: string(hashPass),
		Email:        email,
	}

	err = r.db.Create(&NewUser).Error
	return err
}

func (r *authRepository) GetUserByEmail(ctx context.Context, email string) (*po.AuthUser, error) {
	var user po.AuthUser

	err := r.db.
		Where("email = ?", email).
		First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}
func (r *authRepository) UpdateLastLogin(ctx context.Context, userID string) error {
	now := time.Now()

	return r.db.
		WithContext(ctx).
		Model(&po.AuthUser{}).
		Where("id = ?", userID).
		Update("last_login", now).
		Error
}
