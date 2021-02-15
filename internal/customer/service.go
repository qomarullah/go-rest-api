package customer

import (
	"context"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/qomarullah/go-rest-api/internal/entity"
	"github.com/qomarullah/go-rest-api/pkg/log"
)

// Service encapsulates usecase logic for customers.
type Service interface {
	Get(ctx context.Context, id string) (Customer, error)
	Query(ctx context.Context, offset, limit int) ([]Customer, error)
	Count(ctx context.Context) (int, error)
	Create(ctx context.Context, input CreateRequest) (Customer, error)
	Update(ctx context.Context, id string, input UpdateRequest) (Customer, error)
	Delete(ctx context.Context, id string) (Customer, error)
}

// Customer represents the data about an customer.
type Customer struct {
	entity.Customer
}

// CreateRequest represents an customer creation request.
type CreateRequest struct {
	Msisdn string `json:"msisdn"`
}

// Validate validates the CreateRequest fields.
func (m CreateRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Msisdn, validation.Required, validation.Length(0, 128)),
	)
}

// UpdateRequest represents an customer update request.
type UpdateRequest struct {
	Name string `json:"name"`
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

// NewService creates a new customer service.
func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

// Get returns the customer with the specified the customer ID.
func (s service) Get(ctx context.Context, id string) (Customer, error) {
	customer, err := s.repo.Get(ctx, id)
	if err != nil {
		return Customer{}, err
	}
	return Customer{customer}, nil
}

// Create creates a new customer.
func (s service) Create(ctx context.Context, req CreateRequest) (Customer, error) {
	if err := req.Validate(); err != nil {
		return Customer{}, err
	}
	id := entity.GenerateID()
	//now := time.Now()
	err := s.repo.Create(ctx, entity.Customer{
		Msisdn: req.Msisdn,
	})
	if err != nil {
		return Customer{}, err
	}
	return s.Get(ctx, id)
}

// Update updates the customer with the specified ID.
func (s service) Update(ctx context.Context, id string, req UpdateRequest) (Customer, error) {
	if err := req.Validate(); err != nil {
		return Customer{}, err
	}

	customer, err := s.Get(ctx, id)
	if err != nil {
		return customer, err
	}
	//customer.Name = req.Name
	//customer.UpdatedAt = time.Now()

	if err := s.repo.Update(ctx, customer.Customer); err != nil {
		return customer, err
	}
	return customer, nil
}

// Delete deletes the customer with the specified ID.
func (s service) Delete(ctx context.Context, id string) (Customer, error) {
	customer, err := s.Get(ctx, id)
	if err != nil {
		return Customer{}, err
	}
	if err = s.repo.Delete(ctx, id); err != nil {
		return Customer{}, err
	}
	return customer, nil
}

// Count returns the number of customers.
func (s service) Count(ctx context.Context) (int, error) {
	return s.repo.Count(ctx)
}

// Query returns the customers with the specified offset and limit.
func (s service) Query(ctx context.Context, offset, limit int) ([]Customer, error) {
	items, err := s.repo.Query(ctx, offset, limit)
	if err != nil {
		return nil, err
	}
	result := []Customer{}
	for _, item := range items {
		result = append(result, Customer{item})
	}
	return result, nil
}
