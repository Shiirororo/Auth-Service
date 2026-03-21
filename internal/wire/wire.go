//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
	"github.com/user_service/internal/auth/application/service"
	auth_worker "github.com/user_service/internal/auth/application/worker"
	auth_router "github.com/user_service/internal/auth/controller"
	auth_http "github.com/user_service/internal/auth/controller/http"
	"github.com/user_service/internal/auth/infrastructure/persistence"
	"github.com/user_service/internal/event"
	health_router "github.com/user_service/internal/health/controller"
	health_http "github.com/user_service/internal/health/controller/http"
	"github.com/user_service/internal/initialize"
	"github.com/user_service/internal/middleware"
	"github.com/user_service/internal/router"
	user_service "github.com/user_service/internal/user/application/service"
	user_worker "github.com/user_service/internal/user/application/worker"
	user_router "github.com/user_service/internal/user/controller"
	user_http "github.com/user_service/internal/user/controller/http"
	user_persistence "github.com/user_service/internal/user/infrastrucutre/persistence"
	"gorm.io/gorm"
)

func provideEventQueue() chan event.Event {
	return make(chan event.Event, 1000)
}

func provideWorkerCount() int {
	return 5
}

func InitApp(db *gorm.DB, rdb *redis.Client) (*router.App, error) {
	wire.Build(
		provideEventQueue,
		provideWorkerCount,
		event.NewDispatcher,
		auth_worker.NewLoginWorker,
		user_worker.NewRegisterWorker,
		persistence.NewAuthRepository,
		persistence.NewUserRepository,
		persistence.NewRoleRepository,
		user_persistence.NewProfileRepository,
		persistence.NewRedisOTPRepository,
		initialize.InitJWT,
		service.NewRedisBlacklist,
		service.NewAuthorizationService,
		service.NewAuthService,
		auth_http.NewAuthHandler,
		middleware.NewAuthMiddleware,
		middleware.NewRateLimitMiddleware,
		auth_router.NewAuthRouter,
		health_http.NewHealthHandler,
		health_router.NewHealthRouter,
		user_service.NewUserService,
		user_http.NewUserHandler,
		user_router.NewUserRouter,
		router.NewRouter,
		router.NewApp,
	)
	return new(router.App), nil
}
