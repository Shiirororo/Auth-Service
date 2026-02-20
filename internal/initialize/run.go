package initialize

import (
	"fmt"

	"github.com/auth_service/global"
)

func Run() {
	LoadConfig()
	m := global.Config.Databases
	fmt.Println("Loading MySQL configuration", m.Username, m.Password)
	InitLogger()
	InitMySQL()
	InitRedis()
	r := InitRouter()
	r.Run(":8002")
}
