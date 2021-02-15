package entity

import (
	"time"
)

// User represents a user.
type History struct {
	TransactionId        string     `json:"transaction_id" db:"pk"`
	ChannelTransactionId string     `json:"channel_transaction_id"`
	Msisdn               string     `json:"msisdn"`
	InitialDate          *time.Time `json:"initial_date"`
	DueDate              *time.Time `json:"due_date"`
	PaidStatus           int        `json:"paid_status"`
	PaidDate             *time.Time `json:"paid_date"`
	LoanValue            int        `json:"loan_value"`
	ActivityState        *string    `json:"activity_state"`
	TrxTime              *time.Time `json:"trx_time"`
	Remarks              *string    `json:"remarks"`
	Exposure             *int       `json:"exposure"`
	Metadata             *string    `json:"metadata"`
	Product_offer        *string    `json:"product_offer"`
	MumOfProductOffer    *int       `json:"num_of_product_offer"`
	EligibleChannel      *string    `json:"eligible_channel"`
	MaximumSequence      *int       `json:"maximum_sequence"`
	UpdatedAtScv         *time.Time `json:"updated_at_scv"`
	CurrSequence         *int       `json:"curr_sequence"`
	LatestOfferProfile   *string    `json:"latest_offer_profile"`
}

// get table real
func (c History) TableName() string {
	return "history"
}
