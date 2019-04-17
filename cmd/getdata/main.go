package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/ayubmalik/nhsfinder"
)

var postcodes map[string]nhsfinder.Postcode
var pharmacies []nhsfinder.Pharmacy

func search(postcodeValue string) {
	var distances = make(map[float64]nhsfinder.Pharmacy)
	start := time.Now()
	for _, pharmacy := range pharmacies {
		postcode := postcodes[postcodeValue]
		dist := nhsfinder.Distance(postcode.LatLng, pharmacy.Address.Postcode.LatLng)
		distances[dist] = pharmacy
	}
	end := time.Now().Sub(start)
	fmt.Printf("Calculated %d distances from %s\n", len(distances), postcodeValue)
	fmt.Printf("Took %v ms\n", end)

	keys := make([]float64, 0, len(distances))
	for k := range distances {
		keys = append(keys, k)
	}
	sort.Float64s(keys)
	for _, key := range keys[0:10] {
		p := distances[key]
		fmt.Printf("%7.2f %s %35s %25s %s\n", key, p.ID, p.Name, p.Address.Line1, p.Address.Postcode.Value)
	}
}

func main() {
	fmt.Println("Loading data...")
	postcodesfile := "data/ukpostcodes.csv"
	postcodes := nhsfinder.LoadPostcodes(postcodesfile)
	fmt.Printf("Loaded %d postcodes\n", len(postcodes))

	pharmaciesfile := "data/Pharmacy.csv"
	pharmacies = nhsfinder.LoadPharmacies(pharmaciesfile)
	fmt.Printf("Loaded %d pharmacies with lat/lng\n", len(pharmacies))

	pcode1 := postcodes["M4 4BF"]
	pcode2 := postcodes["LS2 7UE"]
	dist1 := nhsfinder.Distance(pcode1.LatLng, pcode2.LatLng)
	fmt.Printf("Distance from '%s' to '%s': %fm\n", pcode1.Value, pcode2.Value, dist1)

	fmt.Println()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter postcode with spaces: ")
		pcode, _ := reader.ReadString('\n')
		pcode = strings.Replace(pcode, "\n", "", -1)
		pcode = strings.ToUpper(pcode)
		if len(pcode) < 2 {
			fmt.Println("Goodbye!")
			return
		}

		search(pcode)
	}
}
