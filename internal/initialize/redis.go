package initialize

import (
	"fmt"

	"github.com/auth_service/global"
	"github.com/redis/go-redis/v9"
)

func InitRedis() {
	r := global.Config.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%v", r.Host, r.Port),
		Password: r.Password,
		DB:       r.Database,
		PoolSize: 10,
	})

	//TODO: Do check connection
	global.Rdb = rdb
	fmt.Println("Connected to redis")

}
