package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ayubmalik/pharmacyfinder"
)

// TODO: delete this
func main() {
	fmt.Println("Loading data...")
	latLngs := pharmacyfinder.LoadLatLngs("data/ukpostcodes.csv")
	fmt.Printf("Loaded %d postcodes\n", len(latLngs))

	pharmacies := pharmacyfinder.LoadPharmacies("data/pharmacies.csv")
	fmt.Printf("Loaded %d pharmacies with lat/lng\n", len(pharmacies))

	pcode1 := "BD18 2DS"
	pcode2 := "M4 4BF"
	from := latLngs[pcode1]
	to := latLngs[pcode2]
	dist1 := pharmacyfinder.Distance(from, to)
	fmt.Printf("Distance from '%s' to '%s': %fm\n", pcode1, pcode2, dist1)

	inmem := pharmacyfinder.InMemFinder{LatLngs: latLngs, Pharmacies: pharmacies}
	fmt.Println()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter postcode in format M4 4BF: ")
		pcode, _ := reader.ReadString('\n')
		pcode = strings.Replace(pcode, "\n", "", -1)
		pcode = strings.ToUpper(pcode)
		if len(pcode) < 2 {
			fmt.Println("Goodbye!")
			return
		}

		results := inmem.ByPostcode(pcode)
		display(results)
	}
}

func display(results []pharmacyfinder.FindResult) {
	for i, r := range results {
		fmt.Printf("%2d %7.2f %-30s %-30s %s\n", i, r.Distance, r.Pharmacy.Name, r.Pharmacy.Address.Line1, r.Pharmacy.Address.Postcode)
	}
}
