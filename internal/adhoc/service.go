package adhoc

import (
	"context"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/qomarullah/go-rest-api/internal/entity"
	"github.com/qomarullah/go-rest-api/pkg/log"
)

// Service encapsulates usecase logic for adhocs.
type Service interface {
	Get(ctx context.Context, id string) (Adhoc, error)
	Query(ctx context.Context, offset, limit int) ([]Adhoc, error)
	Count(ctx context.Context) (int, error)
	Create(ctx context.Context, input CreateRequest) error
	Update(ctx context.Context, id string, input UpdateRequest) (Adhoc, error)
	Delete(ctx context.Context, id string) (Adhoc, error)
}

// Adhoc represents the data about an adhoc.
type Adhoc struct {
	entity.Adhoc
}

// CreateRequest represents an adhoc creation request.
type CreateRequest struct {
	Name        string `json:"name"`
	Filename    string `json:"filename"`
	Status      string `json:"status"`
	ScheduledAt string `json:"scheduled_at"`
}

// Validate validates the CreateRequest fields.
func (m CreateRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Name, validation.Required, validation.Length(0, 128)),
	)
}

// UpdateRequest represents an adhoc update request.
type UpdateRequest struct {
	Name        string `json:"name"`
	Filename    string `json:"filename"`
	Status      string `json:"status"`
	ScheduledAt string `json:"scheduled_at"`
}

// Validate validates the CreateRequest fields.
func (m UpdateRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Name, validation.Required, validation.Length(0, 128)),
	)
}

type service struct {
	repo   Repository
	logger log.Logger
}

// NewService creates a new adhoc service.
func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

// Get returns the adhoc with the specified the adhoc ID.
func (s service) Get(ctx context.Context, id string) (Adhoc, error) {
	adhoc, err := s.repo.Get(ctx, id)
	if err != nil {
		return Adhoc{}, err
	}
	return Adhoc{adhoc}, nil
}

// Create creates a new adhoc.
func (s service) Create(ctx context.Context, req CreateRequest) error {
	if err := req.Validate(); err != nil {
		return err
	}
	//id := entity.GenerateID()
	//now := time.Now()

	err := s.repo.Create(ctx, entity.Adhoc{
		Name:        req.Name,
		Filename:    req.Filename,
		ScheduledAt: req.ScheduledAt,
		Status:      req.Status,
	})

	return err
}

// Update updates the adhoc with the specified ID.
func (s service) Update(ctx context.Context, id string, req UpdateRequest) (Adhoc, error) {
	if err := req.Validate(); err != nil {
		return Adhoc{}, err
	}

	adhoc, err := s.Get(ctx, id)
	if err != nil {
		return adhoc, err
	}

	adhoc.Name = req.Name
	adhoc.Filename = req.Filename
	adhoc.ScheduledAt = req.ScheduledAt
	adhoc.Status = req.Status

	if err := s.repo.Update(ctx, adhoc.Adhoc); err != nil {
		return adhoc, err
	}
	return adhoc, nil
}

// Delete deletes the adhoc with the specified ID.
func (s service) Delete(ctx context.Context, id string) (Adhoc, error) {
	adhoc, err := s.Get(ctx, id)
	if err != nil {
		return Adhoc{}, err
	}
	if err = s.repo.Delete(ctx, id); err != nil {
		return Adhoc{}, err
	}
	return adhoc, nil
}

// Count returns the number of adhocs.
func (s service) Count(ctx context.Context) (int, error) {
	return s.repo.Count(ctx)
}

// Query returns the adhocs with the specified offset and limit.
func (s service) Query(ctx context.Context, offset, limit int) ([]Adhoc, error) {
	items, err := s.repo.Query(ctx, offset, limit)
	if err != nil {
		return nil, err
	}
	result := []Adhoc{}
	for _, item := range items {
		result = append(result, Adhoc{item})
	}
	return result, nil
}
