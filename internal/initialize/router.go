package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/user_service/global"
	"github.com/user_service/internal/router"
)

func InitRouter(r *gin.Engine) {
	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	AuthRouter := router.RouterGroupApp.Auth
	HealthRouter := router.RouterGroupApp.Health

	api := r.Group("/v1")

	{
		AuthRouter.InitAuthRouter(api) //<- MainGroup
		HealthRouter.InitHealthRouter(api)
	}

}

//Middleware use here:
//logging
//cross
//limiter global
