package order_router

import (
	"github.com/gin-gonic/gin"
)

type OrderRouter struct {
}

func (or *OrderRouter) InitOrderRouter(Router *gin.RouterGroup) {
	//public
	order := Router.Group("/order")
	{
		order.GET("/get-product")
	}

	//private router
	private := order.Group("/")
	{
		private.POST("/payment")
	}
}
