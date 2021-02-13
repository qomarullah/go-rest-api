package entity

import (
	"time"
)

// User represents a user.
type User struct {
	ID              string    `json:"id"`
	Name            string    `json:"name"`
	Email           string    `json:"email"`
	IdCmsPrivileges string    `json:"idCmsPrivileges"`
	Photo           *string   `json:"photo"`
	Status          *string   `json:"status"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
	Token           string    `json:"token"`
}

func (c User) TableName() string {
	return "cms_users"
}

// GetID returns the user ID.
func (u User) GetID() string {
	return u.ID
}

// GetName returns the user name.
func (u User) GetName() string {
	return u.Name
}
