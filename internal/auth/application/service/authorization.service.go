package service

import (
	"context"
	"errors"

	"github.com/user_service/internal/auth/domain/repository"
)

type AuthorizationServiceInterface interface {
	CheckUserRole(ctx context.Context, userIDBytes []byte, requiredRoleID uint) (bool, error)
}

type AuthorizationService struct {
	roleRepo repository.RoleRepository
}

func NewAuthorizationService(roleRepo repository.RoleRepository) AuthorizationServiceInterface {
	return &AuthorizationService{
		roleRepo: roleRepo,
	}
}

func (s *AuthorizationService) CheckUserRole(ctx context.Context, userIDBytes []byte, requiredRoleID uint) (bool, error) {
	// Implementation placeholder: normally this pulls the user roles array from the repository and verifies
	// We'll return hardcoded true/false for structure, or a dummy error if repo lacks the exact method
	return false, errors.New("not fully implemented: missing FindUserRoles in RoleRepository")
}
