package entity

import (
	"time"
)

// User represents a user.
type Reminder struct {
	ID            string     `json:"id"`
	Msisdn        string     `json:"msisdn"`
	LoanValue     string     `json:"loan_value"`
	DueDate       *time.Time `json:"due_date"`
	Incremental   int        `json:"incremental"`
	TransactionID string     `json:"transaction_id"`
}

// get table real
func (c Reminder) TableName() string {
	return "queue"
}
