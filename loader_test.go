package pharmacyfinder_test

import (
	"testing"

	"github.com/ayubmalik/pharmacyfinder"
)

func TestLoadLatLngs(t *testing.T) {

	t.Run("exact match from file", func(t *testing.T) {
		tests := []struct {
			input string
			want  pharmacyfinder.LatLng
		}{
			{input: "AB10 1XG", want: pharmacyfinder.LatLng{57.144165160000000, -2.114847768000000}},
			{input: "AB10 6RN", want: pharmacyfinder.LatLng{57.137879760000000, -2.121486688000000}},
			{input: "AB12 5GL", want: pharmacyfinder.LatLng{57.081937920000000, -2.246567389000000}},
			{input: "AB12 9SP", want: pharmacyfinder.LatLng{57.148707080000000, -2.097806027000000}},
		}
		data := pharmacyfinder.LoadLatLngs("testdata/sample_ukpostcodes.csv")

		for _, tc := range tests {
			got, _ := data[tc.input]
			if got != tc.want {
				t.Fatalf("postcode %s expected: %v, got: %v", tc.input, tc.want, got)
			}
		}
	})

}
