package auth_router

import (
	"github.com/auth_service/internal/handler"
	"github.com/auth_service/internal/middleware"
	"github.com/gin-gonic/gin"
)

type AuthRouter struct {
	Handler    *handler.AuthHandler
	Middleware *middleware.AuthMiddleware
}

func NewAuthRouter(Handler *handler.AuthHandler, Middleware *middleware.AuthMiddleware) *AuthRouter {
	return &AuthRouter{
		Handler:    Handler,
		Middleware: Middleware,
	}
}
func (ar *AuthRouter) InitAuthRouter(Router *gin.RouterGroup) {
	//public router
	auth := Router.Group("/auth")
	//AUTH := Router.Group("/auth")
	{
		auth.POST("/login", ar.Handler.LoginHandler)
		auth.POST("/refresh-token", ar.Handler.RefreshHandler)
		auth.POST("/register", ar.Handler.RegisterHandler)
	}

	//private router
	private := auth.Group("/")
	private.Use(ar.Middleware.AuthenticateToken())
	{
		private.GET("/get_info")
		private.POST("/logout", ar.Handler.LogoutHandler)
	}
}
