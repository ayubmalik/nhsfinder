package pharmacyfinder

import (
	"log"
	"sort"
	"time"
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
	ODSCode       string
	ParentODSCode string
	Name          string
	Address       *Address
	Phone         string
	LatLng        LatLng
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

// ByPostcode finds nearest 10
func (pf *InMemFinder) ByPostcode(postcode string) []FindResult {
	distances := make(map[float64]Pharmacy)
	start := time.Now()
	for _, pharmacy := range pf.Pharmacies {
		fromLatLng := pf.LatLngs[postcode]
		dist := Distance(fromLatLng, pharmacy.LatLng)
		distances[dist] = pharmacy
	}
	end := time.Now().Sub(start)
	log.Printf("Calculated %d distances from %s in %v\n", len(distances), postcode, end)

	keys := make([]float64, 0, len(distances))
	for k := range distances {
		keys = append(keys, k)
	}

	sort.Float64s(keys)
	result := make([]FindResult, 0)
	for _, key := range keys {
		p := distances[key]
		result = append(result, FindResult{key, p})
	}
	return result
}
