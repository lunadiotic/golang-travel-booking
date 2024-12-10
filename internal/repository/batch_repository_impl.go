// internal/repository/batch_repository_impl.go
package repository

import (
	"time"

	"github.com/lunadiotic/golang-travel-booking/internal/domain/entity"
	"gorm.io/gorm"
)

type batchRepository struct {
	db *gorm.DB
}

func NewBatchRepository(db *gorm.DB) *batchRepository {
	return &batchRepository{db: db}
}

func (r *batchRepository) Create(batch *entity.Batch) error {
	return r.db.Create(batch).Error
}

func (r *batchRepository) FindByID(id string) (*entity.Batch, error) {
	var batch entity.Batch
	if err := r.db.Preload("Package.Destination").First(&batch, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &batch, nil
}

func (r *batchRepository) FindByPackage(packageID string) ([]entity.Batch, error) {
	var batches []entity.Batch
	if err := r.db.Preload("Package.Destination").Where("package_id = ?", packageID).Find(&batches).Error; err != nil {
		return nil, err
	}
	return batches, nil
}

func (r *batchRepository) FindByDateRange(startDate, endDate time.Time) ([]entity.Batch, error) {
	var batches []entity.Batch
	if err := r.db.Preload("Package.Destination").
		Where("start_date >= ? AND end_date <= ?", startDate, endDate).
		Find(&batches).Error; err != nil {
		return nil, err
	}
	return batches, nil
}

func (r *batchRepository) FindAll(filters map[string]interface{}) ([]entity.Batch, error) {
	var batches []entity.Batch
	query := r.db.Preload("Package.Destination")

	for key, value := range filters {
		query = query.Where(key, value)
	}

	if err := query.Find(&batches).Error; err != nil {
		return nil, err
	}
	return batches, nil
}

func (r *batchRepository) Update(batch *entity.Batch) error {
	return r.db.Save(batch).Error
}

func (r *batchRepository) UpdateQuota(id string, quota int) error {
	return r.db.Model(&entity.Batch{}).
		Where("id = ?", id).
		Update("quota", quota).
		Error
}

func (r *batchRepository) Delete(id string) error {
	return r.db.Delete(&entity.Batch{}, "id = ?", id).Error
}

// Transaction support
func (r *batchRepository) WithTx(tx *gorm.DB) *batchRepository {
	if tx == nil {
		return r
	}
	return &batchRepository{
		db: tx,
	}
}
