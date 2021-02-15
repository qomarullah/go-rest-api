package entity

import (
	"time"
)

// User represents a user.
type Customer struct {
	Msisdn                  string     `json:"msisdn" db:"pk"`
	Exposure                int        `json:"exposure"`
	MetaData                string     `json:"meta_data"`
	ProductOffer            string     `json:"product_offer"`
	NumOfProductOffer       int        `json:"num_of_product_offer"`
	EligibleChannel         string     `json:"eligible_channel"`
	MaximumSequence         int        `json:"maximum_sequence"`
	CurrSequence            int        `json:"curr_sequence"`
	PaidStatus              int        `json:"paid_status"`
	LatestInitiateOfferTime *time.Time `json:"latest_initiate_offer_time"`
	ValidUntil              time.Time  `json:"valid_until"`
	UpdatedAtScv            time.Time  `json:"updated_at_scv"`
	State                   string     `json:"state"`
	LatestTrxid             *string    `json:"latest_trxid"`
	LatestCtrxid            *string    `json:"latest_ctrxid"`
	LatestLoanValue         int        `json:"latest_loan_value"`
	LatestDueDate           *time.Time `json:"latest_due_date"`
	LatestPaidDate          *time.Time `json:"latest_paid_date"`
	LatestOfferProfile      *string    `json:"latest_offer_profile"`
	OfferIdList             *string    `json:"offer_id_list"`
	OfferUmbMenu            *string    `json:"offer_umb_menu"`
	OfferIdChoose           *string    `json:"offer_id_choose"`
	PaymentIdChoose         *string    `json:"payment_id_choose"`
}

// get table real
func (c Customer) TableName() string {
	return "profile"
}
