package initialize

import (
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/user_service/pkg/settings"
)

func InitRedis(config settings.Config) *redis.Client {
	r := config.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%v", r.Host, r.Port),
		Password: r.Password,
		DB:       r.Database,
		PoolSize: 10,
	})

	//TODO: Do check connection
	fmt.Println("Connected to redis")
	return rdb
}
