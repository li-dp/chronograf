package mocks

import (
	"context"

	"github.com/influxdata/chronograf"
)

var _ chronograf.SourceUsersStore = &DBUsersStore{}

// DBUsersStore mock allows all functions to be set for testing
type DBUsersStore struct {
	AllF    func(context.Context) ([]chronograf.SourceUser, error)
	AddF    func(context.Context, *chronograf.SourceUser) (*chronograf.SourceUser, error)
	DeleteF func(context.Context, *chronograf.SourceUser) error
	GetF    func(ctx context.Context, name string) (*chronograf.SourceUser, error)
	UpdateF func(context.Context, *chronograf.SourceUser) error
}

// All lists all users from the DBUsersStore
func (s *DBUsersStore) All(ctx context.Context) ([]chronograf.SourceUser, error) {
	return s.AllF(ctx)
}

// Add a new User in the DBUsersStore
func (s *DBUsersStore) Add(ctx context.Context, u *chronograf.SourceUser) (*chronograf.SourceUser, error) {
	return s.AddF(ctx, u)
}

// Delete the User from the DBUsersStore
func (s *DBUsersStore) Delete(ctx context.Context, u *chronograf.SourceUser) error {
	return s.DeleteF(ctx, u)
}

// Get retrieves a user if name exists.
func (s *DBUsersStore) Get(ctx context.Context, name string) (*chronograf.SourceUser, error) {
	return s.GetF(ctx, name)
}

// Update the user's permissions or roles
func (s *DBUsersStore) Update(ctx context.Context, u *chronograf.SourceUser) error {
	return s.UpdateF(ctx, u)
}