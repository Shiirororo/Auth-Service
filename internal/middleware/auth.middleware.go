package middleware

import (
	"github.com/auth_service/internal/service"
	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	authService service.JWTService
	tokenRepo   service.TokenBlacklist
}

func NewAuthMiddleware(authService service.JWTService, tokenRepo service.TokenBlacklist) *AuthMiddleware {
	return &AuthMiddleware{
		authService: authService,
		tokenRepo:   tokenRepo,
	}
}

func (m *AuthMiddleware) AuthenticateToken() gin.HandlerFunc { //Verify Signature from JWTService
	return func(c *gin.Context) {
		tokenString, err := extractBearerToken(c, "Authorization")
		if err != nil {
			abortUnauthorized(c, err.Error())
			return
		}
		reqCtx := c.Request.Context()
		claims, err := m.authService.VerifyJWT(reqCtx, tokenString)
		if err != nil {
			abortUnauthorized(c, err.Error())
			return
		}

		isBlackListed, err := m.tokenRepo.IsBlacklisted(reqCtx, tokenString, "")

		if isBlackListed {
			abortUnauthorized(c, "Token has been revoked")
			return
		}
		c.Set("claims", claims)
		c.Next()
	}
}
