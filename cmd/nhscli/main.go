package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ayubmalik/nhsfinder"
)

func main() {
	fmt.Println("Loading data...")
	postcodesfile := "data/ukpostcodes.csv"
	pcodeLatLngs := nhsfinder.LoadPostcodes(postcodesfile)
	fmt.Printf("Loaded %d postcodes\n", len(pcodeLatLngs))

	pharmaciesfile := "data/Pharmacy.csv"
	pharmacies := nhsfinder.LoadPharmacies(pharmaciesfile)
	fmt.Printf("Loaded %d pharmacies with lat/lng\n", len(pharmacies))

	pcode1 := "BD18 2DS"
	pcode2 := "M4 4BF"
	from := pcodeLatLngs[pcode1]
	to := pcodeLatLngs[pcode2]
	dist1 := nhsfinder.Distance(from, to)
	fmt.Printf("Distance from '%s' to '%s': %fm\n", pcode1, pcode2, dist1)

	finder := nhsfinder.PharmacyFinder{pcodeLatLngs, pharmacies}
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

		results := finder.FindNearest(pcode)
		display(results)
	}
}

func display(results []nhsfinder.SearchResult) {
	for i, r := range results {
		fmt.Printf("%2d %7.2f %-30s %-30s %s\n", i, r.Distance, r.Pharmacy.Name, r.Pharmacy.Address.Line1, r.Pharmacy.Address.Postcode.Value)
	}
}
