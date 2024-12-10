// internal/delivery/http/handler/destination_handler.go
package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lunadiotic/golang-travel-booking/internal/delivery/http/handler/dto"
	"github.com/lunadiotic/golang-travel-booking/internal/domain/entity"
	"github.com/lunadiotic/golang-travel-booking/internal/domain/usecase"
)

type DestinationHandler struct {
	destinationUseCase usecase.DestinationUseCase
}

func NewDestinationHandler(destinationUseCase usecase.DestinationUseCase) *DestinationHandler {
	return &DestinationHandler{
		destinationUseCase: destinationUseCase,
	}
}

func (h *DestinationHandler) Create(c *gin.Context) {
	var req dto.CreateDestinationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	destination := &entity.Destination{
		Name:        req.Name,
		City:        req.City,
		Province:    req.Province,
		Description: req.Description,
		ImageURL:    req.ImageURL,
		IsActive:    true,
	}

	if err := h.destinationUseCase.Create(destination); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := dto.DestinationResponse{
		ID:          destination.ID,
		Name:        destination.Name,
		City:        destination.City,
		Province:    destination.Province,
		Description: destination.Description,
		ImageURL:    destination.ImageURL,
		IsActive:    destination.IsActive,
	}

	c.JSON(http.StatusCreated, response)
}

func (h *DestinationHandler) GetByID(c *gin.Context) {
	id := c.Param("id")

	destination, err := h.destinationUseCase.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	response := dto.DestinationResponse{
		ID:          destination.ID,
		Name:        destination.Name,
		City:        destination.City,
		Province:    destination.Province,
		Description: destination.Description,
		ImageURL:    destination.ImageURL,
		IsActive:    destination.IsActive,
	}

	c.JSON(http.StatusOK, response)
}

func (h *DestinationHandler) GetAll(c *gin.Context) {
	destinations, err := h.destinationUseCase.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var response []dto.DestinationResponse
	for _, dest := range destinations {
		response = append(response, dto.DestinationResponse{
			ID:          dest.ID,
			Name:        dest.Name,
			City:        dest.City,
			Province:    dest.Province,
			Description: dest.Description,
			ImageURL:    dest.ImageURL,
			IsActive:    dest.IsActive,
		})
	}

	c.JSON(http.StatusOK, response)
}

func (h *DestinationHandler) Update(c *gin.Context) {
	id := c.Param("id")

	var req dto.UpdateDestinationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	destination := &entity.Destination{
		ID:          id,
		Name:        req.Name,
		City:        req.City,
		Province:    req.Province,
		Description: req.Description,
		ImageURL:    req.ImageURL,
	}

	if err := h.destinationUseCase.Update(destination); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Destination updated successfully"})
}

func (h *DestinationHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	if err := h.destinationUseCase.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Destination deleted successfully"})
}
