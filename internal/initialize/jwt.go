package initialize

import (
	"log"
	"os"

	"github.com/user_service/pkg/token"
)

func InitJWT() token.TokenMaker {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		log.Fatal("No secret found in env")
	}
	return token.NewJWTMaker(secret)
}
