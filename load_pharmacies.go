package nhsfinder

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
)

// LoadPharmacies loads pharmacies from specified CSV filename
func LoadPharmacies(filename string) []Pharmacy {
	datafile, _ := os.Open(filename)
	defer datafile.Close()
	r := csv.NewReader(bufio.NewReader(datafile))
	//csvr.Comma = ','

	var pharmacies []Pharmacy
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Printf(filename)
			log.Fatal(err)
		}

		pharmacies = append(pharmacies, Pharmacy{
			ODSCode: record[0],
			Name:    record[1],
			Address: &Address{
				Line1:    record[2],
				Line2:    record[3],
				Line3:    record[4],
				Line4:    record[5],
				Line5:    record[6],
				Postcode: record[7],
			},
			Phone: record[8],
		})
	}
	return pharmacies
}
