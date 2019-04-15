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

var postcodeDB *nhsfinder.PostcodeDB
var pharmacies []nhsfinder.Pharmacy

func search(postcodeValue string) {
	var distances = make(map[float64]nhsfinder.Pharmacy)
	start := time.Now()
	for _, p := range pharmacies {
		postcode := postcodeDB.Postcodes[postcodeValue]
		if p.Address.Postcode.LatLng.Lat != 0 && p.Address.Postcode.LatLng.Lng != 0 {
			dist := nhsfinder.PostcodeDistance(postcode, p.Address.Postcode)
			distances[dist] = p
		}
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
	postcodedata := "data/ukpostcodes.csv"
	postcodeDB = nhsfinder.LoadPostcodeDB(postcodedata)
	fmt.Printf("Loaded %d postcodes\n", len(postcodeDB.Postcodes))

	pharmacydata := "data/Pharmacy.csv"
	pharmacies = nhsfinder.GetPharmacies(pharmacydata)
	fmt.Printf("Loaded %d pharmacies with lat/lng\n", len(pharmacies))

	pcode1 := postcodeDB.Postcodes["M4 4BF"]
	pcode2 := postcodeDB.Postcodes["LS2 7UE"]
	dist1 := nhsfinder.PostcodeDistance(pcode1, pcode2)
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
