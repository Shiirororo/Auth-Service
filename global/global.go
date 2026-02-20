package global

import (
	"github.com/auth_service/pkg/logger"
	"github.com/auth_service/pkg/settings"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Config settings.Config
	Logger *logger.LoggerZap
	Rdb    *redis.Client
	DB     *gorm.DB
)

/*
Config


*/
