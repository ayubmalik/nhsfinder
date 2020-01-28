package pharmacyfinder

import (
	"encoding/csv"

	"fmt"
	"golang.org/x/text/encoding/charmap"
	"os"
	"strings"
)

// Simplify takes NHS ODS data from a file and creates a simplified CSV.
// The resulting pharmacy CSV file contains no header and only the following fields:
//	 ODSCode, Name, Address1, Address2, Address3, Address4, Postcode, Telephone, Email, Lat, Lng
//
// For source data see:
// 	http://media.nhschoices.nhs.uk/data/foi/Pharmacy.csv
// 	http://media.nhschoices.nhs.uk/data/foi/GP.csv
func Simplify(inputCSV string, outputCSV string) error {
	f, err := os.Open(inputCSV)
	if err != nil {
		return err
	}
	defer f.Close()

	iso88591 := charmap.ISO8859_1.NewDecoder().Reader(f)
	reader := csv.NewReader(iso88591)
	reader.Comma = '¬'
	reader.TrimLeadingSpace = true
	reader.Read() // skip first
	rows, _ := reader.ReadAll()

	orgs := []string{}
	for _, row := range rows {
		pcode := row[13]
		lat := row[14]
		lng := row[15]
		p := fmt.Sprintf("%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s",
			row[1], clean(row[7]), clean(row[8]), clean(row[9]), clean(row[10]), clean(row[11]), pcode, row[18], row[19], lat, lng)
		orgs = append(orgs, p)
	}
	return write(outputCSV, orgs)
}

func clean(src string) string {
	return strings.Replace(strings.Trim(src, " "), ",", "", -1)
}

func appendLatLon(pharmacyRows []string, postcodes map[string]LatLng) {
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
