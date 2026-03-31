package router

import (
	auth_worker "github.com/user_service/internal/auth/application/worker"
	"github.com/user_service/internal/event"
	user_worker "github.com/user_service/internal/user/application/worker"
)

// App aggregates the HTTP Router with background services (Dispatcher, Workers).
type App struct {
	Router              *Router
	Dispatcher          *event.Dispatcher
	LoginWorker         *auth_worker.LoginWorker
	RegisterWorker      *user_worker.RegisterWorker
	EmailCheckWorker    *auth_worker.EmailCheckWorker
	UsernameCheckWorker *auth_worker.UsernameCheckWorker
}

func NewApp(
	r *Router,
	dispatcher *event.Dispatcher,
	loginWorker *auth_worker.LoginWorker,
	registerWorker *user_worker.RegisterWorker,
	emailCheckWorker *auth_worker.EmailCheckWorker,
	usernameCheckWorker *auth_worker.UsernameCheckWorker,
) *App {
	return &App{
		Router:              r,
		Dispatcher:          dispatcher,
		LoginWorker:         loginWorker,
		RegisterWorker:      registerWorker,
		EmailCheckWorker:    emailCheckWorker,
		UsernameCheckWorker: usernameCheckWorker,
	}
}
