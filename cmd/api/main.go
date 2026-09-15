package main

import (
	"github.com/rentomodemo1/rentomo_api/internal/config"
	"github.com/rentomodemo1/rentomo_api/internal/fleet"
	"github.com/rentomodemo1/rentomo_api/internal/payments"
	"log"
	"net/http"
)

func routes() http.Handler {
	mux := http.NewServeMux()
	service := payments.New(func(string) error { return nil }, func(string) {})
	mux.HandleFunc("/charge", service.Handler)
	mux.HandleFunc("/availability", fleet.Handler)
	return mux
}
func main() { log.Fatal(http.ListenAndServe(config.Load().Address, routes())) }
