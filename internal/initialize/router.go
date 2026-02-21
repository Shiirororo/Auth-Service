package initialize

import (
	// "github.com/auth_service/global"
	// "github.com/auth_service/internal/handler"
	// "github.com/auth_service/internal/repository"
	// "github.com/auth_service/internal/service"
	"github.com/auth_service/global"
	"github.com/auth_service/internal/router"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	var r *gin.Engine
	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}
	AuthRouter := router.RouterGroupApp.Auth

	MainGroup := r.Group("v1")
	{
		AuthRouter.InitAuthRouter(MainGroup)
	}
	return r

	//Middleware use here:
	//logging
	//cross
	//limiter global
}
