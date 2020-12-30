package model

import "github.com/brauliodev29/go-guest/pkg/entity"

// Guest struct
type Guest struct {
	ID      entity.ID `json:"id"`
	EventID string    `json:"event_id"`
	UserID  string    `json:"user_id"`
	Confirm bool      `json:"confirm"`
}
