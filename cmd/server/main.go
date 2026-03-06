package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/user_service/global"
	"github.com/user_service/internal/initialize"
	"github.com/user_service/internal/router"
	"github.com/user_service/internal/wire"
)

func main() {
	//go console.Console()
	initialize.Run()

	routerApp, err := wire.InitRouter(global.DB, global.Rdb)
	if err != nil {
		fmt.Printf("Failed to initialize router: %v\n", err)
		return
	}

	router.RouterGroupApp = routerApp

	r := gin.New()
	initialize.InitRouter(r)
	fmt.Println("Server is running at port: ", global.Config.Server.Port)
	fmt.Println("Auth Service is running...")
	r.Run(":" + strconv.Itoa(global.Config.Server.Port))
}
