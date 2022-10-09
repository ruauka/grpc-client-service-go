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

type Handler struct {
	strategy strategy.Strategy
}

func NewHandler(grpcConn strategy.Strategy) *Handler {
	return &Handler{
		strategy: grpcConn,
	}
}

func (h *Handler) Call(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Парсинг входящего JSON
	req := &request.Request{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	// Валидация объекта структуры Request
	if err := functions.Validate(req); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	resp, err := h.strategy.Execute(context.Background(), req)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp) //nolint:errcheck
}
