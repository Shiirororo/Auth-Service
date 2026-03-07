package initialize

import (
	"log"
	"os"

	"github.com/user_service/internal/auth"
)

func InitJWT() auth.JWTService {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Fatal("No secret found in env")
	}
	return auth.NewJWTSToken(secret)
}
