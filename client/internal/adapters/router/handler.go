// Package router - Package router
package router

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"client/internal/adapters/grpc/strategy"
	"client/internal/entities"
	"client/internal/utils/functions"
)

// Handler - handler struct.
type Handler struct {
	strategy strategy.Strategy
}

// NewHandler - builder for handler struct.
func NewHandler(strategy strategy.Strategy) *Handler {
	return &Handler{
		strategy: strategy,
	}
}

// HealthCheck - handler checks gRPC service health.
func (h *Handler) HealthCheck(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if err := h.strategy.HealthCheck(context.Background()); err != nil {
		functions.MakeJSONResponse(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	functions.MakeJSONResponse(w, http.StatusOK, map[string]string{"status": "ok"})
}

// Call - main http handler.
func (h *Handler) Call(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Парсинг входящего JSON
	req := &entities.Request{}
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

	functions.MakeJSONResponse(w, http.StatusOK, resp)
}
