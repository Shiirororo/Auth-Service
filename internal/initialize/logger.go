package initialize

import (
	"github.com/auth_service/global"
	"github.com/auth_service/pkg/logger"
)

func InitLogger() {
	// Initialize logger here
	global.Logger = logger.NewLogger(global.Config.Logger)
}
