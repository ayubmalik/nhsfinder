package nhsfinder

import (
	"encoding/json"
	"net/http"

	"goji.io"
	"goji.io/pat"
)

// PharmacyHandler handles pharmacy http API
type PharmacyHandler struct {
	finder Finder
	http.Handler
}

func (ph *PharmacyHandler) findByPostcode(w http.ResponseWriter, r *http.Request) {
	results := ph.finder.ByPostcode("TODO")
	json.NewEncoder(w).Encode(results)
}

// NewPharmacyHandler constructor to create handlers
func NewPharmacyHandler(finder Finder) *PharmacyHandler {
	ph := new(PharmacyHandler)
	ph.finder = finder
	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/pharmacies/postcode/:postcode"), ph.findByPostcode)
	ph.Handler = mux
	return ph
}
