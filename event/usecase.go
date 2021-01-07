package event

import (
	"github.com/brauliodev29/go-guest/models"
	"github.com/brauliodev29/go-guest/pkg/entity"
)

type UseCase interface {
	CreateEvent(e *models.Event) (entity.ID, error)
}
