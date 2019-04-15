package main

import (
	"fmt"

	"github.com/ayubmalik/nhsfinder"
)

func main() {
	// postcodedata := "data/ukpostcodes.csv"
	// postcodeDB := nhsfinder.LoadPostcodeDB(postcodedata)
	// fmt.Printf("Loaded %d postcodes", len(postcodeDB.Postcodes))

	pharmacydata := "data/Pharmacy.csv"
	pharmacies := nhsfinder.GetPharmacies(pharmacydata, nil)
	fmt.Printf("Loaded %d pharmacies with lat/lng\n", len(pharmacies))

	// pcode1 := postcodeDB.Postcodes["M4 4BF"]
	// pcode2 := postcodeDB.Postcodes["LS2 7UE"]
	// dist1 := nhsfinder.PostcodeDistance(pcode1, pcode2)
	// fmt.Printf("Distance from '%s' to '%s': %fm\n", pcode1.Value, pcode2.Value, dist1)

	// fmt.Println()
	// var distances = make(map[float64]nhsfinder.Pharmacy)
	// start := time.Now()
	// for _, p := range pharmacies {
	// 	if p.Address.Postcode.LatLng.Lat != 0 && p.Address.Postcode.LatLng.Lng != 0 {
	// 		// fmt.Println(p.Name, p.Address.Postcode)
	// 		dist := nhsfinder.PostcodeDistance(pcode1, p.Address.Postcode)
	// 		distances[dist] = p
	// 	}
	// }
	// end := time.Now().Sub(start)
	// fmt.Printf("Calculated %d distances from %s\n", len(distances), pcode1.Value)
	// fmt.Printf("Took %v ms\n", end)

	// keys := make([]float64, 0, len(distances))
	// for k := range distances {
	// 	keys = append(keys, k)
	// }
	// sort.Float64s(keys)
	// for _, key := range keys[0:10] {
	// 	p := distances[key]
	// 	fmt.Printf("%7.2f %35s %25s %s\n", key, p.Name, p.Address.Line1, p.Address.Postcode.Value)
	// }
}
