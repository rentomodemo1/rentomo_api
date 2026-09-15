package payments

import (
	"errors"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRetriesTwiceThenNotifies(t *testing.T) {
	calls, notices := 0, 0
	s := New(func(string) error { calls++; return errors.New("card declined") }, func(id string) {
		notices++
		if calls != 3 || id != "b" {
			t.Fatal("notification before final failure")
		}
	})
	r := httptest.NewRequest("POST", "/charge", strings.NewReader("booking=b"))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	s.Handler(w, r)
	if calls != 3 || notices != 1 || w.Code != 502 {
		t.Fatal(calls, notices, w.Code)
	}
}
func TestSuccessfulAttemptStopsRetries(t *testing.T) {
	for success := 1; success <= 3; success++ {
		calls, notices := 0, 0
		s := New(func(string) error {
			calls++
			if calls == success {
				return nil
			}
			return errors.New("temporary decline")
		}, func(string) { notices++ })
		if err := s.Charge("b"); err != nil || calls != success || notices != 0 {
			t.Fatal(calls, notices, err)
		}
	}
}
