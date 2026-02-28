package global

import (
	"github.com/redis/go-redis/v9"
	"github.com/user_service/pkg/logger"
	"github.com/user_service/pkg/settings"
	"gorm.io/gorm"
)

var (
	Config settings.Config
	Logger *logger.LoggerZap
	Rdb    *redis.Client // -> Blacklist
	DB     *gorm.DB
)

/*
Config


*/
