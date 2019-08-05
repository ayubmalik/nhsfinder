package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

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

	disp := fetchURL(tmp, dispensariesURL)
	hq := fetchURL(tmp, headquartersURL)

	archiver.Unarchive(disp, tmp)
	archiver.Unarchive(hq, tmp)
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
