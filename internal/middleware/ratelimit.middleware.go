package middleware

import "github.com/gin-gonic/gin"

type RateLimitMiddleware struct {
}

func (r *RateLimitMiddleware) NewRateLimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
