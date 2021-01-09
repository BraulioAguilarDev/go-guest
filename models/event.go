package models

import (
	"time"

	"github.com/brauliodev29/go-guest/pkg/entity"
)

// Event struct
type Event struct {
	ID        entity.ID `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	Location  string    `db:"location" json:"location"`
	Date      string    `db:"date" json:"date"`
	Time      string    `db:"time" json:"time"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
