package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/user_service/internal/initialize"
	"github.com/user_service/internal/router"
	"github.com/user_service/internal/wire"
)

func main() {
	//go console.Console()
	db, rdb, config, _ := initialize.Run()

	app, err := wire.InitApp(db, rdb)
	if err != nil {
		fmt.Printf("Failed to initialize app: %v\n", err)
		return
	}

	router.RouterGroupApp = app.Router

	ctx := context.Background()
	go app.Dispatcher.Start(ctx)
	go app.LoginWorker.Start(ctx)

	r := gin.New()
	initialize.InitRouter(r, config)
	fmt.Println("Server is running at port: ", config.Server.Port)
	fmt.Println("Auth Service is running...")
	r.Run(":" + strconv.Itoa(config.Server.Port))
}
