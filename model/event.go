package model

import (
	"time"

	"github.com/brauliodev29/go-guest/pkg/entity"
)

// Event struct
type Event struct {
	ID       entity.ID `json:"id"`
	Name     string    `json:"name"`
	Location string    `json:"location"`
	Date     string    `json:"date"`
	Time     time.Time `json:"time"`
}
