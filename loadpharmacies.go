package nhsfinder

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"

	"golang.org/x/text/encoding/charmap"
)

// LoadPharmacies loads pharmacies from specified CSV filename
func LoadPharmacies(filename string) []Pharmacy {
	datafile, _ := os.Open(filename)
	defer datafile.Close()
r := charmap.Windows1252.NewDecoder().Reader(datafile)
	csvr := csv.NewReader(r)
	csvr.Comma = '¬'

	var pharmacies []Pharmacy
	for {
		record, err := csvr.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		lat, _ := strconv.ParseFloat(record[14], 64)
		lng, _ := strconv.ParseFloat(record[15], 64)
		pharmacies = append(pharmacies, Pharmacy{
			ODSCode: record[0],
			Name:    record[7],

			Address: &Address{
				Line1: record[8],
				Line2: record[9],
				Line3: record[10],
				Line4: record[11],
				Line5: record[12],
				Postcode: Postcode{
					Value: record[13],
					LatLng: LatLng{
						Lat: lat,
						Lng: lng,
					},
				},
			},
			Phone: record[18],
		})
	}
	return pharmacies
}
