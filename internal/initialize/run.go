package initialize

import (
	"fmt"
	"os"
	"strconv"

	"github.com/auth_service/global"
	"github.com/joho/godotenv"
)

func LoadEnv() {
	_ = godotenv.Load()

	if os.Getenv("JWT_SECRET") == "" {
		return
	}
}

func Run() {

	LoadConfig()
	m := global.Config.Databases
	fmt.Println("Loading MySQL configuration", m.Username, m.Password)
	InitLogger()
	LoadEnv()
	InitMySQL()
	InitRedis()
	r := InitRouter()
	r.Run(":" + strconv.Itoa(global.Config.Server.Port))
	fmt.Println("Server is running at port: ")
}
