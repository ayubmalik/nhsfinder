package data

import (
	"io/ioutil"
	"strings"
	"testing"
)

func TestPharmacySummaries(t *testing.T) {

	sampleFile := "testdata/sample-edispensary.csv"
	t.Run("Only include active pharmacies", func(t *testing.T) {
		if err := PharmacySummaries(sampleFile, "/tmp/pharmacies.csv"); err != nil {
			t.Fatalf("%v", err)
		}

		buf, _ := ioutil.ReadFile("/tmp/pharmacies.csv")
		contents := string(buf)

		if strings.Contains(contents, "CLOSED1") || strings.Contains(contents, "CLOSED2") {
			t.Errorf("Did not expect closed pharmacy")
		}

		if !strings.Contains(contents, "ACTIVE1") {
			t.Errorf("Expected active pharmacy in '%s'", contents)
		}
	})
}
