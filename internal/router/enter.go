package router

import (
	"github.com/user_service/internal/event"
	"github.com/user_service/internal/event/worker"
	"github.com/user_service/internal/router/auth_router"
	"github.com/user_service/internal/router/health_check"
)

type Router struct {
	Auth        *auth_router.AuthRouter
	Health      *health_check.HealthRouter
	Dispatcher  *event.Dispatcher
	LoginWorker *worker.LoginWorker
}

func NewRouter(
	Auth *auth_router.AuthRouter,
	Health *health_check.HealthRouter,
	Dispatcher *event.Dispatcher,
	LoginWorker *worker.LoginWorker,
) *Router {
	return &Router{
		Auth:        Auth,
		Health:      Health,
		Dispatcher:  Dispatcher,
		LoginWorker: LoginWorker,
	}
}

var RouterGroupApp *Router
