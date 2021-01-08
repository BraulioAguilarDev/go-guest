package event

import (
	"github.com/brauliodev29/go-guest/pkg/entity"
)

type UseCase interface {
	CreateEvent(name, location, date, time string) (entity.ID, error)
	UpdateEvent(name, location, date, time string, id entity.ID) error
}
