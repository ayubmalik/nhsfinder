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

// *********************************************************************** //
// TODO use type alias as GP and Pharmacy are currently identical structs? //
// *********************************************************************** //

// Org is an NHS organisation in England.
type Org struct {
	ODSCode string
	Name    string
	Address Address
	Phone   string
	LatLng  LatLng
}

// Pharmacy is pharmacy in England.
type Pharmacy Org

// GP is General Practioner in England.
type GP Org

// FindResult is a given item and distance from search query
type FindResult struct {
	Distance float64
	Pharmacy Pharmacy
	GP       GP
}

type finder interface {
	ByPostcode(postcode string) []FindResult
}

// InMemFinder is an in memory finder
type InMemFinder struct {
	LatLngs    map[string]LatLng
	Pharmacies []Pharmacy
	GPs        []GP
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
func (f InMemFinder) FindPharmacy(postcode string) []FindResult {
	postcode = strings.ReplaceAll(postcode, " ", "")
	distances := make(map[float64]Pharmacy)
	for _, pharmacy := range f.Pharmacies {
		from := f.LatLngs[postcode]
		dist := Distance(from, pharmacy.LatLng)
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
		result = append(result, FindResult{Distance: keys[i], Pharmacy: p})
	}
	return result
}

// FindGPs finds GPs in England.
// 10 TODO: should make param
func (f InMemFinder) FindGPs(postcode string) []FindResult {
	postcode = strings.ReplaceAll(postcode, " ", "")
	distances := make(map[float64]GP)
	for _, gp := range f.GPs {
		from := f.LatLngs[postcode]
		dist := Distance(from, gp.LatLng)
		distances[dist] = gp
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
		gp := distances[keys[i]]
		result = append(result, FindResult{Distance: keys[i], GP: gp})
	}
	return result
}
