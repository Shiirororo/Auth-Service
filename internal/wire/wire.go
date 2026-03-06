//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
	"github.com/user_service/internal/handler"
	"github.com/user_service/internal/initialize"
	"github.com/user_service/internal/middleware"
	"github.com/user_service/internal/repository"
	"github.com/user_service/internal/router"
	"github.com/user_service/internal/router/auth_router"
	"github.com/user_service/internal/router/health_check"
	"github.com/user_service/internal/service"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, rdb *redis.Client) (*router.Router, error) {
	wire.Build(
		repository.NewAuthRepository,
		initialize.InitJWT,
		service.NewRedisBlacklist,
		service.NewAuthService,
		handler.NewAuthHandler,
		middleware.NewAuthMiddleware,
		middleware.NewRateLimitMiddleware,
		auth_router.NewAuthRouter,
		handler.NewHealthHandler,
		health_check.NewHealthRouter,
		router.NewRouter,
	)
	return new(router.Router), nil
}
