package middleware

import (
	"github.com/gin-gonic/gin"
	auth_service "github.com/user_service/internal/auth/application/service"
	"github.com/user_service/pkg/token"
)

type AuthMiddleware struct {
	authService token.TokenMaker
	tokenRepo   auth_service.TokenBlacklist
	authz       auth_service.AuthorizationServiceInterface
}

func NewAuthMiddleware(authService token.TokenMaker, tokenRepo auth_service.TokenBlacklist, authz auth_service.AuthorizationServiceInterface) *AuthMiddleware {
	return &AuthMiddleware{
		authService: authService,
		tokenRepo:   tokenRepo,
		authz:       authz,
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

func (m *AuthMiddleware) AuthorizationUser(requiredRoleID uint) gin.HandlerFunc {
	return func(c *gin.Context) {
		val, exists := c.Get("claims")
		if !exists {
			abortUnauthorized(c, "missing claims")
			return
		}
		claims, ok := val.(*token.Claims)
		if !ok || claims == nil {
			abortUnauthorized(c, "invalid claims")
			return
		}

		userIDBytes := []byte(claims.UserID) // Normally parsed back to bytes
		
		hasRole, err := m.authz.CheckUserRole(c.Request.Context(), userIDBytes, requiredRoleID)
		if err != nil || !hasRole {
			abortUnauthorized(c, "insufficient permissions")
			return
		}

		c.Next()
	}
}
