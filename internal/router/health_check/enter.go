package health_check

import (
	"github.com/auth_service/internal/handler"
	"github.com/gin-gonic/gin"
)

type HealthRouter struct {
	Handler *handler.HealthHandler
}

func NewHealthRouter(h *handler.HealthHandler) *HealthRouter {
	return &HealthRouter{Handler: h}
}

func (hr *HealthRouter) InitHealthRouter(Router *gin.RouterGroup) {
	Router.GET("/health", hr.Handler.Ping)
}
