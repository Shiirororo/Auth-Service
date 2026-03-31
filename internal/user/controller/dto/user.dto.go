package dto

import "github.com/user_service/internal/user/domain/model/entity"

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

type GetUserRequest struct {
	UserID string `json:"userID" binding:"required"`
}

type UserProfileResponse struct {
	UserID      string         `json:"user_id"`
	ProfileName string         `json:"profile_name"`
	Mobile      *string        `json:"mobile,omitempty"`
	Gender      *uint8         `json:"gender,omitempty"`
	Birthday    *entity.Date   `json:"birthday,omitempty"`
}

type UserUpdateRequest struct {
	UserID string `json:"user_id" binding:"required"`
	Data   struct {
		ProfileName string       `json:"profile_name"`
		Mobile      *string      `json:"mobile,omitempty"`
		Gender      *uint8       `json:"gender,omitempty"`
		Birthday    *entity.Date `json:"birthday,omitempty"`
	} `json:"data"`
}
