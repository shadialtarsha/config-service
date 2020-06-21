package api

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/shadialtarsha/config-service/service"
)

// NewServer runs httpServer.
func NewServer(config *service.Config) *http.Server {

	controller := Controller{Config: config}

	r := mux.NewRouter()

	r.HandleFunc("/api/v1/config/{serviceName}", controller.ReadConfig).Methods("GET")

	return &http.Server{
		Addr:         ":8080",
		Handler:      r,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
}
