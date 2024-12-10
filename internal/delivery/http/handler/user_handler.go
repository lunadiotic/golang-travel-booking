package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lunadiotic/golang-travel-booking/internal/domain/entity"
	"github.com/lunadiotic/golang-travel-booking/internal/domain/usecase"
	customErrors "github.com/lunadiotic/golang-travel-booking/internal/usecase"
)

type UserHandler struct {
	userUseCase usecase.UserUseCase
}

type UpdateProfileRequest struct {
	FullName string `json:"full_name" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
}

func NewUserHandler(userUseCase usecase.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: userUseCase,
	}
}

func (h *UserHandler) GetProfile(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	user, err := h.userUseCase.GetProfile(userID.(string))
	if err != nil {
		switch err {
		case customErrors.ErrUserNotFound:
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		case customErrors.ErrInvalidInput:
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get profile"})
		}
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) UpdateProfile(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	var req UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := &entity.User{
		ID:       userID.(string),
		FullName: req.FullName,
		Phone:    req.Phone,
	}

	if err := h.userUseCase.UpdateProfile(user); err != nil {
		switch err {
		case customErrors.ErrUserNotFound:
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		case customErrors.ErrInvalidInput:
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update profile"})
		}
		return
	}

	// Get updated profile
	updatedUser, err := h.userUseCase.GetProfile(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get updated profile"})
		return
	}

	c.JSON(http.StatusOK, updatedUser)
}
