package menu

import (
	"context"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/qomarullah/go-rest-api/internal/entity"
	"github.com/qomarullah/go-rest-api/pkg/log"
)

// Service encapsulates usecase logic for menus.
type Service interface {
	Get(ctx context.Context, id string) (Menu, error)
	GetByRoles(ctx context.Context, id string) ([]MenuRoles, error)
	Query(ctx context.Context, offset, limit int) ([]Menu, error)
	Count(ctx context.Context) (int, error)
	Create(ctx context.Context, input CreateRequest) error
	Update(ctx context.Context, id string, input UpdateRequest) (Menu, error)
	Delete(ctx context.Context, id string) (Menu, error)
}

// Menu represents the data about an menu.
type Menu struct {
	entity.Menu
}

// Menu represents the data about an menu.
type MenuRoles struct {
	entity.MenuRoles
}

// CreateRequest represents an menu creation request.
type CreateRequest struct {
	Name     string `json:"name"`
	Path     string `json:"path"`
	Icon     string `json:"icon"`
	IsActive int    `json:"is_active"`
	ParentId int    `json:"parent_id"`
	Sorting  int    `json:"sorting"`
}

// Validate validates the CreateRequest fields.
func (m CreateRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Name, validation.Required, validation.Length(0, 128)),
	)
}

// UpdateRequest represents an menu update request.
type UpdateRequest struct {
	Name     string `json:"name"`
	Path     string `json:"path"`
	Icon     string `json:"icon"`
	IsActive int    `json:"is_active"`
	ParentId int    `json:"parent_id"`
	Sorting  int    `json:"sorting"`
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

// NewService creates a new menu service.
func NewService(repo Repository, logger log.Logger) Service {
	return service{repo, logger}
}

// Get returns the menu with the specified the menu ID.
func (s service) Get(ctx context.Context, id string) (Menu, error) {
	menu, err := s.repo.Get(ctx, id)
	if err != nil {
		return Menu{}, err
	}
	return Menu{menu}, nil
}

// Create creates a new menu.
func (s service) Create(ctx context.Context, req CreateRequest) error {
	if err := req.Validate(); err != nil {
		return err
	}
	//id := entity.GenerateID()
	//now := time.Now()
	err := s.repo.Create(ctx, entity.Menu{
		Name:     req.Name,
		Path:     req.Path,
		Icon:     &req.Icon,
		IsActive: &req.IsActive,
		Sorting:  &req.Sorting,
	})

	return err
}

// Update updates the menu with the specified ID.
func (s service) Update(ctx context.Context, id string, req UpdateRequest) (Menu, error) {
	if err := req.Validate(); err != nil {
		return Menu{}, err
	}

	menu, err := s.Get(ctx, id)
	if err != nil {
		return menu, err
	}

	menu.Name = req.Name
	menu.Path = req.Path
	menu.Icon = &req.Icon
	menu.IsActive = &req.IsActive
	menu.Sorting = &req.Sorting

	if err := s.repo.Update(ctx, menu.Menu); err != nil {
		return menu, err
	}
	return menu, nil
}

// Delete deletes the menu with the specified ID.
func (s service) Delete(ctx context.Context, id string) (Menu, error) {
	menu, err := s.Get(ctx, id)
	if err != nil {
		return Menu{}, err
	}
	if err = s.repo.Delete(ctx, id); err != nil {
		return Menu{}, err
	}
	return menu, nil
}

// Count returns the number of menus.
func (s service) Count(ctx context.Context) (int, error) {
	return s.repo.Count(ctx)
}

// Query returns the menus with the specified offset and limit.
func (s service) Query(ctx context.Context, offset, limit int) ([]Menu, error) {
	items, err := s.repo.Query(ctx, offset, limit)
	if err != nil {
		return nil, err
	}
	result := []Menu{}
	for _, item := range items {
		result = append(result, Menu{item})
	}
	return result, nil
}

// Query returns the menus with the specified offset and limit.
func (s service) GetByRoles(ctx context.Context, id string) ([]MenuRoles, error) {
	items, err := s.repo.GetByRoles(ctx, id)
	if err != nil {
		return nil, err
	}
	result := []MenuRoles{}
	for _, item := range items {
		result = append(result, MenuRoles{item})
	}
	return result, nil
}
