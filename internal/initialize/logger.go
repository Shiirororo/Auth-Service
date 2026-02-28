package initialize

import (
	"github.com/user_service/global"
	"github.com/user_service/pkg/logger"
)

func InitLogger() {
	// Initialize logger here
	global.Logger = logger.NewLogger(global.Config.Logger)
}
