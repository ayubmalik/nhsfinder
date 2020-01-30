package nhsfinder

import (
	"encoding/csv"
	"io"
	"strconv"
	"strings"
)

// LoadLatLngs loads a map of postcode to LatLng.
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

// LoadPharmacies loads pharmacies from specified reader.
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

		lat, _ := strconv.ParseFloat(record[8], 64)
		lng, _ := strconv.ParseFloat(record[9], 64)

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
			Phone:  record[7],
			LatLng: LatLng{lat, lng},
		}
		pharmacies = append(pharmacies, p)
	}
	return pharmacies, nil
}

// LoadGPs loads pharmacies from specified reader.
func LoadGPs(r io.Reader) ([]GP, error) {
	cr := csv.NewReader(r)
	var gps []GP
	for {
		record, err := cr.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}

		lat, _ := strconv.ParseFloat(record[8], 64)
		lng, _ := strconv.ParseFloat(record[9], 64)
		g := GP{
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
		gps = append(gps, g)

	}
	return gps, nil
}
