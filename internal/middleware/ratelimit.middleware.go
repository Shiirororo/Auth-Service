package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

var (
	mu      sync.Mutex
	clients = make(map[string]*Client)
)

type Client struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}
type RateLimitMiddleware struct {
}

// Key design:
// rl:<prefix>:ip:{ip}
// rl:<prefix>:user:{username}
// rl:<prefix>:ip_user:{ip}:{username}

func getClientIP(ctx *gin.Context) string {
	ip := ctx.ClientIP()
	if ip == "" {
		ip = ctx.Request.RemoteAddr
	}
	return ip
}
func getRateLimiter(ip string) *rate.Limiter {
	mu.Lock()
	client, exist := clients[ip]
	if !exist {
		limiter := rate.NewLimiter(5, 10)
		newClient := &Client{limiter, time.Now()}
		clients[ip] = newClient

		return limiter
	}
	client.lastSeen = time.Now()
	mu.Unlock()
	return client.limiter
}

func CleanUpClients() {
	for {

	}
}

func RateLimiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := getClientIP(c)

		limiter := getRateLimiter(ip)

		if !limiter.Allow() {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error": "Too many request",
			})
		}
	}
}
