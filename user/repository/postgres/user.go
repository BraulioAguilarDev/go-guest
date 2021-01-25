package postgres

import (
	"github.com/brauliodev29/go-guest/models"
	"github.com/brauliodev29/go-guest/pkg/entity"
	"github.com/jmoiron/sqlx"
)

// User struct
type User struct {
	ID           entity.ID `json:"id"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	FirebaseUser string    `json:"firebase_user"`
	Email        string    `json:"email"`
}

// UserRepository struct
type UserRepository struct {
	db *sqlx.DB
}

// NewUserRepository func
func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

// CreateUser func
func (r UserRepository) CreateUser(user *models.User) (entity.ID, error) {
	tx := r.db.MustBegin()

	tx.NamedExec(`INSERT INTO "user" (id, first_name, last_name, email, firebase_user) VALUES (:id, :first_name, :last_name, :email, :firebase_user)`, user)
	if err := tx.Commit(); err != nil {
		return entity.NewID(), err
	}

	return user.ID, nil
}

// UpdateUser func
func (r UserRepository) UpdateUser(user *models.User) error {
	tx := r.db.MustBegin()

	tx.MustExec(`UPDATE "user" SET first_name=$1, last_name=$2 WHERE id=$2`, user.FirstName, user.LastName, user.ID)
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
