// internal/domain/repository/batch_repository.go
package repository

import (
	"time"

	"github.com/lunadiotic/golang-travel-booking/internal/domain/entity"
)

type BatchRepository interface {
	Create(batch *entity.Batch) error
	FindByID(id string) (*entity.Batch, error)
	FindByPackage(packageID string) ([]entity.Batch, error)
	FindByDateRange(startDate, endDate time.Time) ([]entity.Batch, error)
	FindAll(filters map[string]interface{}) ([]entity.Batch, error)
	Update(batch *entity.Batch) error
	UpdateQuota(id string, quota int) error
	Delete(id string) error
}
