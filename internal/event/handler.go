package event

import (
	"context"

	"github.com/user_service/internal/repository"
)

type EventHandler interface {
	Handle(ctx context.Context, event Event) error
}

type LoginEventHandler struct {
	authRepo repository.AuthRepository
}

func NewLoginEventHandler(authRepo repository.AuthRepository) EventHandler {
	return &LoginEventHandler{authRepo: authRepo}
}

func (h *LoginEventHandler) Handle(ctx context.Context, event Event) error {
	if loginEvent, ok := event.(*LoginEvent); ok {
		return h.authRepo.UpdateLastLogin(ctx, loginEvent.UserID)
	}
	return nil
}
