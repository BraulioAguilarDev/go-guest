package user

import "github.com/brauliodev29/go-guest/pkg/entity"

// UseCase interface
type UseCase interface {
	CreateUser(firstName, lastName, email, phone, password string) (entity.ID, error)
	UpdateUser(firstName, lastName, phone string, id entity.ID) error
}
