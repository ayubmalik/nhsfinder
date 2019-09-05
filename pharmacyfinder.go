package nhsfinder

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

// Address is a UK address
type Address struct {
	Line1    string
	Line2    string
	Line3    string
	Line4    string
	Line5    string
	Postcode string
}

// Pharmacy in the UK
type Pharmacy struct {
	ODSCode       string
	ParentODSCode string
	Name          string
	Address       *Address
	Phone         string
	LatLng        LatLng
}

// SearchResult is a given item and distance from search query
type SearchResult struct {
	Distance float64
	Pharmacy Pharmacy
}

// PharmacyFinder finds the nearest pharmacies
type PharmacyFinder struct {
	LatLngs    map[string]LatLng
	Pharmacies []Pharmacy
}

// FindNearest finds nearest 10
func (pf PharmacyFinder) FindNearest(searchPostcode string) []SearchResult {
	distances := make(map[float64]Pharmacy, 0)
	start := time.Now()
	for _, pharmacy := range pf.Pharmacies {
		fromLatLng := pf.LatLngs[searchPostcode]
		dist := Distance(fromLatLng, pharmacy.LatLng)
		distances[dist] = pharmacy
	}

	end := time.Now().Sub(start)
	log.Printf("Calculated %d distances from %s in %v\n", len(distances), searchPostcode, end)

	keys := make([]float64, 0, len(distances))
	for k := range distances {
		keys = append(keys, k)
	}
	sort.Float64s(keys)
	result := make([]SearchResult, 0)
	for _, key := range keys[0:10] {
		p := distances[key]
		result = append(result, SearchResult{key, p})
	}
	return result
}
