package health_check

import (
	"github.com/gin-gonic/gin"
	"github.com/user_service/internal/handler"
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
