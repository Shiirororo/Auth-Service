package auth

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/user_service/pkg/request"
)

type AuthHandler struct {
	authService AuthServiceInterface
}

func NewAuthHandler(authService AuthServiceInterface) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) LoginHandler(c *gin.Context) {

	var req = request.LoginRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		//TODO: Hide server status
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	accessToken, refreshToken, err := h.authService.LoginService(c.Request.Context(), req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
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

	claims, ok := val.(*Claims)
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
	var req = request.RefreshRequest{}

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

func (h *AuthHandler) RegisterHandler(c *gin.Context) {
	var req = request.RegisterRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	err := h.authService.RegisterService(c.Request.Context(), req.Username, req.Password, req.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Registration successful",
	})
}
