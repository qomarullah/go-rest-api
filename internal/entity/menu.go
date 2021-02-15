package entity

import (
	"time"
)

// User represents a user.
type Menu struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Path      string    `json:"path"`
	Icon      *string   `json:"icon"`
	IsActive  *int      `json:"is_active"`
	ParentId  *int      `json:"parent_id"`
	Sorting   *int      `json:"sorting"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	//Roles     MenuRoles `db:"-" json:"roles,omitempty"`
}
type MenuRoles struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Path      string    `json:"path"`
	Icon      *string   `json:"icon"`
	IsActive  *int      `json:"is_active"`
	ParentId  *int      `json:"parent_id"`
	Sorting   *int      `json:"sorting"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	IsCreate  *bool     `json:"is_create,omitempty"`
	IsRead    *bool     `json:"is_read,omitempty"`
	IsEdit    *bool     `json:"is_edit,omitempty"`
	IsDelete  *bool     `json:"is_delete,omitempty"`
}

// get table real
func (c Menu) TableName() string {
	return "menus"
}
