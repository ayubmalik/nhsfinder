package data

import (
	"encoding/csv"
	"fmt"
	"os"
)

// CreatePharmacies takes NHS ODS pharmacy CSV files and creates CSV of active phamacies.
// See https://digital.nhs.uk/services/organisation-data-service/data-downloads/gp-and-gp-practice-related-data
// for source data files.
func CreatePharmacies(dispensaryCsv string, outputCsv string) error {
	d, err := os.Open(dispensaryCsv)
	if err != nil {
		return err
	}
	defer d.Close()

	drows, _ := csv.NewReader(d).ReadAll()

	fmt.Printf("File %s has rows = %d\n", dispensaryCsv, len(drows))

	dmap := make(map[string][]string)
	for _, row := range drows {
		if row[12] == "A" && row[13] == "1" {
			dmap[row[0]] = row
			fmt.Println(row[0], row[1], row[4:9], row[9], row[17], row[23])
		}
	}

	fmt.Println(len(dmap))
	return nil
}
