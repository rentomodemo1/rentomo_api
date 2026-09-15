package main

import (
	"net/http/httptest"
	"testing"
)

func TestPublicRoutes(t *testing.T) {
	for _, url := range []string{"/charge", "/availability?branch=north&date=2030-06-01"} {
		w := httptest.NewRecorder()
		routes().ServeHTTP(w, httptest.NewRequest("GET", url, nil))
		if w.Code == 404 {
			t.Fatal("missing route", url)
		}
	}
}
