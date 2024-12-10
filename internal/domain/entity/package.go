package entity

import (
	"time"
)

type Package struct {
	ID            string    `json:"id" gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	DestinationID string    `json:"destination_id" gorm:"type:uuid;not null"`
	Title         string    `json:"title" gorm:"not null"`
	Description   string    `json:"description" gorm:"type:text"`
	Duration      int       `json:"duration" gorm:"not null"` // dalam hari
	ImageURL      string    `json:"image_url"`
	IsActive      bool      `json:"is_active" gorm:"default:true"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`

	Destination Destination `json:"destination,omitempty" gorm:"foreignKey:DestinationID"`
	Batches     []Batch     `json:"batches,omitempty" gorm:"foreignKey:PackageID"`

	// Informasi itinerary disimpan sebagai JSON
	Itinerary []DaySchedule `json:"itinerary" gorm:"type:jsonb"`
}

type DaySchedule struct {
	Day        int      `json:"day"`
	Activities []string `json:"activities"`
	MealPlan   string   `json:"meal_plan"` // e.g., "Breakfast, Lunch, Dinner"
}
