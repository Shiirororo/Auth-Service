package health_check

import (
	"github.com/gin-gonic/gin"
	"github.com/user_service/internal/health"
)

type HealthRouter struct {
	HealthHandler *health.HealthHandler
}

func NewHealthRouter(healthHandler *health.HealthHandler) *HealthRouter {
	return &HealthRouter{HealthHandler: healthHandler}
}

func (hr *HealthRouter) InitHealthRouter(Router *gin.RouterGroup) {
	Router.GET("/health", hr.HealthHandler.Ping)
}
