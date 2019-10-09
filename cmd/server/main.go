package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	finder "github.com/ayubmalik/pharmacyfinder"

	"goji.io/pat"
)

type finderRoute struct {
	finder finder.Finder
}

func (fr finderRoute) serveHTTP(w http.ResponseWriter, r *http.Request) {
	postcode := strings.ToUpper(pat.Param(r, "postcode"))
	postcode = strings.Replace(postcode, "+", " ", -1) // allow M4+4BF
	result := fr.finder.ByPostcode(postcode)
	jsonOut, _ := json.Marshal(result)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, string(jsonOut))
}

func search(w http.ResponseWriter, r *http.Request) {
	content, err := ioutil.ReadFile("search.html")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, string(content))
}

func main() {
	log.Println("Loading data from CSV files")
	latLngs := finder.LoadLatLngs("data/ukpostcodes.csv")
	pharmacies := finder.LoadPharmacies("data/Pharmacy.csv")
	inMemFinder := finder.InMemFinder{LatLngs: latLngs, Pharmacies: pharmacies}
	handler := finder.NewPharmacyHandler(&inMemFinder)

	fmt.Println("Started Pharmacys API: http://localhost:8000/pharmacies/:postcode")
	http.ListenAndServe("0.0.0.0:8000", handler)
}
