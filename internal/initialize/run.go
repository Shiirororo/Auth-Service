package initialize

import (
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/user_service/global"
	"gorm.io/gorm"
)

func Run() (*gorm.DB, *redis.Client) {

	LoadConfig()
	m := global.Config.Databases
	fmt.Println("Loading MySQL configuration", m.Username, m.Password)
	InitLogger()
	db := InitMySQL()
	rdb := InitRedis()
	InitJWT()
	fmt.Println("Successfully Initialized ")
	return db, rdb
}
