package data

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"

	"github.com/ayubmalik/nhsfinder"
)

// CreatePharmacies takes NHS ODS pharmacy CSV files and creates an condensed CSV of active pharmacies.
// The pharmacy CSV file contains only fields required by this app i.e. ODSCode, Name, Address 1, Address2, Address3, Address4, Address5, Postcode, Telephone.
// For source data see:
// https://digital.nhs.uk/services/organisation-data-service/data-downloads/gp-and-gp-practice-related-data
func CreatePharmacies(dispensaryCsv string, pcodesLatLng map[string]nhsfinder.LatLng, outputCsv string) error {
	d, err := os.Open(dispensaryCsv)
	if err != nil {
		return err
	}
	defer d.Close()

	rows, _ := csv.NewReader(d).ReadAll()

	pharmacies := []string{}
	for _, row := range rows {
		// A for active and 1 for type Pharmacy
		if row[12] == "A" && row[13] == "1" {
			pcode := row[9]
			latlng := pcodesLatLng[pcode]
			p := fmt.Sprintf("%s,%s,%s,%s,%s,%s,%s,%s,%s,%f,%f",
				row[0], clean(row[1]), clean(row[4]), clean(row[5]), clean(row[6]), clean(row[7]), clean(row[8]), pcode, row[17], latlng.Lat, latlng.Lng)
			pharmacies = append(pharmacies, p)
		}
	}
	return write(outputCsv, pharmacies)
}

func clean(src string) string {
	return strings.Replace(src, ", ", " ", -1)
}

//appendLatLon adds latitude (lat) and longtitude to pharmacy rows
func appendLatLon(pharmacyRows []string, postcodes map[string]nhsfinder.LatLng) {
	fmt.Println(postcodes)
}

func write(file string, pharmacies []string) error {
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()
	last := len(pharmacies) - 1
	for i, v := range pharmacies {
		nl := "\n"
		if i == last {
			nl = ""
		}
		fmt.Fprintf(f, "%s%s", v, nl)
	}
	return nil
}
