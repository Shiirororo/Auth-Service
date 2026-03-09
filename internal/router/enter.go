package router

import (
	"github.com/user_service/internal/auth/application/worker"
	auth_router "github.com/user_service/internal/auth/controller"
	"github.com/user_service/internal/event"
	health_router "github.com/user_service/internal/health/controller"
	user_router "github.com/user_service/internal/user/controller"
)

type Router struct {
	Auth           *auth_router.AuthRouter
	Health         *health_router.HealthRouter
	User           *user_router.UserRouter
	Dispatcher     *event.Dispatcher
	LoginWorker    *worker.LoginWorker
	RegisterWorker *worker.RegisterWorker
}

func NewRouter(
	Auth *auth_router.AuthRouter,
	Health *health_router.HealthRouter,
	User *user_router.UserRouter,
	Dispatcher *event.Dispatcher,
	LoginWorker *worker.LoginWorker,
	RegisterWorker *worker.RegisterWorker,
) *Router {
	return &Router{
		Auth:           Auth,
		Health:         Health,
		User:           User,
		Dispatcher:     Dispatcher,
		LoginWorker:    LoginWorker,
		RegisterWorker: RegisterWorker,
	}
}

var RouterGroupApp *Router
