// internal/domain/repository/package_repository.go
package repository

import (
	"github.com/lunadiotic/golang-travel-booking/internal/domain/entity"
)

type PackageRepository interface {
	Create(pkg *entity.Package) error
	FindByID(id string) (*entity.Package, error)
	FindByDestination(destinationID string) ([]entity.Package, error)
	FindAll(filters map[string]interface{}) ([]entity.Package, error)
	Update(pkg *entity.Package) error
	Delete(id string) error
}
