package entity

import (
	"time"
)

type Batch struct {
	ID        string    `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	PackageID string    `json:"package_id" gorm:"type:uuid;not null"`
	StartDate time.Time `json:"start_date" gorm:"not null"`
	EndDate   time.Time `json:"end_date" gorm:"not null"`
	Price     float64   `json:"price" gorm:"not null"`
	Quota     int       `json:"quota" gorm:"not null"`
	MinQuota  int       `json:"min_quota" gorm:"not null"`
	IsActive  bool      `json:"is_active" gorm:"default:true"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Package Package `json:"package,omitempty" gorm:"foreignKey:PackageID"`
}
