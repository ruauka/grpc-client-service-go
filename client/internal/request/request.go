// Package request - входящая структура.
package request

import (
	"time"

	"github.com/go-playground/validator/v10"
)

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
	Payment           int32     `json:"payment" validate:"required"`
	PaymentDate       time.Time `json:"payment_date" validate:"required"`
	ActualPaymentDate time.Time `json:"actual_payment_date" validate:"required"`
}

// Validate - движок валидации.
var Validate = validator.New()

// ValidateStruct - валидация структуры Request.
func ValidateStruct(inputMessage *Request) error {
	err := Validate.Struct(inputMessage)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return err
		}
	}
	return nil
}
