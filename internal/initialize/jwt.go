package initialize

import (
	"log"
	"os"

	"github.com/auth_service/internal/service"
)

func InitJWT() service.JWTService {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Fatal("No secret found in env")
	}
	return service.NewJWTSToken(secret)
}
