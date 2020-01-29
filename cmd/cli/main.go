package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	finder "github.com/ayubmalik/finder"
)

// TODO: delete this
func main() {
	fmt.Println("Loading data...")
	latLngs := finder.LoadLatLngs("data/ukpostcodes.csv")
	fmt.Printf("Loaded %d postcodes\n", len(latLngs))

	pharmacies := finder.LoadPharmacies("data/pharmacies.csv")
	fmt.Printf("Loaded %d pharmacies with lat/lng\n", len(pharmacies))

	pcode1 := "BD18 2DS"
	pcode2 := "M4 4BF"
	from := latLngs[pcode1]
	to := latLngs[pcode2]
	dist1 := finder.Distance(from, to)
	fmt.Printf("Distance from '%s' to '%s': %fm\n", pcode1, pcode2, dist1)

	find := finder.InMemFinder{LatLngs: latLngs, Pharmacies: pharmacies}
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

		results := find.ByPostcode(pcode)
		display(results)
	}
}

