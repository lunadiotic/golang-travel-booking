package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/lunadiotic/golang-travel-booking/internal/domain/usecase"
)

type UserHandler struct {
	userUseCase usecase.UserUseCase // akan di-inject
}

func NewUserHandler(userUseCase usecase.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: userUseCase,
	}
}

func (h *UserHandler) GetProfile(c *gin.Context) {
	// Implementasi get profile
	c.JSON(200, gin.H{
		"message": "Get profile endpoint",
	})
}

func (h *UserHandler) UpdateProfile(c *gin.Context) {
	// Implementasi update profile
	c.JSON(200, gin.H{
		"message": "Update profile endpoint",
	})
}
