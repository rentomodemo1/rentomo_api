package payments

import (
	"net/http"
)

type PaymentFunc func(string) error
type Service struct {
	pay    PaymentFunc
	notify func(string)
}

func New(p PaymentFunc, n func(string)) *Service { return &Service{p, n} }
func (s *Service) Charge(id string) error        { return s.pay(id) }
func (s *Service) Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "use POST", 405)
		return
	}
	id := r.FormValue("booking")
	if id == "" {
		http.Error(w, "booking required", 400)
		return
	}
	if err := s.Charge(id); err != nil {
		http.Error(w, err.Error(), 502)
		return
	}
	w.Write([]byte("Payment accepted"))
}
