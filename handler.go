package pharmacyfinder

import (
	"encoding/json"
	"net/http"

	"goji.io"
	"goji.io/pat"
)

const (
	minlen = 5
	maxlen = 8
)

// PharmacyHandler handles pharmacy http API
type PharmacyHandler struct {
	finder finder
	http.Handler
}

func (ph *PharmacyHandler) findByPostcode(w http.ResponseWriter, r *http.Request) {
	postcode := pat.Param(r, "postcode")
	if len(postcode) < minlen || len(postcode) > maxlen {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	results := ph.finder.ByPostcode(postcode)
	json.NewEncoder(w).Encode(results)
}

// NewPharmacyHandler constructor to create handlers
func NewPharmacyHandler(finder finder) *PharmacyHandler {
	ph := new(PharmacyHandler)
	ph.finder = finder
	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/pharmacies/postcode/:postcode"), ph.findByPostcode)
	ph.Handler = mux
	return ph
}
