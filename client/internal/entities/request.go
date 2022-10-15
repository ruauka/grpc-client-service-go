// Package entities - input and output structs.
package entities

import "time"

// Request - input JSON struct.
type Request struct {
	Name     string    `json:"name" validate:"required"`
	Surname  string    `json:"surname" validate:"required"`
	Accounts []Account `json:"account" validate:"required,dive,required"`
	Loans    []Loan    `json:"loan" validate:"required,dive,required"`
}

// Account - struct with account data.
type Account struct {
	PaymentDateBalance *int32 `json:"payment_date_balance" validate:"required"`
}

// Loan - struct with loan data.
type Loan struct {
	PaymentDate       time.Time `json:"payment_date" validate:"required"`
	ActualPaymentDate time.Time `json:"actual_payment_date" validate:"required"`
	Payment           int32     `json:"payment" validate:"required"`
}
