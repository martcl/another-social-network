package user

import (
	"context"
	"fmt"

	"github.com/martcl/another-social-network/pkg/db"
)

func NewUser(ctx context.Context, name string) (*db.User, error) {
	user := &db.User{
		Username: name,
	}

	client, err := db.NewClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %w", err)
	}

	if err := user.Save(client); err != nil {
		return nil, fmt.Errorf("failed to save user: %w", err)
	}

	return user, nil
}
