// internal/domain/usecase/destination_usecase.go
package usecase

import "github.com/lunadiotic/golang-travel-booking/internal/domain/entity"

type DestinationUseCase interface {
	Create(destination *entity.Destination) error
	GetByID(id string) (*entity.Destination, error)
	GetAll() ([]entity.Destination, error)
	Update(destination *entity.Destination) error
	Delete(id string) error
}
