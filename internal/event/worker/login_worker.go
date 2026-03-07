package worker

import (
	"context"

	"github.com/user_service/internal/auth/domain/repository"
	"github.com/user_service/internal/event"
)

type LoginWorker struct {
	authRepo repository.UserRepository
	queue    chan string
	workers  int
}

// NewLoginWorker creates a worker pool that handles login events.
// It subscribes to the dispatcher so it can receive events.
func NewLoginWorker(authRepo repository.UserRepository, dispatcher *event.Dispatcher, workers int) *LoginWorker {
	worker := &LoginWorker{
		authRepo: authRepo,
		queue:    make(chan string, 1000), // Buffer to handle login spikes Without Blocking
		workers:  workers,
	}

	// Register this worker as the handler for LoginEvents
	dispatcher.Register(event.LoginEvent, worker)
	return worker
}

// Handle implements event.EventHandler. It receives events from the dispatcher and queues them.
func (w *LoginWorker) Handle(ctx context.Context, e event.Event) error {
	if userID, ok := e.Payload.(string); ok {
		w.queue <- userID
	}
	return nil
}

// Start spawns the goroutines that will pull from the queue and execute DB queries sequentially per worker.
func (w *LoginWorker) Start(ctx context.Context) {
	for i := 0; i < w.workers; i++ {
		go w.processQueue(ctx)
	}
}

func (w *LoginWorker) processQueue(ctx context.Context) {
	for {
		select {
		case userID := <-w.queue:
			w.authRepo.UpdateLastLogin(ctx, userID)
		case <-ctx.Done():
			return
		}
	}
}
