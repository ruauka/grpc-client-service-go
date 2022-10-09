// Package entities - входящая структура.
package entities

import "time"

// Request - структура для входящего JSON.
type Request struct {
	Name     string    `json:"name" validate:"required"`
	Surname  string    `json:"surname" validate:"required"`
	Accounts []Account `json:"account" validate:"required,dive,required"`
	Loans    []Loan    `json:"loan" validate:"required,dive,required"`
}

// Account - Структура с данными по счетам.
type Account struct {
	PaymentDateBalance *int32 `json:"payment_date_balance" validate:"required"`
}

// Loan - структура с данными по кредиту.
type Loan struct {
	PaymentDate       time.Time `json:"payment_date" validate:"required"`
	ActualPaymentDate time.Time `json:"actual_payment_date" validate:"required"`
	Payment           int32     `json:"payment" validate:"required"`
}
