package user

import (
	"github.com/brauliodev29/go-guest/models"
	"github.com/brauliodev29/go-guest/pkg/entity"
)

// Repository interface
type Repository interface {
	CreateUser(user *models.User) (entity.ID, error)
	UpdateUser(user *models.User) error
}
