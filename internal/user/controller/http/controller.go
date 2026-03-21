package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/user_service/internal/user/application/service"
	"github.com/user_service/internal/user/controller/dto"
)

type UserHandler struct {
	userService service.UserServiceInterface
}

func NewUserHandler(userService service.UserServiceInterface) *UserHandler {
	return &UserHandler{userService: userService}
}

func (uh *UserHandler) GetUserInfoHandler(c *gin.Context) {
	var req = dto.GetUserRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	profile, err := uh.userService.GetUserInfo(c.Request.Context(), req.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, profile)
}

func (uh *UserHandler) RegisterHandler(c *gin.Context) {
	var req = dto.RegisterRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}
	err := uh.userService.RegisterService(c.Request.Context(), req.Username, req.Password, req.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Registration successful",
	})
}
func (uh *UserHandler) UpdateUserInfoHandler(c *gin.Context) {
	var req = dto.UserUpdateRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
	}
	err := uh.userService.UpdateUserInfo(c.Request.Context(), req.UserID, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Update successful",
	})
}
