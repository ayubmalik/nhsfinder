package data

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func TestPharmacySummaries(t *testing.T) {

	t.Run("Only include active pharmacies", func(t *testing.T) {

		goldenFile := "testdata/pharmacies.golden.csv"
		inputFile := "testdata/sample-edispensary.csv"
		outputFile := "/tmp/pharmacies.csv"

		if err := PharmacySummaries(inputFile, outputFile); err != nil {
			t.Fatalf("%v", err)
		}

		expected, _ := ioutil.ReadFile(goldenFile)
		actual, _ := ioutil.ReadFile(outputFile)

		if !bytes.Equal(expected, actual) {
			t.Fatalf("actual/expected contents not equal:\n%s\n\n%s", actual, expected)
		}
	})
}
