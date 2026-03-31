package product_router

import (
	"github.com/gin-gonic/gin"
)

type ProductRouter struct {
}

func NewProductRouter() *ProductRouter {
	return &ProductRouter{}
}

func (or *ProductRouter) InitOrderRouter(Router *gin.RouterGroup) {
	//public
	product := Router.Group("/order")
	{
		product.GET("/get-product")
	}

}
