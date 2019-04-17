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

// Postcode is a single UK postcode with latlng
type Postcode struct {
	Value  string
	LatLng LatLng
}

// Address is a UK address
type Address struct {
	Line1    string
	Line2    string
	Line3    string
	Line4    string
	Line5    string
	Postcode Postcode
}

// Pharmacy in the UK
type Pharmacy struct {
	ID      string
	Name    string
	Address *Address
	Phone   string
}

// SearchResult is a given item and distance from search query
type SearchResult struct {
	pharmacy Pharmacy
	distance float64
}

// PharmacyFinder finds the nearest pharmacies
type PharmacyFinder struct {
	postcodes  map[string]Postcode
	pharmacies []Pharmacy
}

func (pf PharmacyFinder) findNearest(searchPostcode string) []SearchResult {
	distances := make(map[float64]Pharmacy)
	start := time.Now()
	for _, p := range pf.pharmacies {
		postcode := pf.postcodes[searchPostcode]
		if p.Address.Postcode.LatLng.Lat != 0 && p.Address.Postcode.LatLng.Lng != 0 {
			dist := Distance(postcode.LatLng, p.Address.Postcode.LatLng)
			distances[dist] = p
		}
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
		result = append(result, SearchResult{p, key})
	}
	return result
}
