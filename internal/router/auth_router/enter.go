package auth_router

import (
	"github.com/gin-gonic/gin"
	"github.com/user_service/internal/handler"
	"github.com/user_service/internal/middleware"
)

type AuthRouter struct {
	authHandler    *handler.AuthHandler
	authMiddleware *middleware.AuthMiddleware
	rateLimit      *middleware.RateLimitMiddleware
}

func NewAuthRouter(authHandler *handler.AuthHandler, authMiddleware *middleware.AuthMiddleware, ratelimitMiddleware *middleware.RateLimitMiddleware) *AuthRouter {
	return &AuthRouter{
		authHandler:    authHandler,
		authMiddleware: authMiddleware,

		rateLimit: ratelimitMiddleware,
	}
}
func (ar *AuthRouter) InitAuthRouter(Router *gin.RouterGroup) {
	//public router
	auth := Router.Group("/auth")

	{
		auth.POST("/login", ar.rateLimit.UserLoginLimiter(), ar.authHandler.LoginHandler)
		auth.POST("/refresh-token", ar.authHandler.RefreshHandler)
		auth.POST("/register", ar.rateLimit.UserLoginLimiter(), ar.authHandler.RegisterHandler)
	}

	//private router
	private := auth.Group("/")
	private.Use(ar.authMiddleware.AuthenticateToken())
	{
		private.GET("/user_info" /*ar.Handler.GetUserInfoHandler*/)
		private.POST("/logout", ar.authHandler.LogoutHandler)
	}
}
