package main

import (
	"errors"
	"net/http"
	"time"

	health "github.com/docker/go-healthcheck"
	"github.com/gorilla/mux"
)

func main() {

	health.RegisterPeriodicThresholdFunc("postgresql", time.Second*5, 3, postgresqlCheck)
	health.RegisterPeriodicThresholdFunc("gateway", time.Second*5, 3, gateWayCheck)

	r := mux.NewRouter()
	r.HandleFunc("/health", health.StatusHandler)
	srv := &http.Server{
		Handler:     r,
		Addr:        "0.0.0.0:8080",
		ReadTimeout: 15 * time.Second,
	}

	srv.ListenAndServe()
}

func postgresqlCheck() error {
	return nil
}

func gateWayCheck() error {
	return errors.New("no connection")
}
