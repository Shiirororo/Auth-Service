package router

import (
	auth_router "github.com/user_service/internal/auth/controller"
	health_router "github.com/user_service/internal/health/controller"
	product_router "github.com/user_service/internal/product/controller"
	user_router "github.com/user_service/internal/user/controller"
)

type Router struct {
	Auth    *auth_router.AuthRouter
	Health  *health_router.HealthRouter
	User    *user_router.UserRouter
	Product *product_router.ProductRouter
}

func NewRouter(
	Auth *auth_router.AuthRouter,
	Health *health_router.HealthRouter,
	User *user_router.UserRouter,
	Product *product_router.ProductRouter,
) *Router {
	return &Router{
		Auth:    Auth,
		Health:  Health,
		User:    User,
		Product: Product,
	}
}

var RouterGroupApp *Router
