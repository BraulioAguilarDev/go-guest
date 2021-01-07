package models

import (
	"time"

	"github.com/brauliodev29/go-guest/pkg/entity"
)

// Event struct
type Event struct {
	ID        entity.ID `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"size:100;not null" json:"name"`
	Location  string    `gorm:"size:100;not null" json:"location"`
	Date      string    `gorm:"size:50;not null" json:"date"`
	Time      time.Time `json:"time"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// TableName func
func (e *Event) TableName() string {
	return "event"
}
