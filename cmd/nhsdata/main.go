package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"

	"github.com/ayubmalik/nhsfinder/data"
	"github.com/mholt/archiver"
)

const (
	dispensariesURL = "https://files.digital.nhs.uk/assets/ods/current/edispensary.zip"
	headquartersURL = "https://files.digital.nhs.uk/assets/ods/current/epharmacyhq.zip"
)

func main() {
	tmp, _ := ioutil.TempDir("", "nhs")
	defer func() {
		os.RemoveAll(tmp)
	}()
	fmt.Println(tmp)

	dispZip := fetchURL(tmp, dispensariesURL)
	hqZip := fetchURL(tmp, headquartersURL)

	archiver.Unarchive(dispZip, tmp)
	archiver.Unarchive(hqZip, tmp)

	dispCsv := strings.TrimSuffix(dispZip, ".zip") + ".csv"
	hqCsv := strings.TrimSuffix(hqZip, ".zip") + ".csv"
	data.CreatePharmacies(dispCsv, hqCsv, "")
}

func fetchURL(tmp string, url string) string {
	fetcher := data.HTTPFetcher{}
	dest := path.Join(tmp, path.Base(url))
	err := fetcher.Fetch(url, dest)
	if err != nil {
		log.Fatalf("could not download %s!\n", url)
	}
	return dest
}
