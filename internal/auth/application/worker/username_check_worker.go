package worker

import (
	"context"
	"log"

	"github.com/user_service/internal/auth/domain/repository"
	"github.com/user_service/internal/event"
)

// UsernameCheckWorker handles CheckUsernameEvent on behalf of the Auth domain.
type UsernameCheckWorker struct {
	userRepo repository.UserRepository
}

func NewUsernameCheckWorker(userRepo repository.UserRepository, dispatcher *event.Dispatcher) *UsernameCheckWorker {
	w := &UsernameCheckWorker{userRepo: userRepo}
	dispatcher.Register(event.CheckUsernameEvent, w)
	return w
}

func (w *UsernameCheckWorker) Handle(ctx context.Context, e event.Event) error {
	payload, ok := e.Payload.(event.CheckUsernamePayload)
	if !ok {
		log.Println("UsernameCheckWorker: invalid payload")
		return nil
	}

	_, err := w.userRepo.GetUserByUsername(ctx, payload.Username)
	payload.ReplyCh <- (err == nil) // true = username exists
	return nil
}
