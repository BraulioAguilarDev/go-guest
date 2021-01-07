package usecase

import (
	"github.com/brauliodev29/go-guest/event"
	"github.com/brauliodev29/go-guest/models"
	"github.com/brauliodev29/go-guest/pkg/entity"
)

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
func (e EventUseCase) CreateEvent(event *models.Event) (entity.ID, error) {
	return e.eventRepo.CreateEvent(event)
}
