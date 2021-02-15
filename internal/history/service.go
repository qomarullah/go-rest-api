package history

import (
	"context"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/qomarullah/go-rest-api/internal/entity"
	"github.com/qomarullah/go-rest-api/pkg/log"
)

// Service encapsulates usecase logic for historys.
type Service interface {
	Get(ctx context.Context, id string) (History, error)
	GetByTrxId(ctx context.Context, id string) (History, error)
	Query(ctx context.Context, offset, limit int) ([]History, error)
	Count(ctx context.Context) (int, error)
	Create(ctx context.Context, input CreateRequest) (History, error)
	Update(ctx context.Context, id string, input UpdateRequest) (History, error)
	Delete(ctx context.Context, id string) (History, error)
}

// History represents the data about an history.
type History struct {
	entity.History
}

// CreateRequest represents an history creation request.
type CreateRequest struct {
	Msisdn string `json:"msisdn"`
}

// Validate validates the CreateRequest fields.
func (m CreateRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Msisdn, validation.Required, validation.Length(0, 128)),
	)
}

// UpdateRequest represents an history update request.
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

// NewService creates a new history service.
func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

// Get returns the history with the specified the history ID.
func (s service) Get(ctx context.Context, id string) (History, error) {
	history, err := s.repo.Get(ctx, id)
	if err != nil {
		return History{}, err
	}
	return History{history}, nil
}

// Get returns the history with the specified the history ID.
func (s service) GetByTrxId(ctx context.Context, id string) (History, error) {
	history, err := s.repo.GetByTrxId(ctx, id)
	if err != nil {
		return History{}, err
	}
	return History{history}, nil
}

// Create creates a new history.
func (s service) Create(ctx context.Context, req CreateRequest) (History, error) {
	if err := req.Validate(); err != nil {
		return History{}, err
	}
	id := entity.GenerateID()
	//now := time.Now()
	err := s.repo.Create(ctx, entity.History{
		Msisdn: req.Msisdn,
	})
	if err != nil {
		return History{}, err
	}
	return s.Get(ctx, id)
}

// Update updates the history with the specified ID.
func (s service) Update(ctx context.Context, id string, req UpdateRequest) (History, error) {
	if err := req.Validate(); err != nil {
		return History{}, err
	}

	history, err := s.Get(ctx, id)
	if err != nil {
		return history, err
	}
	//history.Name = req.Name
	//history.UpdatedAt = time.Now()

	if err := s.repo.Update(ctx, history.History); err != nil {
		return history, err
	}
	return history, nil
}

// Delete deletes the history with the specified ID.
func (s service) Delete(ctx context.Context, id string) (History, error) {
	history, err := s.Get(ctx, id)
	if err != nil {
		return History{}, err
	}
	if err = s.repo.Delete(ctx, id); err != nil {
		return History{}, err
	}
	return history, nil
}

// Count returns the number of historys.
func (s service) Count(ctx context.Context) (int, error) {
	return s.repo.Count(ctx)
}

// Query returns the historys with the specified offset and limit.
func (s service) Query(ctx context.Context, offset, limit int) ([]History, error) {
	items, err := s.repo.Query(ctx, offset, limit)
	if err != nil {
		return nil, err
	}
	result := []History{}
	for _, item := range items {
		result = append(result, History{item})
	}
	return result, nil
}
