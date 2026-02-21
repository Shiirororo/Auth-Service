package auth_router

import (
	"os"

	"github.com/auth_service/internal/middleware"
	"github.com/auth_service/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type AuthRouter struct {
	redisClient *redis.Client
}

func NewAuthRouter(redisClient *redis.Client) *AuthRouter {
	return &AuthRouter{
		redisClient: redisClient,
	}
}

func (ar *AuthRouter) InitAuthRouter(Router *gin.RouterGroup) {
	//public router
	AUTH := Router.Group("/auth")
	{
		AUTH.GET("/login")
		AUTH.POST("/register")
	}
	jwtService := service.NewJWTService(os.Getenv("JWT_SECRET"))
	blacklist := service.NewRedisBlacklist(ar.redisClient)

	authMiddleware := middleware.NewAuthMiddleware(
		jwtService,
		blacklist,
	)

	//private router
	REQUIRE := Router.Group("/auth")
	REQUIRE.Use(authMiddleware.AuthenticateToken())
}
