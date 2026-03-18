package user_router

import (
	"github.com/gin-gonic/gin"
	"github.com/user_service/internal/middleware"
	user_http "github.com/user_service/internal/user/controller/http"
)

type UserRouter struct {
	userHandler    *user_http.UserHandler
	authMiddleware *middleware.AuthMiddleware
}

func NewUserRouter(userHandler *user_http.UserHandler, authMiddleware *middleware.AuthMiddleware) *UserRouter {
	return &UserRouter{
		userHandler:    userHandler,
		authMiddleware: authMiddleware,
	}
}

func (ur *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	usr := Router.Group("/user")
	usr.Use(ur.authMiddleware.AuthenticateToken())
	{
		usr.POST("/profile", ur.userHandler.GetUserInfoHandler)
	}
}
