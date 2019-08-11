package data

import (
	"encoding/csv"
	"fmt"
	"os"
)

// CreatePharmacies takes NHS ODS pharmacy CSV files and creates CSV of active phamacies.
// See https://digital.nhs.uk/services/organisation-data-service/data-downloads/gp-and-gp-practice-related-data
// for source data files.
func CreatePharmacies(dispensaryCsv string, hqCsv string, outputCsv string) error {
	d, err := os.Open(dispensaryCsv)
	if err != nil {
		return err
	}
	defer d.Close()

	h, err2 := os.Open(hqCsv)
	if err2 != nil {
		return err2
	}

	drows, _ := csv.NewReader(d).ReadAll()
	hrows, _ := csv.NewReader(h).ReadAll()

	fmt.Printf("File %s has rows = %d\n", dispensaryCsv, len(drows))
	fmt.Printf("File %s has rows = %d\n", hqCsv, len(hrows))

	hmap := make(map[string][]string)
	for _, row := range hrows {
		hmap[row[0]] = row
	}

	dmap := make(map[string][]string)
	for _, row := range drows {
		if row[12] == "A" && row[13] == "1" {
			dmap[row[0]] = row

			pkey := row[14]
			parent := hmap[pkey]
			fmt.Println(parent[0], parent[1], row[14], row[1])
		}
	}

	fmt.Println(len(dmap))
	fmt.Println(len(hmap))
	return nil
}
