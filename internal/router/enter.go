package router

import (
	"github.com/auth_service/internal/router/auth_router"
)

type Router struct {
	Auth auth_router.AuthRouter
}

var RouterGroupApp = new(Router)
