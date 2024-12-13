// internal/delivery/http/handler/dto/destination_dto.go
package dto

type CreateDestinationRequest struct {
	Name        string `json:"name" binding:"required"`
	City        string `json:"city" binding:"required"`
	Province    string `json:"province" binding:"required"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
}

type UpdateDestinationRequest struct {
	Name        string `json:"name" binding:"required"`
	City        string `json:"city" binding:"required"`
	Province    string `json:"province" binding:"required"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
	IsActive    bool   `json:"is_active"`
}

type DestinationResponse struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	City        string `json:"city"`
	Province    string `json:"province"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
	IsActive    bool   `json:"is_active"`
}
