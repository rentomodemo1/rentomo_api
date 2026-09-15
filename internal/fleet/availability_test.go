package fleet

import "testing"

func TestAvailabilityDate(t *testing.T) {
	v := Inventory()[0]
	if !Available(v, "2030-06-01") || Available(v, "2030-06-02") {
		t.Fatal("date availability mismatch")
	}
}
