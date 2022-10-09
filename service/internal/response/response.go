// Package response - ответ сервиса.
package response

import (
	"time"

	"service/internal/logic"
	"service/internal/utils/dictionary"
)

// Response - Структура ответа.
type Response struct {
	Version     string   `json:"version"`
	ExecuteDate string   `json:"execute_date"`
	Results     []Result `json:"resp"`
}

// Result - блок из основной логики.
type Result struct {
	logic.Result
}

// NewResponse - конструктор объекта ответа.
func NewResponse(data *logic.Data) *Response {
	return &Response{
		Version:     "v.1.0.0",
		ExecuteDate: time.Now().Format(dictionary.LayoutToString),
		Results: []Result{
			{
				data.Result,
			},
		},
	}
}
