package health_router

import (
	"github.com/gin-gonic/gin"
	"github.com/user_service/internal/health/controller/http"
)

type HealthRouter struct {
	HealthHandler *http.HealthHandler
}

func NewHealthRouter(healthHandler *http.HealthHandler) *HealthRouter {
	return &HealthRouter{HealthHandler: healthHandler}
}

func (hr *HealthRouter) InitHealthRouter(Router *gin.RouterGroup) {
	Router.GET("/health", hr.HealthHandler.Ping)
}
