package nhsfinder

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
)

// PostcodeDB is simple db of postcodes loaded from CSV
type PostcodeDB struct {
	Postcodes map[string]Postcode
}

// LoadPostcodeDB loads a PostcodeDB from CSV filename
func LoadPostcodeDB(filename string) *PostcodeDB {
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
	return &PostcodeDB{
		Postcodes: postcodes,
	}
}
