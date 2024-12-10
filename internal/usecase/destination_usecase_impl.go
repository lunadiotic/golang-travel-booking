// internal/usecase/destination_usecase_impl.go
package usecase

import (
	"github.com/lunadiotic/golang-travel-booking/internal/domain/entity"
	"github.com/lunadiotic/golang-travel-booking/internal/domain/repository"
)

type destinationUseCase struct {
	destinationRepo repository.DestinationRepository
}

func NewDestinationUseCase(destinationRepo repository.DestinationRepository) *destinationUseCase {
	return &destinationUseCase{
		destinationRepo: destinationRepo,
	}
}

func (u *destinationUseCase) Create(destination *entity.Destination) error {
	if destination.Name == "" || destination.City == "" || destination.Province == "" {
		return ErrInvalidInput
	}
	return u.destinationRepo.Create(destination)
}

func (u *destinationUseCase) GetByID(id string) (*entity.Destination, error) {
	if id == "" {
		return nil, ErrInvalidInput
	}

	destination, err := u.destinationRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if destination == nil {
		return nil, ErrDestinationNotFound
	}

	return destination, nil
}

func (u *destinationUseCase) GetAll() ([]entity.Destination, error) {
	return u.destinationRepo.FindAll(map[string]interface{}{
		"is_active": true,
	})
}

func (u *destinationUseCase) Update(destination *entity.Destination) error {
	if destination.ID == "" {
		return ErrInvalidInput
	}

	existing, err := u.destinationRepo.FindByID(destination.ID)
	if err != nil {
		return err
	}
	if existing == nil {
		return ErrDestinationNotFound
	}

	return u.destinationRepo.Update(destination)
}

func (u *destinationUseCase) Delete(id string) error {
	if id == "" {
		return ErrInvalidInput
	}

	existing, err := u.destinationRepo.FindByID(id)
	if err != nil {
		return err
	}
	if existing == nil {
		return ErrDestinationNotFound
	}

	return u.destinationRepo.Delete(id)
}
