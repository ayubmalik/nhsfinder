package nhsfinder

import (
	"bytes"
	"fmt"

	"io/ioutil"
	"testing"
)

func TestSimplifyPharmacies(t *testing.T) {

	t.Run("Simplify pharmacies", func(t *testing.T) {
		inputFile := "testdata/Pharmacy.csv"
		goldenFile := "testdata/pharmacy.golden.csv"
		outputFile := "/tmp/pharmacy.csv"

		if err := SimplifyODS(inputFile, outputFile); err != nil {
			t.Fatalf("%v", err)
		}

		expected, _ := ioutil.ReadFile(goldenFile)
		actual, _ := ioutil.ReadFile(outputFile)

		if !bytes.Equal(expected, actual) {
			fmt.Println("----------------------")
			fmt.Println(string(expected))
			fmt.Println("----------------------")
			fmt.Println(string(actual))
			fmt.Println("----------------------")
			t.Fatalf("actual/expected contents not equal")
		}
	})
}
