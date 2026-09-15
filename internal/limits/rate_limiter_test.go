package limits

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestLimitClearErrorAndWindowReset(t *testing.T) {
	now := time.Unix(1000, 0)
	l := New(2, time.Minute)
	l.Now = func() time.Time { return now }
	handler := l.Wrap(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("ok")) }))
	for i, want := range []int{200, 200, 429} {
		w := httptest.NewRecorder()
		handler.ServeHTTP(w, httptest.NewRequest("GET", "/availability", nil))
		if w.Code != want {
			t.Fatal(i, w.Code)
		}
		if want == 429 && !strings.Contains(w.Body.String(), "Request limit exceeded") {
			t.Fatal(w.Body.String())
		}
	}
	now = now.Add(time.Minute)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, httptest.NewRequest("GET", "/availability", nil))
	if w.Code != 200 {
		t.Fatal("window did not reset")
	}
}
