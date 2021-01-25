package usecase

import (
	"time"

	"github.com/brauliodev29/go-guest/models"
	"github.com/brauliodev29/go-guest/pkg/entity"
	"github.com/brauliodev29/go-guest/user"
)

// AuthUseCase struct
type AuthUseCase struct {
	userRepo user.Repository
}

// NewAuthUseCase func
func NewAuthUseCase(userCase user.Repository) *AuthUseCase {
	return &AuthUseCase{
		userRepo: userCase,
	}
}

// CreateUser func
func (a AuthUseCase) CreateUser(firstName, lastName, email string) (entity.ID, error) {

	// Add FirebaseSDK

	user := &models.User{
		ID:           entity.NewID(),
		FirstName:    firstName,
		LastName:     lastName,
		FirebaseUser: "dedefrfrgtghtgt",
		Email:        email,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	return a.userRepo.CreateUser(user)
}

// UpdateUser func
func (a AuthUseCase) UpdateUser(firstName, lastName string, id entity.ID) error {

	user := &models.User{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
		UpdatedAt: time.Now(),
	}
	return a.userRepo.UpdateUser(user)
}
