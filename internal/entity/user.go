package entity

import (
	"time"
)

// User represents a user.
type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	RolesId   int       `json:"roles_id"`
	Photo     *string   `json:"photo"`
	Status    *string   `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Password  string    `db:"-" json:"password,omitempty"`
	Token     string    `db:"-" json:"token,omitempty"`
}

// get table real
func (c User) TableName() string {
	return "users"
}

// GetID returns the user ID.
func (u User) GetID() string {
	return u.ID
}

// GetName returns the user name.
func (u User) GetName() string {
	return u.Name
}
