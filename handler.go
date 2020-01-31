package nhsfinder

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

// FinderHandler handles pharmacy http API
type FinderHandler struct {
	pharmacyFinder pharmacyFinder
	gpFinder       gpFinder
	http.Handler
}

func (h *FinderHandler) findPharmacies(w http.ResponseWriter, r *http.Request) {
	postcode := pat.Param(r, "postcode")
	if len(postcode) < minlen || len(postcode) > maxlen {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	results := h.pharmacyFinder.FindPharmacy(postcode)
	json.NewEncoder(w).Encode(results)
}

// NewFinderHandler creates a http.Handler for finding NHS organisations
func NewFinderHandler(pf pharmacyFinder) *FinderHandler {
	h := new(FinderHandler)
	h.pharmacyFinder = pf
	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/pharmacies/postcode/:postcode"), h.findPharmacies)
	h.Handler = mux
	return h
}
