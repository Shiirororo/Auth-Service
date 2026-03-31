package worker

import (
	"context"
	"log"

	"github.com/user_service/internal/auth/domain/repository"
	"github.com/user_service/internal/event"
)

// EmailCheckWorker handles CheckEmailEvent on behalf of the Auth domain.
// It queries the auth repository and replies on the payload's ReplyCh.
type EmailCheckWorker struct {
	authRepo repository.AuthRepository
}

func NewEmailCheckWorker(authRepo repository.AuthRepository, dispatcher *event.Dispatcher) *EmailCheckWorker {
	w := &EmailCheckWorker{authRepo: authRepo}
	dispatcher.Register(event.CheckEmailEvent, w)
	return w
}

func (w *EmailCheckWorker) Handle(ctx context.Context, e event.Event) error {
	payload, ok := e.Payload.(event.CheckEmailPayload)
	if !ok {
		log.Println("EmailCheckWorker: invalid payload")
		return nil
	}

	_, err := w.authRepo.GetUserByEmail(ctx, payload.Email)
	payload.ReplyCh <- (err == nil) // true = email exists
	return nil
}
