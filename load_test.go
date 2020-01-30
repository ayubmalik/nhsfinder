package nhsfinder

import (
	"fmt"
	"os"
	"testing"
)

func TestLoadLatLngs(t *testing.T) {

	postcodeFile, err := os.Open("testdata/postcode.csv")
	if err != nil {
		t.Errorf("could not load postcode file: %v\n", err)
	}
	defer postcodeFile.Close()

	t.Run("load lat/lng with postcode spaces removed", func(t *testing.T) {
		tests := []struct {
			input string
			want  LatLng
		}{
			{input: "AB101XG", want: LatLng{57.144165160000000, -2.114847768000000}},
			{input: "AB106RN", want: LatLng{57.137879760000000, -2.121486688000000}},
			{input: "AB129SP", want: LatLng{57.148707080000000, -2.097806027000000}},
		}

		latLngs, _ := LoadLatLngs(postcodeFile)
		for _, tc := range tests {
			got, _ := latLngs[tc.input]
			if got != tc.want {
				t.Fatalf("postcode %s expected: %v, got: %v", tc.input, tc.want, got)
			}
		}
	})
}

func TestLoadPharmacies(t *testing.T) {

	pharmacyFile, err := os.Open("testdata/pharmacy.golden.csv")
	if err != nil {
		t.Errorf("could not open pharmacy file: %v\n", err)
	}
	defer pharmacyFile.Close()

	tests := []Pharmacy{
		{ODSCode: "FA512", Name: "Lords Pharmacy", LatLng: LatLng{52.244796752929688, 0.4055977463722229}},
		{ODSCode: "FAP38", Name: "LloydsPharmacy Inside Sainsbury's", LatLng: LatLng{54.894672393798828, -2.9472866058349609}},
		{ODSCode: "FC826", Name: "Rutland Late Night Pharmacy", LatLng: LatLng{52.670024871826172, -0.7302858829498291}},
	}

	t.Run("load pharmacies from CSV", func(t *testing.T) {
		pharmacies, _ := LoadPharmacies(pharmacyFile)

		fmt.Println(pharmacies)
		for _, want := range tests {
			found := false
			for _, p := range pharmacies {
				if want.ODSCode == p.ODSCode && want.Name == p.Name && want.LatLng == p.LatLng {
					found = true
					break
				}
			}

			if !found {
				t.Fatalf("wanted pharmacy (%s/%s/%v) not found in pharmacies", want.ODSCode, want.Name, want.LatLng)
			}
		}
	})
}
