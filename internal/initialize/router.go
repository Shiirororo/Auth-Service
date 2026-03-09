package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/user_service/internal/middleware"
	"github.com/user_service/internal/router"
	"github.com/user_service/pkg/settings"
)

func InitRouter(r *gin.Engine, config settings.Config) {
	if config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	AuthRouter := router.RouterGroupApp.Auth
	HealthRouter := router.RouterGroupApp.Health
	UserRouter := router.RouterGroupApp.User
	r.Use(middleware.ConcurrencyLimiterHandler(config.Server.Max_Request))
	api := r.Group("/v1")
	go middleware.CleanUpClients()
	// api.Use(middleware.ConcurenncyLimiterHandler())
	{
		AuthRouter.InitAuthRouter(api) //<- MainGroup
		HealthRouter.InitHealthRouter(api)
		UserRouter.InitUserRouter(api)
	}
}

//Middleware use here:
//logging
//cross
//limiter global
