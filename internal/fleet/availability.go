package fleet

import (
	"encoding/json"
	"net/http"
)

type Vehicle struct {
	ID, Branch string
	Dates      []string
	SoldOut    bool
}

func Inventory() []Vehicle {
	return []Vehicle{{ID: "v1", Branch: "north", Dates: []string{"2030-06-01"}}, {ID: "v2", Branch: "north", Dates: []string{"2030-06-01"}, SoldOut: true}}
}
func Available(v Vehicle, date string) bool {
	for _, d := range v.Dates {
		if d == date {
			return !v.SoldOut
		}
	}
	return false
}
func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Search(Inventory(), r.URL.Query().Get("branch"), r.URL.Query().Get("date")))
}
