package pharmacyfinder

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

// LatLng is point with latitude and longitude
type LatLng struct {
	Lat float64
	Lng float64
}

// Address in England
type Address struct {
	Line1    string
	Line2    string
	Line3    string
	Line4    string
	Line5    string
	Postcode string
}

// Pharmacy in England
type Pharmacy struct {
	ODSCode string
	Name    string
	Address Address
	Phone   string
	LatLng  LatLng
}

// FindResult is a given item and distance from search query
type FindResult struct {
	Distance float64
	Pharmacy Pharmacy
}

type finder interface {
	ByPostcode(postcode string) []FindResult
}

// InMemFinder is an in memory finder
type InMemFinder struct {
	LatLngs    map[string]LatLng
	Pharmacies []Pharmacy
}

// NewInMemFinder returns an in memory finder with files loaded from default paths
// TODO: add dataDir param from viper
func NewInMemFinder() (*InMemFinder, error) {
	postcodeFile, err := os.Open("data/postcode.csv")
	if err != nil {
		return nil, err
	}
	defer postcodeFile.Close()

	latLngs, err := LoadLatLngs(postcodeFile)
	if err != nil {
		return nil, err
	}

	pharmacyFile, err := os.Open("data/pharmacy.csv")
	if err != nil {
		return nil, err
	}
	defer pharmacyFile.Close()

	pharmacies, err := LoadPharmacies(pharmacyFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not load pharmacy file: %v", err)
		os.Exit(1)
	}
	return &InMemFinder{LatLngs: latLngs, Pharmacies: pharmacies}, nil
}

// FindPharmacy finds nearest 10 TODO: should make param
func (pf InMemFinder) FindPharmacy(postcode string) []FindResult {
	postcode = strings.ReplaceAll(postcode, " ", "")
	distances := make(map[float64]Pharmacy)
	for _, pharmacy := range pf.Pharmacies {
		fromLatLng := pf.LatLngs[postcode]
		dist := Distance(fromLatLng, pharmacy.LatLng)
		distances[dist] = pharmacy
	}

	keys := make([]float64, 0, len(distances))
	for k := range distances {
		keys = append(keys, k)
	}
	sort.Float64s(keys)

	max := 10
	if len(keys) < max {
		max = len(keys)
	}

	result := make([]FindResult, 0, max)
	for i := 0; i < max; i++ {
		p := distances[keys[i]]
		result = append(result, FindResult{keys[i], p})
	}
	return result
}
