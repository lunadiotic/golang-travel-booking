// internal/repository/package_repository_impl.go
package repository

import (
	"github.com/lunadiotic/golang-travel-booking/internal/domain/entity"
	"gorm.io/gorm"
)

type packageRepository struct {
	db *gorm.DB
}

func NewPackageRepository(db *gorm.DB) *packageRepository {
	return &packageRepository{db: db}
}

func (r *packageRepository) Create(pkg *entity.Package) error {
	return r.db.Create(pkg).Error
}

func (r *packageRepository) FindByID(id string) (*entity.Package, error) {
	var pkg entity.Package
	if err := r.db.Preload("Destination").First(&pkg, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &pkg, nil
}

func (r *packageRepository) FindByDestination(destinationID string) ([]entity.Package, error) {
	var packages []entity.Package
	if err := r.db.Where("destination_id = ?", destinationID).Find(&packages).Error; err != nil {
		return nil, err
	}
	return packages, nil
}

func (r *packageRepository) FindAll(filters map[string]interface{}) ([]entity.Package, error) {
	var packages []entity.Package
	query := r.db.Preload("Destination")

	for key, value := range filters {
		query = query.Where(key, value)
	}

	if err := query.Find(&packages).Error; err != nil {
		return nil, err
	}
	return packages, nil
}

func (r *packageRepository) Update(pkg *entity.Package) error {
	return r.db.Save(pkg).Error
}

func (r *packageRepository) Delete(id string) error {
	return r.db.Delete(&entity.Package{}, "id = ?", id).Error
}
