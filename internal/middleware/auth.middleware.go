package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/user_service/internal/commons"
	"github.com/user_service/pkg/token"
)

type AuthMiddleware struct {
	authService token.TokenMaker
	tokenRepo   commons.TokenBlacklist
}

func NewAuthMiddleware(authService token.TokenMaker, tokenRepo commons.TokenBlacklist) *AuthMiddleware {
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

		// Check Blacklist
		isSessionBlocked, err := m.tokenRepo.IsSessionBlacklisted(reqCtx, claims.SessionID)
		if err != nil {
			abortUnauthorized(c, err.Error())
			return
		}
		if isSessionBlocked {
			abortUnauthorized(c, "Session has been revoked")
			return
		}

		isJTIBlocked, _ := m.tokenRepo.IsJTIBlacklisted(reqCtx, claims.ID)
		if isJTIBlocked {
			abortUnauthorized(c, "Token has been revoked")
			return
		}
		c.Set("claims", claims)
		c.Next()
	}
}

func (m *AuthMiddleware) AuthorizationUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		//TODO: Authorization Middlware
	}
}
