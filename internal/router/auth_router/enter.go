package auth_router

import (
	"github.com/auth_service/internal/handler"
	"github.com/auth_service/internal/middleware"
	"github.com/gin-gonic/gin"
)

type AuthRouter struct {
	authHandler    *handler.AuthHandler
	authMiddleware *middleware.AuthMiddleware
}

func NewAuthRouter(authHandler *handler.AuthHandler, authMiddleware *middleware.AuthMiddleware) *AuthRouter {
	return &AuthRouter{
		authHandler:    authHandler,
		authMiddleware: authMiddleware,
	}
}
func (ar *AuthRouter) InitAuthRouter(Router *gin.RouterGroup) {
	//public router
	auth := Router.Group("/auth")

	{
		auth.POST("/login", ar.authHandler.LoginHandler)
		auth.POST("/refresh-token", ar.authHandler.RefreshHandler)
		auth.POST("/register", ar.authHandler.RegisterHandler)
	}

	//private router
	private := auth.Group("/")
	private.Use(ar.authMiddleware.AuthenticateToken())
	{
		private.GET("/user_info" /*ar.Handler.GetUserInfoHandler*/)
		private.POST("/logout", ar.authHandler.LogoutHandler)
	}
}
