package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type RateLimitMiddleware struct {
	rdb *redis.Client
}

type RateLimitMiddlewareInterface interface {
	Allow(keyPrefix string, limit int, window time.Duration) gin.HandlerFunc
}

// Key design:
// rl:<prefix>:ip:{ip}
// rl:<prefix>:user:{username}
// rl:<prefix>:ip_user:{ip}:{username}

func NewRateLimitMiddleware(rdb *redis.Client) RateLimitMiddlewareInterface {
	return &RateLimitMiddleware{rdb: rdb}
}

func getClientIP(ctx *gin.Context) string {
	ip := ctx.ClientIP()

	return ip
}

func (r *RateLimitMiddleware) Allow(keyPrefix string, limit int, window time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		ip := getClientIP(c)

		// Attempt to get a user identifier from the context, if set by an auth middleware
		// You can customize the key "username" or "userID" based on your actual auth logic.
		username := c.GetString("user-id")

		var key string
		if username != "" {
			key = fmt.Sprintf("rl:%s:ip_user:%s:%s", keyPrefix, ip, username)
		} else {
			key = fmt.Sprintf("rl:%s:ip:%s", keyPrefix, ip)
		}

		count, err := r.rdb.Incr(ctx, key).Result()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Rate limiter error"})
			return
		}

		if count == 1 {
			// Set expiration window on the first request
			r.rdb.Expire(ctx, key, window)
		}

		if count > int64(limit) {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "Too Many Requests"})
			return
		}

		c.Next()
	}
}
