package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
}

func (uh *UserHandler) GetUserInfoHandler(c *gin.Context) {
	var req struct {
		AccessToken string `json:"access_token" binding:"required"`
		UserID      string `json:"userID" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
	}
}
