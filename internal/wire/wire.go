//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/auth_service/internal/handler"
	"github.com/auth_service/internal/initialize"
	"github.com/auth_service/internal/middleware"
	"github.com/auth_service/internal/repository"
	"github.com/auth_service/internal/router"
	"github.com/auth_service/internal/router/auth_router"
	"github.com/auth_service/internal/router/health_check"
	"github.com/auth_service/internal/service"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
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
		auth_router.NewAuthRouter,
		handler.NewHealthHandler,
		health_check.NewHealthRouter,
		router.NewRouter,
	)
	return new(router.Router), nil
}
