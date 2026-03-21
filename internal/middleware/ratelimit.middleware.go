package middleware

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"golang.org/x/time/rate"
)

var (
	mu        sync.Mutex
	clients   = make(map[string]*Client)
	semaphore chan struct{}
	once      sync.Once
)

type Client struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

// Key design:
// rl:<prefix>:ip:{ip}
// rl:<prefix>:user:{username}
// rl:<prefix>:ip_user:{ip}:{username}

func getLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err == nil {
		for _, address := range addrs {
			if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					return ipnet.IP.String()
				}
			}
		}
	}
	return "127.0.0.1"
}

func getClientIP(ctx *gin.Context) string {
	ip := ctx.ClientIP()
	switch ip {
	case "127.0.0.1", "::1":
		ip = getLocalIP()
	case "":
		ip = ctx.Request.RemoteAddr
	}
	return ip
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

/*
	Global ratelimit

*/

func ConcurrencyLimiterHandler(maxRequest int) gin.HandlerFunc {
	return func(c *gin.Context) {
		once.Do(func() {
			semaphore = make(chan struct{}, maxRequest)
		})

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
	AllowRequest(ctx context.Context, service_userid string, window time.Duration) (bool, error)
}

/*
Interface code
*/
func (r *RateLimitMiddleware) UpdateRequest(ctx context.Context, service string, ip string) error {
	key := fmt.Sprintf("ratelimit:%s:%s", service, ip)
	_, err := r.client.Incr(ctx, key).Result()
	return err
}
func (r *RateLimitMiddleware) AllowRequest(
	ctx context.Context,
	service_userid string,
	ip string,
	window time.Duration,
) (bool, error) {

	key := fmt.Sprintf("ratelimit:%s:%s", service_userid, ip)

	now := time.Now().UnixMilli()
	windowStart := now - window.Milliseconds()

	member := fmt.Sprintf("%d:%s", now, uuid.New().String())

	pipe := r.client.TxPipeline()

	pipe.ZRemRangeByScore(ctx, key, "0", fmt.Sprintf("%d", windowStart))

	pipe.ZAdd(ctx, key, redis.Z{
		Score:  float64(now),
		Member: member,
	})

	countCmd := pipe.ZCard(ctx, key)

	pipe.Expire(ctx, key, window)

	if _, err := pipe.Exec(ctx); err != nil {
		return false, err
	}

	count := countCmd.Val()

	if count > 10 {
		// rollback request
		r.client.ZRem(ctx, key, member)
		return false, nil
	}

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
func (r *RateLimitMiddleware) GetInforLimiter() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Next()
	}
}
