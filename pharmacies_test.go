package pharmacyfinder

import (
	"bytes"
	"fmt"

	"io/ioutil"
	"testing"
)

func TestSimplifyPharmacies(t *testing.T) {

	t.Run("Simplify pharmacies", func(t *testing.T) {
		inputFile := "testdata/Pharmacy.csv"
		goldenFile := "testdata/pharmacies.golden.csv"
		outputFile := "/tmp/pharmacies.csv"

		if err := SimplifyODS(inputFile, outputFile); err != nil {
			t.Fatalf("%v", err)
		}

		expected, _ := ioutil.ReadFile(goldenFile)
		actual, _ := ioutil.ReadFile(outputFile)

		fmt.Println(string(actual))

		if !bytes.Equal(expected, actual) {
			t.Fatalf("actual/expected contents not equal:\n%s\n\n%s", actual, expected)
		}
	})
}
