package pharmacyfinder

import (
	"encoding/csv"
	"io"

	"fmt"
	"os"
	"strings"

	"golang.org/x/text/encoding/charmap"
)

// SimplifyODS takes NHS ODS data from a file and creates a simplified CSV.
// The resulting CSV file contains no header and only the following fields:
//	 ODSCode, Name, Address1, Address2, Address3, Address4, Postcode, Telephone, Email, Lat, Lng
//
// For source data see:
// 	http://media.nhschoices.nhs.uk/data/foi/Pharmacy.csv
// 	http://media.nhschoices.nhs.uk/data/foi/GP.csv
//
// TODO:use io.Reader / io.Writer params
func SimplifyODS(inputCSV string, outputCSV string) error {
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

// SimplifyPostcodes takes UK postcode and geo data from a file and creates a simplified CSV.
// The resulting CSV file contains no header and "live" postcodes from England only, with the following fields:
// 	Postcode, Name, Address1, Address2, Address3, Address4, Postcode, Telephone, Email, Lat, Lng
//
// For source data see:
// 	https://www.getthedata.com/open-postcode-geo
func SimplifyPostcodes(r io.Reader, w io.Writer) error {
	reader := csv.NewReader(r)
	reader.Comma = ','
	reader.TrimLeadingSpace = true
	reader.Read() // skip first
	rows, _ := reader.ReadAll()

	postcodes := []string{}
	for _, row := range rows {
		if row[1] == "live" && row[6] == "England" {
			pcode := row[0]
			lat := row[7]
			lng := row[8]
			p := fmt.Sprintf("%s,%s,%s", pcode, lat, lng)
			postcodes = append(postcodes, p)
		}
	}

	last := len(postcodes) - 1
	for i, v := range postcodes {
		nl := "\n"
		if i == last {
			nl = ""
		}
		fmt.Fprintf(w, "%s%s", v, nl)
	}
	return nil
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
