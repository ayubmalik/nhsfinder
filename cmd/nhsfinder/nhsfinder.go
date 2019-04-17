package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/ayubmalik/nhsfinder"

	goji "goji.io"
	"goji.io/pat"
)

type finderRoute struct {
	finder nhsfinder.PharmacyFinder
}

func (fr finderRoute) serveHTTP(w http.ResponseWriter, r *http.Request) {
	postcode := strings.ToUpper(pat.Param(r, "postcode"))
	postcode = strings.Replace(postcode, "+", " ", -1) // allow M4+4BF
	result := fr.finder.FindNearest(postcode)
	jsonOut, _ := json.Marshal(result)
	fmt.Fprintf(w, string(jsonOut))
}

func main() {
	log.Println("Loading data from CSV files")
	postcodes := nhsfinder.LoadPostcodes("data/ukpostcodes.csv")
	pharmacies := nhsfinder.LoadPharmacies("data/Pharmacy.csv")
	finderRoute := finderRoute{nhsfinder.PharmacyFinder{postcodes, pharmacies}}

	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/find-pharmacies/:postcode"), finderRoute.serveHTTP)
	fmt.Println("Started server API: http://localhost:8000/find-pharmacies/:postcode")
	http.ListenAndServe("localhost:8000", mux)
}
