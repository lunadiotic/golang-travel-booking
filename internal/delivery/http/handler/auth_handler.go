// internal/delivery/http/handler/auth_handler.go
package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lunadiotic/golang-travel-booking/internal/delivery/http/handler/dto"
	"github.com/lunadiotic/golang-travel-booking/internal/domain/entity"
	"github.com/lunadiotic/golang-travel-booking/internal/domain/usecase"
)

type AuthHandler struct {
	userUseCase usecase.UserUseCase // akan di-inject
}

func NewAuthHandler(userUseCase usecase.UserUseCase) *AuthHandler {
	return &AuthHandler{
		userUseCase: userUseCase,
	}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req dto.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user := &entity.User{
		Email:    req.Email,
		Password: req.Password, // Nanti akan di-hash
		FullName: req.FullName,
		Phone:    req.Phone,
		Role:     "customer", // default role
	}

	if err := h.userUseCase.Register(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"data": gin.H{
			"id":        user.ID,
			"email":     user.Email,
			"full_name": user.FullName,
			"phone":     user.Phone,
		},
	})
}

func (h *AuthHandler) Login(c *gin.Context) {
	// Implementasi login
	c.JSON(200, gin.H{
		"message": "Login endpoint",
	})
}
