package models

import (
	"time"

	"github.com/brauliodev29/go-guest/pkg/entity"
)

// Event struct
type Event struct {
	ID        entity.ID `db:"id"`
	Name      string    `db:"name"`
	Location  string    `db:"location"`
	Date      string    `db:"date"`
	Time      string    `db:"time"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
