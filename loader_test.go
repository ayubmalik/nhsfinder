package pharmacyfinder

import (
	"testing"
)

func TestLoadLatLngs(t *testing.T) {

	t.Run("load lat/lng with postcode spaces removed", func(t *testing.T) {
		tests := []struct {
			input string
			want  LatLng
		}{
			{input: "AB101XG", want: LatLng{57.144165160000000, -2.114847768000000}},
			{input: "AB106RN", want: LatLng{57.137879760000000, -2.121486688000000}},
			{input: "AB129SP", want: LatLng{57.148707080000000, -2.097806027000000}},
		}

		latLngs := LoadLatLngs("testdata/sample_ukpostcodes.csv")

		for _, tc := range tests {
			got, _ := latLngs[tc.input]
			if got != tc.want {
				t.Fatalf("postcode %s expected: %v, got: %v", tc.input, tc.want, got)
			}
		}
	})
}

func TestLoadPharmacies(t *testing.T) {

	tests := []Pharmacy{
		Pharmacy{ODSCode: "FA002", Name: "ROWLANDS PHARMACY", LatLng: LatLng{53.372375, -2.127355}},
		Pharmacy{ODSCode: "FA007", Name: "ROWLANDS PHARMACY", LatLng: LatLng{51.740496, 0.689189}},
		Pharmacy{ODSCode: "FA008", Name: "BOOTS UK LIMITED", LatLng: LatLng{53.802136, -1.544251}},
	}

	t.Run("load pharmacies from CSV", func(t *testing.T) {
		pharmacies := LoadPharmacies("testdata/sample_pharmacies.csv")

		for _, want := range tests {
			found := false
			for _, p := range pharmacies {
				if want.ODSCode == p.ODSCode && want.Name == p.Name && want.LatLng == p.LatLng {
					found = true
					break
				}
			}

			if !found {
				t.Fatalf("wanted pharmmacy (%s/%s/%v) not found in pharmacies", want.ODSCode, want.Name, want.LatLng)
			}
		}
	})
}
