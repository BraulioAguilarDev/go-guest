package models

import (
	"time"

	"github.com/brauliodev29/go-guest/pkg/entity"
)

// User struct
type User struct {
	ID           entity.ID `db:"id" json:"id"`
	FirstName    string    `db:"first_name" json:"first_name"`
	LastName     string    `db:"last_name" json:"last_name"`
	FirebaseUser string    `db:"firebase_user" json:"firebase_user"`
	Email        string    `db:"email" json:"email"`
	Phone        string    `db:"phone" json:"phone"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time `db:"updated_at" json:"updated_at"`
}
