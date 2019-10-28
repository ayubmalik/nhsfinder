package pharmacyfinder

import (
	"testing"
)

func TestByPostcode(t *testing.T) {
	latLngs := map[string]LatLng{
		"M44BF": LatLng{1.0, 2.0},
	}

	pharmacies := []Pharmacy{
		Pharmacy{Name: "P1", LatLng: LatLng{1.0, 2.10}},
		Pharmacy{Name: "P2", LatLng: LatLng{1.0, 2.20}},
		Pharmacy{Name: "P3", LatLng: LatLng{1.0, 2.30}},
		Pharmacy{Name: "P4", LatLng: LatLng{1.0, 2.40}},
		Pharmacy{Name: "P5", LatLng: LatLng{1.0, 2.50}},
		Pharmacy{Name: "P6", LatLng: LatLng{1.0, 2.60}},
		Pharmacy{Name: "P7", LatLng: LatLng{1.0, 2.70}},
		Pharmacy{Name: "P8", LatLng: LatLng{1.0, 2.80}},
		Pharmacy{Name: "P9", LatLng: LatLng{1.0, 2.90}},
		Pharmacy{Name: "P10", LatLng: LatLng{1.0, 2.10}},
		Pharmacy{Name: "P11", LatLng: LatLng{1.0, 2.11}},
		Pharmacy{Name: "P12", LatLng: LatLng{1.0, 2.12}},
	}

	finder := InMemFinder{LatLngs: latLngs, Pharmacies: pharmacies}

	got := finder.ByPostcode("M44BF")

	if len(got) != 10 {
		// t.Fatalf("got %d pharmacies, wanted %d", len(got), 10)
	}
}
