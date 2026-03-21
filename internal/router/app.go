package router

import (
	auth_worker "github.com/user_service/internal/auth/application/worker"
	"github.com/user_service/internal/event"
	user_worker "github.com/user_service/internal/user/application/worker"
)

// App aggregates the HTTP Router with background services (Dispatcher, Workers).
// This keeps the pure HTTP routing concern of Router separate from event-driven workers.
type App struct {
	Router         *Router
	Dispatcher     *event.Dispatcher
	LoginWorker    *auth_worker.LoginWorker
	RegisterWorker *user_worker.RegisterWorker
}

func NewApp(
	r *Router,
	dispatcher *event.Dispatcher,
	loginWorker *auth_worker.LoginWorker,
	registerWorker *user_worker.RegisterWorker,
) *App {
	return &App{
		Router:         r,
		Dispatcher:     dispatcher,
		LoginWorker:    loginWorker,
		RegisterWorker: registerWorker,
	}
}
