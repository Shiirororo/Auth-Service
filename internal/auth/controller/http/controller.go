package http

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/user_service/internal/auth/application/service"
	"github.com/user_service/internal/auth/controller/dto"
	"github.com/user_service/pkg/token"
)

type AuthHandler struct {
	authService service.AuthServiceInterface
}

func NewAuthHandler(authService service.AuthServiceInterface) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) LoginHandler(c *gin.Context) {

	var req = dto.LoginRequestWithEmail{}

	if err := c.ShouldBindJSON(&req); err != nil {
		//TODO: Hide server status
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	accessToken, refreshToken, userID, err := h.authService.LoginServiceWithEmail(c.Request.Context(), req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user_id":       userID,
		"refresh_token": refreshToken,
		"message":       "login success",
		"access_token":  accessToken,
	})
}
func (h *AuthHandler) LogoutHandler(c *gin.Context) {
	val, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "claims not found"})
		return
	}

	claims, ok := val.(*token.Claims)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid claims type"})
		return
	}

	ttl := time.Until(claims.ExpiresAt.Time)

	err := h.authService.LogoutService(c.Request.Context(), claims.SessionID, ttl)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Logout success",
	})

}

func (h *AuthHandler) RefreshHandler(c *gin.Context) {
	var req = dto.RefreshRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	accessToken, refreshToken, err := h.authService.RefreshService(c.Request.Context(), req.RefreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

