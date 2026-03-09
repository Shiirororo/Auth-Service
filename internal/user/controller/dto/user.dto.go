package dto

import (
	"time"
)

type GetUserRequest struct {
	AccessToken string `json:"access_token" binding:"required"`
	UserID      string `json:"userID" binding:"required"`
}

type UserProfileResponse struct {
	UserID      string     `json:"user_id"`
	ProfileName string     `json:"profile_name"`
	Mobile      *string    `json:"mobile,omitempty"`
	Gender      *uint8     `json:"gender,omitempty"`
	Birthday    *time.Time `json:"birthday,omitempty"`
}
