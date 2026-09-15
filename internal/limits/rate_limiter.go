package limits

import (
	"net/http"
	"sync"
	"time"
)

type Limiter struct {
	mu           sync.Mutex
	limit, count int
	window       time.Duration
	start        time.Time
	Now          func() time.Time
}

func New(limit int, window time.Duration) *Limiter {
	return &Limiter{limit: limit, window: window, Now: time.Now}
}
func (l *Limiter) Wrap(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		l.mu.Lock()
		now := l.Now()
		if l.start.IsZero() || !now.Before(l.start.Add(l.window)) {
			l.start = now
			l.count = 0
		}
		l.count++
		allowed := l.count <= l.limit
		l.mu.Unlock()
		if !allowed {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusTooManyRequests)
			w.Write([]byte("{\"error\":\"Request limit exceeded. Try again later.\"}"))
			return
		}
		next.ServeHTTP(w, r)
	})
}
