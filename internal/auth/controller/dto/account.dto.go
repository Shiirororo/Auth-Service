package dto

import "github.com/user_service/internal/auth/domain/vo"

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

type LoginRequestWithUsername struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type LoginRequestWithEmail struct {
	Email    vo.Email `json:"email" binding:"required"`
	Password string   `json:"password" binding:"required"`
}

type UserLoginResponse struct {
	UserID string `json:"user_id"`
	Data   struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
		TTL          string `json:"ttl"`
		AccountInfo  struct {
		} `json:"accountInfo"`
	} `json:"data"`
}
type RefreshRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type LogoutRequest struct {
	SessionID    string `json:"session_id" binding:"required"`
	RefreshToken string `json:"refresh_token" binding:"required"`
}
