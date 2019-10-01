package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"

	"github.com/ayubmalik/nhsfinder"

	"github.com/ayubmalik/nhsfinder/data"
	"github.com/mholt/archiver"
)

const (
	dispensariesURL    = "https://files.digital.nhs.uk/assets/ods/current/edispensary.zip"
	postcodesLatLngURL = "https://www.freemaptools.com/download/full-postcodes/ukpostcodes.zip"
)

func main() {
	tmpDir, _ := ioutil.TempDir("", "nhs")
	defer func() {
		os.RemoveAll(tmpDir)
	}()
	fmt.Println("tmp dir", tmpDir)

	fmt.Println("fetching", dispensariesURL)
	dispZip := fetchURL(dispensariesURL, tmpDir)
	archiver.Unarchive(dispZip, tmpDir)

	fmt.Println("fetching", postcodesLatLngURL)
	pcodesZip := fetchURL(postcodesLatLngURL, tmpDir)
	archiver.Unarchive(pcodesZip, tmpDir)
	pcodesCsv := strings.TrimSuffix(pcodesZip, ".zip") + ".csv"

	pcodes := nhsfinder.LoadPostcodes(pcodesCsv)

	fmt.Println("create pharmacies")
	dispCsv := strings.TrimSuffix(dispZip, ".zip") + ".csv"
	data.CreatePharmacies(dispCsv, pcodes, "/tmp/pharmacies.csv")

	// keep copy of UK postcodes csv
	os.Rename(pcodesCsv, path.Join(os.TempDir(), path.Base(pcodesCsv)))
}

func fetchURL(url string, dir string) string {
	fetcher := data.HTTPFetcher{}
	dest := path.Join(dir, path.Base(url))
	err := fetcher.Fetch(url, dest)
	if err != nil {
		log.Fatalf("could not download %s!\n", url)
	}
	return dest
}
