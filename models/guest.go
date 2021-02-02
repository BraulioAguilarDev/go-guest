package models

import "github.com/brauliodev29/go-guest/pkg/entity"

// Guest struct
type Guest struct {
	ID       entity.ID `json:"id"`
	EventID  entity.ID `json:"event_id"`
	DetailID entity.ID `json:"detail_id"`
	Email    string    `json:"email"`
	FullName string    `json:"full_name"`
	Status   string    `json:"status"`
}
