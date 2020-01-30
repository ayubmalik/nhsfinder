package pharmacyfinder

import (
	"testing"
)

func TestByPostcode(t *testing.T) {

	t.Run("returns nearest 10", func(t *testing.T) {
		latLngs := map[string]LatLng{
			"M44BF": {1.0, 0.0},
		}

		pharmacies := []Pharmacy{
			{Name: "P1", LatLng: LatLng{1.0, 1}},
			{Name: "P2", LatLng: LatLng{1.0, 2}},
			{Name: "P3", LatLng: LatLng{1.0, 3}},
			{Name: "P4", LatLng: LatLng{1.0, 4}},
			{Name: "P5", LatLng: LatLng{1.0, 5}},
			{Name: "P6", LatLng: LatLng{1.0, 6}},
			{Name: "P7", LatLng: LatLng{1.0, 7}},
			{Name: "P8", LatLng: LatLng{1.0, 8}},
			{Name: "P9", LatLng: LatLng{1.0, 9}},
			{Name: "P10", LatLng: LatLng{1.0, 10}},
			{Name: "P11", LatLng: LatLng{1.0, 11}},
			{Name: "P12", LatLng: LatLng{1.0, 12}},
		}

		finder := InMemFinder{LatLngs: latLngs, Pharmacies: pharmacies}
		got := finder.FindPharmacy("M44BF")

		if len(got) != 10 {
			t.Fatalf("got %d pharmacies, wanted %d", len(got), 10)
		}
	})

	t.Run("returns all if less than 10", func(t *testing.T) {
		latLngs := map[string]LatLng{
			"M44BF": {1.0, 0.0},
		}

		pharmacies := []Pharmacy{
			{Name: "P1", LatLng: LatLng{1.0, 1}},
			{Name: "P2", LatLng: LatLng{1.0, 2}},
		}

		finder := InMemFinder{LatLngs: latLngs, Pharmacies: pharmacies}
		got := finder.FindPharmacy("M44BF")

		if len(got) != 2 {
			t.Fatalf("got %d pharmacies, wanted %d", len(got), 2)
		}
	})
}
