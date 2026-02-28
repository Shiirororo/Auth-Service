package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/user_service/internal/po"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type authRepository struct {
	db *gorm.DB
}

type AuthRepository interface {
	FindByUsername(ctx context.Context, username string) (*po.AuthUser, error)
	UpdateLastLogin(ctx context.Context, userID string) error
	GetUserByEmail(ctx context.Context, email string) (*po.AuthUser, error)
	CreateNewUser(ctx context.Context, usename string, password string, email string) error
	GetUserByUserID(ctx context.Context, userID string) (*po.AuthUser, error)
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db: db}
}

func (r *authRepository) FindByUsername(ctx context.Context, username string) (*po.AuthUser, error) {
	var user po.AuthUser

	err := r.db.
		WithContext(ctx).
		Where("username = ?", username).
		First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *authRepository) GetUserByUserID(ctx context.Context, userID string) (*po.AuthUser, error) {
	var user po.AuthUser
	err := r.db.
		WithContext(ctx).
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

	err = r.db.WithContext(ctx).Create(&NewUser).Error
	return err
}

func (r *authRepository) GetUserByEmail(ctx context.Context, email string) (*po.AuthUser, error) {
	var user po.AuthUser

	err := r.db.
		WithContext(ctx).
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
