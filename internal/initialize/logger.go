package initialize

import (
	"github.com/user_service/pkg/logger"
	"github.com/user_service/pkg/settings"
)

func InitLogger(config settings.Config) *logger.LoggerZap {
	// Initialize logger here
	return logger.NewLogger(config.Logger)
}
