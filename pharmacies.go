package nhsfinder

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

// GetPharmacies loads pharmacies from specified csvFile file and adds postcode info
func GetPharmacies(csvFile string, postcodeDB *PostcodeDB) []Pharmacy {
	pharmacies := loadPharmacies(csvFile)
	for _, p := range pharmacies {
		pcode := p.Address.Postcode.Value
		p.Address.UpdatePostcode(postcodeDB.Postcodes[pcode])
	}
	return pharmacies
}

func loadPharmacies(filename string) []Pharmacy {
	datafile, _ := os.Open(filename)
	defer datafile.Close()
	r := csv.NewReader(datafile)
	var pharmacies []Pharmacy
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		pharmacies = append(pharmacies, Pharmacy{
			ID:   record[0],
			Name: record[1],
			Address: &Address{
				Line1: record[4],
				Line2: record[5],
				Line3: record[6],
				Line4: record[7],
				Line5: record[8],
				Postcode: Postcode{
					Value: record[9],
				},
			},
			Phone: record[17],
		})
	}
	return pharmacies
}
