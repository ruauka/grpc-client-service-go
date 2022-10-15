// Package functions - пакет вспомогательных функций.
package functions

import (
	"time"

	"github.com/go-playground/validator/v10"

	"service/internal/strategy/request"
	"service/internal/strategy/utils/dictionary"
)

// DateDiff - вспомогательная функция расчета разницы в датах.
func DateDiff(date1, date2 time.Time) int32 {
	duration := date2.Sub(date1).Hours() / dictionary.HourDivider
	return int32(duration)
}

// Validate - валидация структуры Request.
func Validate(req *request.Request) error {
	err := dictionary.Validate.Struct(req)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			return err
		}
	}
	return nil
}
