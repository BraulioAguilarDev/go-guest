package event

import (
	"github.com/brauliodev29/go-guest/models"
	"github.com/brauliodev29/go-guest/pkg/entity"
)

// UseCase interface
type UseCase interface {
	CreateEvent(name, location, date, time string) (entity.ID, error)
	UpdateEvent(name, location, date, time string, id entity.ID) error
	FindAllEvent() ([]*models.Event, error)
	FindOneEvent(id entity.ID) (*models.Event, error)
}
