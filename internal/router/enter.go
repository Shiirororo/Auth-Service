package router

import (
	"github.com/auth_service/internal/router/auth_router"
	"github.com/auth_service/internal/router/health_check"
)

type Router struct {
	Auth   *auth_router.AuthRouter
	Health *health_check.HealthRouter
}

func NewRouter(Auth *auth_router.AuthRouter, Health *health_check.HealthRouter) *Router {
	return &Router{
		Auth:   Auth,
		Health: Health,
	}
}

var RouterGroupApp *Router
