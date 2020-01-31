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
	finder finder
	http.Handler
}

func (h *FinderHandler) findPharmacies(w http.ResponseWriter, r *http.Request) {
	postcode := pat.Param(r, "postcode")
	if len(postcode) < minlen || len(postcode) > maxlen {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	results := h.finder.FindPharmacies(postcode)
	json.NewEncoder(w).Encode(results)
}

// NewFinderHandler creates a http.Handler for finding NHS organisations
func NewFinderHandler(f finder) *FinderHandler {
	h := new(FinderHandler)
	h.finder = f
	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/pharmacies/postcode/:postcode"), h.findPharmacies)
	h.Handler = mux
	return h
}
