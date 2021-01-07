package usecase

import (
	"github.com/brauliodev29/go-guest/event"
	"github.com/brauliodev29/go-guest/models"
	"github.com/brauliodev29/go-guest/pkg/entity"
)

// EventUseCase struct
type EventUseCase struct {
	eventRepo event.Repository
}

// NewEventUseCase func
func NewEventUseCase(eventRepo event.Repository) *EventUseCase {
	return &EventUseCase{
		eventRepo: eventRepo,
	}
}

// CreateEvent func
func (e EventUseCase) CreateEvent(name, location, date, time string) (entity.ID, error) {
	event := &models.Event{
		ID:       entity.NewID(),
		Name:     name,
		Location: location,
		Date:     date,
		Time:     time,
	}

	return e.eventRepo.CreateEvent(event)
}
