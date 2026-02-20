package initialize

import (
	// "github.com/auth_service/global"
	// "github.com/auth_service/internal/handler"
	// "github.com/auth_service/internal/repository"
	// "github.com/auth_service/internal/service"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	// authRepo := repository.NewAuthRepository(global.DB)
	// authService := service.NewAuthService(authRepo, to)
	// authHandler := handler.NewAuthHandler(authService)

	// authGroup := r.Group("/auth")
	// {
	// 	authGroup.POST("/login", authHandler.LoginHandler)
	// 	authGroup.POST("/refresh", authHandler.RefreshHandler)
	// 	// authGroup.POST("/logout", authHandler.LogoutHandler) // Should ideally use auth middleware, but handler checks token too
	// }

	return r
}
