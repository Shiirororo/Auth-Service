package user_router

import (
	"github.com/gin-gonic/gin"
	"github.com/user_service/internal/middleware"
	user_http "github.com/user_service/internal/user/controller/http"
)

type UserRouter struct {
	userHandler    *user_http.UserHandler
	authMiddleware *middleware.AuthMiddleware
	rateLimit      *middleware.RateLimitMiddleware
}

func NewUserRouter(userHandler *user_http.UserHandler, authMiddleware *middleware.AuthMiddleware, ratelimitMiddleware *middleware.RateLimitMiddleware) *UserRouter {
	return &UserRouter{
		userHandler:    userHandler,
		authMiddleware: authMiddleware,
		rateLimit:      ratelimitMiddleware,
	}
}

func (ur *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	usr := Router.Group("/user")
	
	// public routes
	{
		usr.POST("/register", ur.rateLimit.UserLoginLimiter(), ur.userHandler.RegisterHandler)
	}

	// private routes
	privateUsr := usr.Group("/")
	privateUsr.Use(ur.authMiddleware.AuthenticateToken())
	{
		privateUsr.POST("/profile", ur.userHandler.GetUserInfoHandler)
	}
}
