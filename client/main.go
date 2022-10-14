// Package main
package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/julienschmidt/httprouter"

	s "client/internal/adapters/grpc/strategy"
	"client/internal/adapters/router"
)

func main() {
	strategy, err := s.NewStrategy()
	if err != nil {
		log.Fatal(err)
	}
	handler := router.NewHandler(strategy)

	r := httprouter.New()

	r.POST("/execute", handler.Call)
	r.GET("/health", handler.HealthCheck)

	loggedRouter := handlers.LoggingHandler(os.Stdout, r)

	log.Println("Starting server...")

	log.Fatal(http.ListenAndServe(":8000", loggedRouter)) //nolint:gosec
}
