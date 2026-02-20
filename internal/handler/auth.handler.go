package handler

import (
	"net/http"
	// "time"

	"github.com/auth_service/internal/service"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService *service.AuthService
	jwtService  *service.JWTService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginCredit struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (h *AuthHandler) LoginHandler(c *gin.Context) {

	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		//TODO: Hide server status
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	accessToken, refreshToken, err := h.authService.LoginService(c.Request.Context(), req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":       "login success",
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

func (h *AuthHandler) AuthenticateHandler(c *gin.Context) {
	//Im not even write it :(
	var req LoginCredit

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

}

// LogoutHandler handles user logout
// func (h *AuthHandler) LogoutHandler(c *gin.Context) {
// 	// Extract Access Token from header to blacklist it
// 	tokenString := c.GetHeader("Authorization")
// 	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
// 		tokenString = tokenString[7:]
// 	}

// 	userID, exists := service.GetUserID(c)
// 	if !exists {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
// 		return
// 	}

// 	// Verify again to get claims (ID, Exp)
// 	claims, err := h.jwtService.VerifyJWT(c.Request.Context(), tokenString, nil)
// 	if err != nil {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
// 		return
// 	}

// 	if claims.ExpiresAt == nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid token claims"})
// 		return
// 	}

// 	// Calculate TTL
// 	ttl := time.Until(claims.ExpiresAt.Time)

// 	err = h.authService.LogoutService(c.Request.Context(), claims.ID, ttl, userID)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "logout failed"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "logout success"})
// }

// RefreshHandler handles token refresh
func (h *AuthHandler) RefreshHandler(c *gin.Context) {
	var req struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	accessToken, refreshToken, err := h.authService.RefreshTokenService(c.Request.Context(), req.RefreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}
