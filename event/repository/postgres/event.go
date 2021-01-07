package postgres

import (
	"github.com/brauliodev29/go-guest/models"
	"github.com/brauliodev29/go-guest/pkg/entity"
	"github.com/jmoiron/sqlx"
)

// Event struct
type Event struct {
	ID       entity.ID `json:"id"`
	Name     string    `json:"name"`
	Location string    `json:"location"`
	Date     string    `json:"date"`
	Time     string    `json:"time"`
}

// EventRepository func
type EventRepository struct {
	db *sqlx.DB
}

// NewEventRepository func
func NewEventRepository(db *sqlx.DB) *EventRepository {
	return &EventRepository{db: db}
}

// CreateEvent func
func (r EventRepository) CreateEvent(event *models.Event) (entity.ID, error) {

	tx := r.db.MustBegin()
	tx.NamedExec("INSERT INTO event (id, name, location, date, time) VALUES (:id, :name, :location, :date, :time)", event)

	if err := tx.Commit(); err != nil {
		return entity.NewID(), err
	}

	return event.ID, nil
}
