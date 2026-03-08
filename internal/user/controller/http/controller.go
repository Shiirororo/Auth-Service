package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/user_service/internal/user/controller/dto"
)

type UserHandler struct {
}

func (uh *UserHandler) GetUserInfoHandler(c *gin.Context) {
	var req = dto.GetUserRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
