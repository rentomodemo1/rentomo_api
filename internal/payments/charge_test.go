package payments

import "testing"

func TestChargeUsesInjectedPayment(t *testing.T) {
	called := ""
	s := New(func(id string) error { called = id; return nil }, func(string) {})
	if err := s.Charge("b"); err != nil || called != "b" {
		t.Fatal(called, err)
	}
}
