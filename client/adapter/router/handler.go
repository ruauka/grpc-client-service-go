// Package router - Package router
package router

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"client/adapter/grpc/strategy"
	"client/internal/request"
	"client/internal/utils/functions"
)

// Handler - структура хендлера.
type Handler struct {
	strategy strategy.Strategy
}

// NewHandler - конструктор структуры хендлера.
func NewHandler(grpcConn strategy.Strategy) *Handler {
	return &Handler{
		strategy: grpcConn,
	}
}

// Call - основной http хендлер.
func (h *Handler) Call(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Парсинг входящего JSON
	req := &request.Request{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Валидация объекта структуры Request
	if err := functions.Validate(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := h.strategy.Execute(context.Background(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp) //nolint:errcheck,gosec
}
