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
	"client/internal/config"
)

func main() {
	// config create
	cfg := config.GetConfig()
	// create gRPC service instance and connect
	strategy, err := s.NewStrategy(cfg)
	if err != nil {
		log.Fatal(err)
	}
	// create handler
	handler := router.NewHandler(strategy)
	// create router
	r := httprouter.New()
	// register health check endpoint
	r.POST("/execute", handler.Call)
	// register main endpoint for gRPC
	r.GET("/health", handler.HealthCheck)
	// logger
	loggedRouter := handlers.LoggingHandler(os.Stdout, r)

	log.Println("Starting server...")

	log.Fatal(http.ListenAndServe(":8000", loggedRouter)) //nolint:gosec
}
