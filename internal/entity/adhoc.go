package entity

import (
	"time"
)

// User represents a user.
type Adhoc struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Filename    string    `json:"filename"`
	Status      string    `json:"status"`
	ScheduledAt string    `json:"scheduled_at"`
	CreatedAt   time.Time `db:"-" json:"created_at"`
	UpdatedAt   time.Time `db:"-" json:"updated_at"`
}

// get table real
func (c Adhoc) TableName() string {
	return "adhoc"
}
