package auth_router

import (
	"github.com/auth_service/internal/handler"
	"github.com/auth_service/internal/service"
	"github.com/gin-gonic/gin"
)

type AuthRouter struct {
	RedisClient *service.RedisBlacklist
	Handler     *handler.AuthHandler
	Middleware  gin.HandlerFunc
}

func NewAuthRouter(redisClient *service.RedisBlacklist, Handler *handler.AuthHandler, Middleware gin.HandlerFunc) *AuthRouter {
	return &AuthRouter{
		RedisClient: redisClient,
		Handler:     Handler,
		Middleware:  Middleware,
	}
}
func (ar *AuthRouter) InitAuthRouter(Router *gin.RouterGroup) {
	//public router
	auth := Router.Group("/auth")
	//AUTH := Router.Group("/auth")
	{
		auth.POST("/login", ar.Handler.LoginHandler)
	}

	//private router
	private := auth.Group("/")
	private.Use(ar.Middleware)
	{
		private.GET("/get_info")
	}
}
