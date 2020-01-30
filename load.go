package pharmacyfinder

import (
	"encoding/csv"
	"io"
	"strconv"
	"strings"
)

// LoadLatLngs loads a map of string postcodes to its LatLng
func LoadLatLngs(r io.Reader) (map[string]LatLng, error) {
	cr := csv.NewReader(r)
	var postcodes = make(map[string]LatLng)
	for {
		record, err := cr.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}

		// normalise postcode
		postcode := record[0]
		postcode = strings.TrimSpace(postcode)
		postcode = strings.Replace(postcode, " ", "", 2)

		lat, _ := strconv.ParseFloat(record[1], 64)
		lng, _ := strconv.ParseFloat(record[2], 64)

		postcodes[postcode] = LatLng{
			Lat: lat,
			Lng: lng,
		}
	}
	return postcodes, nil
}

// LoadPharmacies loads pharmacies from specified CSV filename
func LoadPharmacies(r io.Reader) ([]Pharmacy, error) {
	cr := csv.NewReader(r)
	var pharmacies []Pharmacy
	for {
		record, err := cr.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}

		lat, _ := strconv.ParseFloat(record[9], 64)
		lng, _ := strconv.ParseFloat(record[10], 64)

		p := Pharmacy{
			ODSCode: record[0],
			Name:    record[1],
			Address: Address{
				Line1:    record[2],
				Line2:    record[3],
				Line3:    record[4],
				Line4:    record[5],
				Postcode: record[6],
			},
			Phone:  record[9],
			LatLng: LatLng{lat, lng},
		}

		// fmt.Printf("name = %s\n", record[1])
		// fmt.Printf("line1 = %s\n", record[2])
		// fmt.Printf("line2 = %s\n", record[3])
		// fmt.Printf("line3 = %s\n", record[4])
		// fmt.Printf("line4 = %s\n", record[5])
		// fmt.Printf("postcode = %s\n", record[6])
		// fmt.Printf("phone = %s\n", record[7])
		// fmt.Printf("lat/lang = %f/%f\n\n", lat, lng)
		pharmacies = append(pharmacies, p)

	}
	return pharmacies, nil
}
