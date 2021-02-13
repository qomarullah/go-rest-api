package user

import (
	"context"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/qiangxue/go-rest-api/internal/entity"
	"github.com/qiangxue/go-rest-api/pkg/log"
)

// Service encapsulates usecase logic for users.
type Service interface {
	Get(ctx context.Context, id string) (User, error)
	Query(ctx context.Context, offset, limit int) ([]User, error)
	Count(ctx context.Context) (int, error)
	Create(ctx context.Context, input CreateUserRequest) (User, error)
	Update(ctx context.Context, id string, input UpdateUserRequest) (User, error)
	Delete(ctx context.Context, id string) (User, error)
}

// User represents the data about an user.
type User struct {
	entity.User
}

// CreateUserRequest represents an user creation request.
type CreateUserRequest struct {
	Name string `json:"name"`
}

// Validate validates the CreateUserRequest fields.
func (m CreateUserRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Name, validation.Required, validation.Length(0, 128)),
	)
}

// UpdateUserRequest represents an user update request.
type UpdateUserRequest struct {
	Name string `json:"name"`
}

// Validate validates the CreateUserRequest fields.
func (m UpdateUserRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Name, validation.Required, validation.Length(0, 128)),
	)
}

type service struct {
	repo   Repository
	logger log.Logger
}

// NewService creates a new user service.
func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

// Get returns the user with the specified the user ID.
func (s service) Get(ctx context.Context, id string) (User, error) {
	user, err := s.repo.Get(ctx, id)
	if err != nil {
		return User{}, err
	}
	return User{user}, nil
}

// Create creates a new user.
func (s service) Create(ctx context.Context, req CreateUserRequest) (User, error) {
	if err := req.Validate(); err != nil {
		return User{}, err
	}
	id := entity.GenerateID()
	now := time.Now()
	err := s.repo.Create(ctx, entity.User{
		ID:        id,
		Name:      req.Name,
		CreatedAt: now,
		UpdatedAt: now,
	})
	if err != nil {
		return User{}, err
	}
	return s.Get(ctx, id)
}

// Update updates the user with the specified ID.
func (s service) Update(ctx context.Context, id string, req UpdateUserRequest) (User, error) {
	if err := req.Validate(); err != nil {
		return User{}, err
	}

	user, err := s.Get(ctx, id)
	if err != nil {
		return user, err
	}
	//user.Name = req.Name
	//user.UpdatedAt = time.Now()

	if err := s.repo.Update(ctx, user.User); err != nil {
		return user, err
	}
	return user, nil
}

// Delete deletes the user with the specified ID.
func (s service) Delete(ctx context.Context, id string) (User, error) {
	user, err := s.Get(ctx, id)
	if err != nil {
		return User{}, err
	}
	if err = s.repo.Delete(ctx, id); err != nil {
		return User{}, err
	}
	return user, nil
}

// Count returns the number of users.
func (s service) Count(ctx context.Context) (int, error) {
	return s.repo.Count(ctx)
}

// Query returns the users with the specified offset and limit.
func (s service) Query(ctx context.Context, offset, limit int) ([]User, error) {
	items, err := s.repo.Query(ctx, offset, limit)
	if err != nil {
		return nil, err
	}
	result := []User{}
	for _, item := range items {
		result = append(result, User{item})
	}
	return result, nil
}
