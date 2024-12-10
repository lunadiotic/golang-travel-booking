package entity

import (
	"time"
)

type Destination struct {
	ID          string    `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name        string    `json:"name" gorm:"not null"`
	City        string    `json:"city" gorm:"not null"`
	Province    string    `json:"province" gorm:"not null"`
	Description string    `json:"description" gorm:"type:text"`
	ImageURL    string    `json:"image_url"`
	IsActive    bool      `json:"is_active" gorm:"default:true"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Packages    []Package `json:"packages,omitempty" gorm:"foreignKey:DestinationID"`
}
