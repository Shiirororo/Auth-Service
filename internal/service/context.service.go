package service

import (
	"github.com/gin-gonic/gin"
)

const (
	ContextUserID = "userID"
	ContextRole   = "role"
)

// GetUserID retrieves the UserID from the Gin context
func GetUserID(c *gin.Context) (string, bool) {
	val, exists := c.Get(ContextUserID)
	if !exists {
		return "", false
	}
	userID, ok := val.(string)
	return userID, ok
}

// SetUserID sets the UserID in the Gin context
func SetUserID(c *gin.Context, userID string) {
	c.Set(ContextUserID, userID)
}
