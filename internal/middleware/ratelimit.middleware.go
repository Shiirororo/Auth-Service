package middleware

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/user_service/global"
	"golang.org/x/time/rate"
)

var (
	mu        sync.Mutex
	clients   = make(map[string]*Client)
	semaphore = make(chan struct{}, global.Config.Server.Max_Request) //Make it changable through setting
)

type Client struct {
	limiter  *rate.Limiter
	lastSeen time.Time
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
		time.Sleep(time.Minute)
		mu.Lock()
		for ip, client := range clients {
			if time.Since(client.lastSeen) > 3*time.Minute {
				delete(clients, ip)
			}
		}
		mu.Unlock()
	}
}

func (r *RateLimitMiddleware) RateLimiter() gin.HandlerFunc {
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

/*
	Global ratelimit

*/

func ConcurrencyLimiterHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		select {
		case semaphore <- struct{}{}:
			defer func() { <-semaphore }()
			c.Next()
		default:
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"error": "Max concurrency request",
			})
			c.Abort()
		}
	}
}

/*
User level rate limit

format: ratelimit:<service>:<userID>/<userIP>:<count>

- Login: MAX 10 request/min per user
*/
type RateLimitMiddleware struct {
	client *redis.Client
}
type RateLimitInterface interface {
	UpdateRequest(ctx context.Context, service string, ip string) error
}

/*
Interface code
*/
func (r *RateLimitMiddleware) UpdateRequest(ctx context.Context, service string, ip string) error {
	key := fmt.Sprintf("ratelimit:%s:%s", service, ip)
	_, err := r.client.Incr(ctx, key).Result()
	return err
}
func (r *RateLimitMiddleware) AllowRequest(ctx context.Context, service string, ip string, window time.Duration) (bool, error) {
	key := fmt.Sprintf("ratelimit:%s:%s", service, ip)
	now := time.Now().UnixMilli()
	windowStart := now - window.Milliseconds()

	// Remove old request count
	err := r.client.ZRemRangeByScore(ctx, key, "0", fmt.Sprintf("%d", windowStart)).Err()

	if err != nil {
		return false, err
	}

	//Count

	count, err := r.client.ZCard(ctx, key).Result()
	if err != nil {
		return false, err
	}
	if count >= 10 {
		return false, nil
	}
	//Update
	err = r.client.ZAdd(ctx, key, redis.Z{
		Score:  float64(now),
		Member: now,
	}).Err()
	if err != nil {
		return false, err
	}

	r.client.Expire(ctx, key, window)
	return true, nil
}

func NewRateLimitMiddleware(limiter *redis.Client) *RateLimitMiddleware {
	return &RateLimitMiddleware{
		client: limiter,
	}
}

func (r *RateLimitMiddleware) UserLoginLimiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := getClientIP(c)
		allow, err := r.AllowRequest(c.Request.Context(), "login", ip, time.Minute)
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
			return
		}
		if !allow {
			c.AbortWithStatusJSON(429, gin.H{
				"error": http.StatusTooManyRequests,
			})
			return
		}

		c.Next()
	}
}
