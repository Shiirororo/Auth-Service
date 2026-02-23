package router

import (
	"github.com/auth_service/internal/router/auth_router"
)

type Router struct {
	Auth auth_router.AuthRouter
}

func NewRouter(Auth auth_router.AuthRouter) *Router {
	return &Router{
		Auth: Auth,
	}
}

var RouterGroupApp = new(Router)
