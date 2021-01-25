package firebase

import (
	"context"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

// Firebase struct
type Firebase struct {
	credentialFile string
}

// NewFirebase func
func NewFirebase(file string) *Firebase {
	return &Firebase{
		credentialFile: file,
	}
}

// NewClient func
func (f *Firebase) NewClient() (*auth.Client, error) {
	ctx := context.Background()
	options := option.WithCredentialsFile(f.credentialFile)

	app, err := firebase.NewApp(ctx, nil, options)
	if err != nil {
		return nil, err
	}

	return app.Auth(ctx)
}
