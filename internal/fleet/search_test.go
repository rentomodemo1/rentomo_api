package fleet

import (
	"encoding/json"
	"net/http/httptest"
	"testing"
)

func TestSearchByBranchAndDateHidesSoldOut(t *testing.T) {
	vehicles := append(Inventory(), Vehicle{ID: "other", Branch: "south", Dates: []string{"2030-06-01"}})
	matches := Search(vehicles, "north", "2030-06-01")
	if len(matches) != 1 || matches[0].ID != "v1" {
		t.Fatal(matches)
	}
	if len(Search(vehicles, "north", "2030-06-02")) != 0 {
		t.Fatal("wrong date included")
	}
	w := httptest.NewRecorder()
	Handler(w, httptest.NewRequest("GET", "/availability?branch=north&date=2030-06-01", nil))
	var got []Vehicle
	if err := json.Unmarshal(w.Body.Bytes(), &got); err != nil || len(got) != 1 || got[0].ID != "v1" {
		t.Fatal(got, err)
	}
}
