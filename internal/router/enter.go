package router

import (
	auth_router "github.com/user_service/internal/auth/controller"
	"github.com/user_service/internal/event"
	"github.com/user_service/internal/event/worker"
	health_router "github.com/user_service/internal/health/controller"
)

type Router struct {
	Auth           *auth_router.AuthRouter
	Health         *health_router.HealthRouter
	Dispatcher     *event.Dispatcher
	LoginWorker    *worker.LoginWorker
	RegisterWorker *worker.RegisterWorker
}

func NewRouter(
	Auth *auth_router.AuthRouter,
	Health *health_router.HealthRouter,
	Dispatcher *event.Dispatcher,
	LoginWorker *worker.LoginWorker,
	RegisterWorker *worker.RegisterWorker,
) *Router {
	return &Router{
		Auth:           Auth,
		Health:         Health,
		Dispatcher:     Dispatcher,
		LoginWorker:    LoginWorker,
		RegisterWorker: RegisterWorker,
	}
}

var RouterGroupApp *Router
