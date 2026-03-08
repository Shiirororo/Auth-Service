//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
	"github.com/user_service/internal/auth/application/service"
	auth_router "github.com/user_service/internal/auth/controller"
	auth_http "github.com/user_service/internal/auth/controller/http"
	"github.com/user_service/internal/auth/infrastructure/messaging"
	"github.com/user_service/internal/auth/infrastructure/persistence"
	"github.com/user_service/internal/commons"
	"github.com/user_service/internal/event"
	"github.com/user_service/internal/event/worker"
	health_router "github.com/user_service/internal/health/controller"
	health_http "github.com/user_service/internal/health/controller/http"
	"github.com/user_service/internal/initialize"
	"github.com/user_service/internal/middleware"
	"github.com/user_service/internal/router"
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
		persistence.NewUserRepository,
		persistence.NewRedisOTPRepository,
		messaging.NewMockEmailSender,
		initialize.InitJWT,
		commons.NewRedisBlacklist,
		service.NewAuthService,
		auth_http.NewAuthHandler,
		middleware.NewAuthMiddleware,
		middleware.NewRateLimitMiddleware,
		auth_router.NewAuthRouter,
		health_http.NewHealthHandler,
		health_router.NewHealthRouter,
		router.NewRouter,
	)
	return new(router.Router), nil
}
