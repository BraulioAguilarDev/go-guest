package event

import (
	"github.com/brauliodev29/go-guest/models"
	"github.com/brauliodev29/go-guest/pkg/entity"
)

type Repository interface {
	CreateEvent(event *models.Event) (entity.ID, error)
	UpdateEvent(event *models.Event) error
}
