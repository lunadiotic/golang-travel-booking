// internal/repository/destination_repository_impl.go
package repository

import (
	"github.com/lunadiotic/golang-travel-booking/internal/domain/entity"
	"gorm.io/gorm"
)

type destinationRepository struct {
	db *gorm.DB
}

func NewDestinationRepository(db *gorm.DB) *destinationRepository {
	return &destinationRepository{db: db}
}

func (r *destinationRepository) Create(destination *entity.Destination) error {
	return r.db.Create(destination).Error
}

func (r *destinationRepository) FindByID(id string) (*entity.Destination, error) {
	var destination entity.Destination
	if err := r.db.First(&destination, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &destination, nil
}

func (r *destinationRepository) FindAll(filters map[string]interface{}) ([]entity.Destination, error) {
	var destinations []entity.Destination
	query := r.db

	for key, value := range filters {
		query = query.Where(key, value)
	}

	if err := query.Find(&destinations).Error; err != nil {
		return nil, err
	}
	return destinations, nil
}

func (r *destinationRepository) Update(destination *entity.Destination) error {
	return r.db.Save(destination).Error
}

func (r *destinationRepository) Delete(id string) error {
	return r.db.Delete(&entity.Destination{}, "id = ?", id).Error
}
