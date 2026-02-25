package auth_router

import (
	"github.com/auth_service/internal/handler"
	"github.com/gin-gonic/gin"
)

type AuthRouter struct {
	Handler    *handler.AuthHandler
	Middleware gin.HandlerFunc
}

func NewAuthRouter(Handler *handler.AuthHandler, Middleware gin.HandlerFunc) *AuthRouter {
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
	}

	//private router
	private := auth.Group("/")
	private.Use(ar.Middleware)
	{
		private.GET("/get_info")
		private.POST("/logout", ar.Handler.LogoutHandler)
	}
}
