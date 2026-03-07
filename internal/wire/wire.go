//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
	"github.com/user_service/internal/auth"
	"github.com/user_service/internal/event"
	"github.com/user_service/internal/event/worker"
	"github.com/user_service/internal/health"
	"github.com/user_service/internal/initialize"
	"github.com/user_service/internal/middleware"
	"github.com/user_service/internal/router"
	"github.com/user_service/internal/router/auth_router"
	"github.com/user_service/internal/router/health_check"
	"gorm.io/gorm"
)

func provideEventQueue() chan event.Event {
	return make(chan event.Event, 1000)
}

func provideWorkerCount() int {
	return 5
}

func InitRouter(db *gorm.DB, rdb *redis.Client) (*router.Router, error) {
	wire.Build(
		provideEventQueue,
		provideWorkerCount,
		event.NewDispatcher,
		worker.NewLoginWorker,
		auth.NewAuthRepository,
		initialize.InitJWT,
		auth.NewRedisBlacklist,
		auth.NewAuthService,
		auth.NewAuthHandler,
		middleware.NewAuthMiddleware,
		middleware.NewRateLimitMiddleware,
		auth_router.NewAuthRouter,
		health.NewHealthHandler,
		health_check.NewHealthRouter,
		router.NewRouter,
	)
	return new(router.Router), nil
}
