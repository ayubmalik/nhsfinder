package pharmacyfinder

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
)

// LoadLatLngs loads a map of string postcodes to its LatLng
func LoadLatLngs(filename string) map[string]LatLng {
	datafile, _ := os.Open(filename)
	defer datafile.Close()
	r := csv.NewReader(datafile)
	var postcodes = make(map[string]LatLng)
	r.Read() // skip header
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		value := record[1]
		lat, _ := strconv.ParseFloat(record[2], 64)
		lng, _ := strconv.ParseFloat(record[3], 64)
		postcodes[value] = LatLng{
			Lat: lat,
			Lng: lng,
		}
	}
	return postcodes
}

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

		lat, _ := strconv.ParseFloat(record[9], 64)
		lng, _ := strconv.ParseFloat(record[10], 64)

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
			Phone:  record[8],
			LatLng: LatLng{lat, lng},
		})
	}
	return pharmacies
}
