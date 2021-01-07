package postgres

import (
	"time"

	"github.com/brauliodev29/go-guest/models"
	"github.com/brauliodev29/go-guest/pkg/entity"
	"github.com/jinzhu/gorm"
)

// Event struct
type Event struct {
	ID       entity.ID `json:"id"`
	Name     string    `json:"name"`
	Location string    `json:"location"`
	Date     string    `json:"date"`
	Time     time.Time `json:"time"`
}

// EventRepository func
type EventRepository struct {
	db *gorm.DB
}

// NewEventRepository func
func NewEventRepository(db *gorm.DB) *EventRepository {
	return &EventRepository{db: db}
}

// CreateEvent func
func (r EventRepository) CreateEvent(e *models.Event) (entity.ID, error) {
	if err := r.db.Create(e).Error; err != nil {
		return e.ID, nil
	}

	return e.ID, nil
}
