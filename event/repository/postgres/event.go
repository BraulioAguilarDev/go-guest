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

// UpdateEvent func
func (r EventRepository) UpdateEvent(event *models.Event) error {
	tx := r.db.MustBegin()
	tx.MustExec("UPDATE event SET name=$1, location=$2, date=$3, time=$4 WHERE id=$5", event.Name, event.Location, event.Date, event.Time, event.ID)
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

// FindAllEvent func
func (r EventRepository) FindAllEvent() ([]*models.Event, error) {

	events := []*models.Event{}
	if err := r.db.Select(&events, "SELECT * FROM event ORDER BY date ASC"); err != nil {
		return nil, err
	}

	return events, nil
}

// FindOneEvent func
func (r EventRepository) FindOneEvent(id entity.ID) (*models.Event, error) {
	event := models.Event{}

	if err := r.db.Get(&event, "SELECT * FROM event WHERE id=$1 LIMIT 1", id); err != nil {
		return nil, err
	}

	return &event, nil
}
