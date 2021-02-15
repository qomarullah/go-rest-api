package reminder

import (
	"context"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/qomarullah/go-rest-api/internal/entity"
	"github.com/qomarullah/go-rest-api/pkg/log"
)

// Service encapsulates usecase logic for reminders.
type Service interface {
	Get(ctx context.Context, id string) (Reminder, error)
	Query(ctx context.Context, offset, limit int) ([]Reminder, error)
	Count(ctx context.Context) (int, error)
	Create(ctx context.Context, input CreateRequest) (Reminder, error)
	Update(ctx context.Context, id string, input UpdateRequest) (Reminder, error)
	Delete(ctx context.Context, id string) (Reminder, error)
}

// Reminder represents the data about an reminder.
type Reminder struct {
	entity.Reminder
}

// CreateRequest represents an reminder creation request.
type CreateRequest struct {
	Msisdn string `json:"msisdn"`
}

// Validate validates the CreateRequest fields.
func (m CreateRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Msisdn, validation.Required, validation.Length(0, 128)),
	)
}

// UpdateRequest represents an reminder update request.
type UpdateRequest struct {
	Msisdn string `json:"msisdn"`
}

// Validate validates the CreateRequest fields.
func (m UpdateRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Msisdn, validation.Required, validation.Length(0, 128)),
	)
}

type service struct {
	repo   Repository
	logger log.Logger
}

// NewService creates a new reminder service.
func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

// Get returns the reminder with the specified the reminder ID.
func (s service) Get(ctx context.Context, id string) (Reminder, error) {
	reminder, err := s.repo.Get(ctx, id)
	if err != nil {
		return Reminder{}, err
	}
	return Reminder{reminder}, nil
}

// Create creates a new reminder.
func (s service) Create(ctx context.Context, req CreateRequest) (Reminder, error) {
	if err := req.Validate(); err != nil {
		return Reminder{}, err
	}
	id := entity.GenerateID()
	//now := time.Now()
	err := s.repo.Create(ctx, entity.Reminder{
		Msisdn: req.Msisdn,
	})
	if err != nil {
		return Reminder{}, err
	}
	return s.Get(ctx, id)
}

// Update updates the reminder with the specified ID.
func (s service) Update(ctx context.Context, id string, req UpdateRequest) (Reminder, error) {
	if err := req.Validate(); err != nil {
		return Reminder{}, err
	}

	reminder, err := s.Get(ctx, id)
	if err != nil {
		return reminder, err
	}
	//reminder.Name = req.Name
	//reminder.UpdatedAt = time.Now()

	if err := s.repo.Update(ctx, reminder.Reminder); err != nil {
		return reminder, err
	}
	return reminder, nil
}

// Delete deletes the reminder with the specified ID.
func (s service) Delete(ctx context.Context, id string) (Reminder, error) {
	reminder, err := s.Get(ctx, id)
	if err != nil {
		return Reminder{}, err
	}
	if err = s.repo.Delete(ctx, id); err != nil {
		return Reminder{}, err
	}
	return reminder, nil
}

// Count returns the number of reminders.
func (s service) Count(ctx context.Context) (int, error) {
	return s.repo.Count(ctx)
}

// Query returns the reminders with the specified offset and limit.
func (s service) Query(ctx context.Context, offset, limit int) ([]Reminder, error) {
	items, err := s.repo.Query(ctx, offset, limit)
	if err != nil {
		return nil, err
	}
	result := []Reminder{}
	for _, item := range items {
		result = append(result, Reminder{item})
	}
	return result, nil
}
