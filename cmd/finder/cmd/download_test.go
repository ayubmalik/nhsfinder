package cmd

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"testing"
)

type mockDownloader struct {
	fromFile string
}

func (m mockDownloader) Download(src, destFile string) error {
	df, err := os.Create(destFile)
	if err != nil {
		fmt.Println("err", err)
	}
	defer df.Close()

	sf, err := os.Open(m.fromFile)
	if err != nil {
		fmt.Println("err", err)
	}
	defer sf.Close()

	io.Copy(df, sf)
	return nil
}

func TestDownloadPharmacy(t *testing.T) {
	destDir := tempDir()
	defer func() { os.RemoveAll(destDir) }()

	want := path.Join(destDir, "pharmacies.csv")
	downloader := mockDownloader{fromFile: "../../../testdata/Pharmacy.csv"}
	downloadODS(downloader, want)

	info, err := os.Stat(want)
	if err != nil || info.Name() != path.Base(want) {
		t.Errorf("wanted file not created err %v ,file %v", err, info)
	}
}

func TestDownloadGP(t *testing.T) {
	destDir := tempDir()
	defer func() { os.RemoveAll(destDir) }()

	want := path.Join(destDir, "gps.csv")
	downloader := mockDownloader{fromFile: "../../../testdata/GP.csv"}
	downloadODS(downloader, want)

	info, err := os.Stat(want)
	if err != nil || info.Name() != path.Base(want) {
		t.Errorf("wanted file not created err %v ,file %v", err, info)
	}
}

func TestDownloadPostcodes(t *testing.T) {
	destDir := tempDir()
	defer func() { os.RemoveAll(destDir) }()

	want := path.Join(destDir, "open_postcode_geo.csv")
	downloader := mockDownloader{fromFile: "../../../testdata/open_postcode_geo.zip"}
	downloadPostcodes(downloader, want)

	info, err := os.Stat(want)
	if err != nil || info.Name() != path.Base(want) {
		t.Errorf("wanted file not created err %v ,file %v", err, info)
	}
}

func tempDir() string {
	dir, _ := ioutil.TempDir("", "findertest-")
	return dir
}
