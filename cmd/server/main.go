package main

import (
	"fmt"
	"os"

	//"github.com/auth_service/cmd/console"
	"github.com/auth_service/global"
	"github.com/auth_service/internal/handler"
	"github.com/auth_service/internal/initialize"
	"github.com/auth_service/internal/middleware"
	"github.com/auth_service/internal/repository"
	"github.com/auth_service/internal/router/auth_router"
	"github.com/auth_service/internal/service"
	// "github.com/auth_service/internal/repository"
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
	fmt.Println("Auth Service is running...")
}
