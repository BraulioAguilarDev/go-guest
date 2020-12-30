package model

import "github.com/brauliodev29/go-guest/pkg/entity"

// User struct
type User struct {
	ID        entity.ID `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Active    bool      `json:"active"`
}
