package main

import (
	"fmt"
	"os"
	"strconv"

	//"github.com/auth_service/cmd/console"
	"github.com/auth_service/global"
	"github.com/auth_service/internal/handler"
	"github.com/auth_service/internal/initialize"
	"github.com/auth_service/internal/middleware"
	"github.com/auth_service/internal/repository"
	"github.com/auth_service/internal/router"
	"github.com/auth_service/internal/router/auth_router"
	"github.com/auth_service/internal/router/health_check"
	"github.com/auth_service/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
	//go console.Console()
	initialize.Run()
	blacklist := service.NewRedisBlacklist(global.Rdb)
	jwtService := service.NewJWTSToken(os.Getenv("JWT_SECRET"))
	auth_repo := repository.NewAuthRepository(global.DB)
	auth_service := service.NewAuthService(auth_repo, blacklist, jwtService)
	auth_handler := handler.NewAuthHandler(*auth_service)

	auth_middleware := middleware.NewAuthMiddleware(
		jwtService,
		blacklist,
	)
	authRouter := auth_router.NewAuthRouter(blacklist, auth_handler, auth_middleware.AuthenticateToken())

	healthHandler := handler.NewHealthHandler()
	healthRouter := health_check.NewHealthRouter(healthHandler)

	router.RouterGroupApp = router.NewRouter(authRouter, healthRouter)
	r := gin.New()
	initialize.InitRouter(r)
	r.Run(":" + strconv.Itoa(global.Config.Server.Port))
	fmt.Println("Server is running at port: ")
	fmt.Println("Auth Service is running...")
}
