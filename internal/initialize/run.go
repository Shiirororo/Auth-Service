package initialize

import (
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/user_service/pkg/logger"
	"github.com/user_service/pkg/settings"
	"gorm.io/gorm"
)

func Run() (*gorm.DB, *redis.Client, settings.Config, *logger.LoggerZap) {

	config := LoadConfig()
	// Set global config temporarily for InitMySQL and InitRedis until we inject them
	// We will inject them manually since we aren't using AppContext
	m := config.Databases
	fmt.Println("Loading MySQL configuration", m.Username, m.Password)
	log := InitLogger(config)
	db := InitMySQL(config)
	rdb := InitRedis(config)
	InitJWT()
	fmt.Println("Successfully Initialized")
	return db, rdb, config, log
}
