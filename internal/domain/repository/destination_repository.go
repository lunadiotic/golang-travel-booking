// internal/domain/repository/destination_repository.go
package repository

import (
	"github.com/lunadiotic/golang-travel-booking/internal/domain/entity"
)

type DestinationRepository interface {
	Create(destination *entity.Destination) error
	FindByID(id string) (*entity.Destination, error)
	FindAll(filters map[string]interface{}) ([]entity.Destination, error)
	Update(destination *entity.Destination) error
	Delete(id string) error
}
