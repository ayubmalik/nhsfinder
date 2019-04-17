package nhsfinder

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
)

// LoadPostcodes loads Postcode struct from CSV filename as map keyed off postcode string
func LoadPostcodes(filename string) map[string]Postcode {
	datafile, _ := os.Open(filename)
	defer datafile.Close()
	r := csv.NewReader(datafile)
	var postcodes = make(map[string]Postcode)
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
		postcodes[value] = Postcode{
			Value: value,
			LatLng: LatLng{
				Lat: lat,
				Lng: lng,
			},
		}
	}
	return postcodes
}
