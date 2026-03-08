package dto

type GetUserRequest struct {
	AccessToken string `json:"access_token" binding:"required"`
	UserID      string `json:"userID" binding:"required"`
}
