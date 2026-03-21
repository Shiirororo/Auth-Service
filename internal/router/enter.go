package router

import (
	auth_router "github.com/user_service/internal/auth/controller"
	health_router "github.com/user_service/internal/health/controller"
	user_router "github.com/user_service/internal/user/controller"
)

type Router struct {
	Auth   *auth_router.AuthRouter
	Health *health_router.HealthRouter
	User   *user_router.UserRouter
}

func NewRouter(
	Auth *auth_router.AuthRouter,
	Health *health_router.HealthRouter,
	User *user_router.UserRouter,
) *Router {
	return &Router{
		Auth:   Auth,
		Health: Health,
		User:   User,
	}
}

var RouterGroupApp *Router
