package usecase

import (
	"context"
	"time"

	"firebase.google.com/go/auth"
	"github.com/brauliodev29/go-guest/models"
	"github.com/brauliodev29/go-guest/pkg/entity"
	"github.com/brauliodev29/go-guest/user"
)

// AuthUseCase struct
type AuthUseCase struct {
	userRepo   user.Repository
	authClient *auth.Client
}

// NewAuthUseCase func
func NewAuthUseCase(
	userCase user.Repository,
	client *auth.Client,
) *AuthUseCase {
	return &AuthUseCase{
		userRepo:   userCase,
		authClient: client,
	}
}

// CreateUser func
func (a AuthUseCase) CreateUser(firstName, lastName, email, phone, password string) (entity.ID, error) {

	user := &models.User{
		ID:        entity.NewID(),
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Phone:     phone,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Add FirebaseSDK
	ctx := context.Background()
	params := (&auth.UserToCreate{}).
		DisplayName(user.FirstName).
		Email(user.Email).
		Password(password).
		Disabled(false)

	remoteUser, err := a.authClient.CreateUser(ctx, params)
	if err != nil {
		return user.ID, err
	}

	user.FirebaseUser = remoteUser.UID

	return a.userRepo.CreateUser(user)
}

// UpdateUser func
func (a AuthUseCase) UpdateUser(firstName, lastName, phone string, id entity.ID) error {

	user := &models.User{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
		Phone:     phone,
		UpdatedAt: time.Now(),
	}
	return a.userRepo.UpdateUser(user)
}
