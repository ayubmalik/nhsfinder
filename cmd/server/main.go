package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	finder "github.com/ayubmalik/pharmacyfinder"

	goji "goji.io"
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
	postcodes := finder.LoadPostcodes("data/ukpostcodes.csv")
	pharmacies := finder.LoadPharmacies("data/Pharmacy.csv")
	finderRoute := finderRoute{finder.InMemFinder{postcodes, pharmacies}}

	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/find-pharmacies/:postcode"), finderRoute.serveHTTP)
	mux.HandleFunc(pat.Get("/"), search)
	fmt.Println("Started server API: http://localhost:8000/")
	fmt.Println("Search API: http://localhost:8000/find-pharmacies/:postcode")
	http.ListenAndServe("0.0.0.0:8000", mux)
}
