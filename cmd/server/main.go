package main

import (
	"fmt"
	"log"
	"net/http"

	finder "github.com/ayubmalik/pharmacyfinder"
)

func main() {
	// TODO
	// w.Header().Set("Access-Control-Allow-Origin", "*")
	// postcode = strings.Replace(postcode, "+", " ", -1) // allow M4+4BF
	log.Println("Loading data from CSV files")
	latLngs := finder.LoadLatLngs("data/ukpostcodes.csv")
	pharmacies := finder.LoadPharmacies("data/pharmacies.csv")
	inMemFinder := finder.InMemFinder{LatLngs: latLngs, Pharmacies: pharmacies}
	handler := finder.NewPharmacyHandler(&inMemFinder)

	fmt.Println("Started Pharmacys API: http://localhost:8000/pharmacies/:postcode")
	http.ListenAndServe("0.0.0.0:8000", handler)
}
